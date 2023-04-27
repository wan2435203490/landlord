package svc

import (
	"encoding/json"
	"fmt"
	"landlord/common/enum"
	"landlord/common/utils"
	"landlord/db"
	"landlord/internal"
	"landlord/internal/component"
	"landlord/pojo"
	"landlord/pojo/ws"
	"landlord/sdk/service"
	"log"
	"math/rand"
	"time"
)

type GameSvc struct {
	service.Service
}

func (s *GameSvc) ReadyGame(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	player := room.GetPlayerByUserId(user.Id)
	player.Ready = true

	component.RC.UpdateRoom(room)

	component.NC.Send2Room(room.Id, ws.NewReadyGame(user.Id))

	isAllReady := room.IsAllReady()

	if isAllReady {
		room.Mu.Lock()
		defer room.Mu.Unlock()
		s.StartGame(room)
	}

	return isAllReady
}

func (s *GameSvc) UnReadyGame(user *db.User) string {
	room := component.RC.GetUserRoom(user.Id)
	player := room.GetPlayerByUserId(user.Id)
	player.Ready = false

	component.RC.UpdateRoom(room)
	component.NC.Send2Room(room.Id, ws.NewUnReadyGame(user.Id))
	return "success"
}

// Want 低于3分时，NextPlayerBid
func (s *GameSvc) Want(user *db.User, score int) string {
	room := component.RC.GetUserRoom(user.Id)

	if room.Multiple >= score {
		return "抢地主失败：比上一家叫分低"
	}
	log.Printf("[%s] 玩家 %s 叫牌，分数为 %d 分\n", room.Id, user.UserName, score)

	player := room.GetPlayerByUserId(user.Id)

	room.Multiple = score
	//记录最近一次叫分的playerId
	room.LatestBidId = player.Id

	if score > 2 {
		var landlord *db.User
		for _, player := range room.PlayerList {
			if player.User.Id == user.Id {
				if player.Id != room.BiddingPlayerId {
					return "不是当前用户的叫牌回合"
				}
				landlord = player.User
				room.StepNum = player.Id
				player.Identity = enum.Landlord
				player.AddCards(room.Distribution.TopCards)
			} else {
				player.Identity = enum.Farmer
			}
		}
		if landlord == nil {
			return "选取的地主玩家不能为空"
		}
		room.PrePlayTime = time.Now().UnixMilli()
		component.RC.UpdateRoom(room)
		component.NC.Send2Room(room.Id, ws.NewBidEnd())
		component.NC.Send2User(landlord.Id, ws.NewPleasePlayCard())
		log.Printf("[%s] 玩家 %s 成为地主", room.Id, landlord.UserName)

	} else {
		nextPlayerId := player.GetNextPlayerId()
		//bid一圈
		if room.BiddingPlayerId == room.EndBidId {
			s.MustWant(room.LatestBidId, room)
			return ""
		}
		room.IncrBiddingPlayer()

		nextUser := room.GetUserByPlayerId(nextPlayerId)
		fmt.Printf("[%s] 玩家 %d 抢地主，分数为", room.Id,
			player.Id, score)

		component.NC.Send2User(nextUser.Id, ws.NewBid(score))
	}

	return ""
}

// MustWant 叫了一圈地主 没有叫3分的情况 都没人叫地主就默认第一家是地主 叫1分
func (s *GameSvc) MustWant(landlordPlayerId int, room *pojo.Room) {
	var landlord *db.User
	for _, player := range room.PlayerList {
		if player.Id == landlordPlayerId {
			landlord = player.User
			room.StepNum = player.Id
			player.Identity = enum.Landlord
			player.AddCards(room.Distribution.TopCards)
		} else {
			player.Identity = enum.Farmer
		}
	}
	if room.Multiple < 1 {
		room.Multiple = 1
	}

	room.PrePlayTime = time.Now().UnixMilli()
	component.RC.UpdateRoom(room)
	component.NC.Send2Room(room.Id, ws.NewBidEnd())
	component.NC.Send2User(landlord.Id, ws.NewPleasePlayCard())
	log.Printf("[%s] 玩家 %s 成为地主", room.Id, landlord.UserName)
}

func (s *GameSvc) NoWant(user *db.User) {
	room := component.RC.GetUserRoom(user.Id)
	player := room.GetPlayerByUserId(user.Id)
	nextPlayerId := player.GetNextPlayerId()

	room.IncrBiddingPlayer()
	//bid一圈
	if room.BiddingPlayerId == room.EndBidId {
		landlordPlayerId := utils.IfThen(room.LatestBidId > 0, room.LatestBidId, nextPlayerId).(int)
		s.MustWant(landlordPlayerId, room)
		return
	}

	nextUser := room.GetUserByPlayerId(nextPlayerId)
	fmt.Printf("[%s] 玩家 %d 选择不叫，由下家 %d 玩家叫牌", room.Id,
		player.Id, nextPlayerId)

	component.NC.Send2User(nextUser.Id, ws.NewBid(0))
}

func (s *GameSvc) PlayCard(user *db.User, cardList []*pojo.Card) (*pojo.RoundResult, string) {
	room := component.RC.GetUserRoom(user.Id)
	marshal, _ := json.Marshal(cardList)
	fmt.Printf("[%s] 玩家 %s 出牌: %s", room.Id, user.UserName, string(marshal))

	player := room.GetPlayerByUserId(user.Id)

	cardType := internal.GetCardsType(cardList...)
	if cardType == -1 {
		fmt.Printf("[%s] 玩家 %s 打出的牌不符合规则", room.Id, user.UserName)
		return nil, "玩家打出的牌不符合规则"
	}
	if room.PreCards != nil && room.PrePlayerId != player.Id {
		preType := internal.GetCardsType(room.PreCards...)
		canPlay := internal.CanPlayCards(cardList, room.PreCards, cardType, preType)
		if !canPlay {
			return nil, "该玩家出的牌管不了上家"
		}
	}
	removeNextPlayerRecentCards(room, player)
	player.RecentCards = cardList
	player.RemoveCards(cardList)

	msg := ws.NewPlayCard(user, cardList, cardType)
	component.NC.Send2Room(room.Id, msg)

	if cardType == enum.Bomb || cardType == enum.JokerBomb {
		room.DoubleMultiple()
	}
	var result *pojo.RoundResult
	if len(player.Cards) == 0 {
		if isSpring(room, player) {
			room.DoubleMultiple()
		}
		fmt.Printf("[%s] 游戏结束，%s 获胜！", room.Id, player.Identity.GetIdentity())
		result = getResult(room, player)
		room.Reset()
	} else {
		fmt.Printf("[%s] 玩家 %s 出牌，类型为 %s，下一个出牌者序号为：%d", room.Id,
			player.User.UserName, cardType.GetType(), player.GetNextPlayerId())
		room.PreCards = cardList
		room.PrePlayerId = player.Id
		room.IncrStep()
		nextUser := room.GetUserByPlayerId(player.GetNextPlayerId())
		component.NC.Send2User(nextUser.Id, ws.NewPleasePlayCard())
	}
	room.PrePlayTime = time.Now().UnixMilli()
	component.RC.UpdateRoom(room)
	return result, ""
}

func (s *GameSvc) PassGame(user *db.User) {
	room := component.RC.GetUserRoom(user.Id)
	player := room.GetPlayerByUserId(user.Id)
	removeNextPlayerRecentCards(room, player)
	room.IncrStep()
	room.PrePlayTime = time.Now().UnixMilli()
	component.RC.UpdateRoom(room)

	fmt.Printf("[%s] 玩家 %s 要不起，下一个出牌者序号为：%d", room.Id, user.UserName, player.GetNextPlayerId())
	nextUser := room.GetUserByPlayerId(player.GetNextPlayerId())
	component.NC.Send2User(nextUser.Id, ws.NewPleasePlayCard())
	component.NC.Send2Room(room.Id, ws.NewPass(user))
}

func (s *GameSvc) StartGame(room *pojo.Room) string {
	if room.RoomStatus == enum.Playing {
		return "房间游戏已经开始"
	} else {
		room.RoomStatus = enum.Playing
	}

	distribution := &pojo.CardDistribution{}
	room.Distribution = distribution
	distribution.Refresh()

	for _, player := range room.PlayerList {
		cards := distribution.GetCards(player.Id)
		player.AddCards(cards)
		player.Ready = true
	}

	roomId := room.Id
	component.NC.Send2Room(roomId, ws.NewStartGame(roomId))

	rand.Seed(time.Now().Unix())
	n := rand.Intn(3) + 1
	room.BiddingPlayerId = n
	room.EndBidId = n + 3
	player := room.GetPlayer(n)
	component.NC.Send2User(player.User.Id, ws.NewBid(0))
	component.RC.UpdateRoom(room)

	return ""
}

// GiveCards 推荐出牌
func (s *GameSvc) GiveCards(user *db.User) [][]*pojo.Card {

	room := component.RC.GetUserRoom(user.Id)
	if room == nil || room.PreCards == nil {
		return nil
	}

	player := room.GetPlayerByUserId(user.Id)

	given := internal.GivePlayCards(player.Cards, room.PreCards)

	return given
}

func isSpring(room *pojo.Room, winner *pojo.Player) bool {
	if winner.IsLandlord() {
		for _, player := range room.GetFarmers() {
			if len(player.Cards) < 17 {
				return false
			}
		}
		return true
	} else {
		return len(room.GetLandlord().Cards) == 17
	}
}

func getResult(room *pojo.Room, player *pojo.Player) *pojo.RoundResult {
	result := &pojo.RoundResult{
		WinIdentity: player.Identity,
		Multiple:    room.Multiple,
	}

	for _, player := range room.PlayerList {
		if player.Identity == enum.Landlord {
			result.LandlordId = player.Id
		}
	}

	return result
}

func removeNextPlayerRecentCards(room *pojo.Room, player *pojo.Player) {
	nextPlayer := room.GetPlayer(player.GetNextPlayerId())
	nextPlayer.ClearRecentCards()
}
