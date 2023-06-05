package component

import (
	"fmt"
	"landlord/common/enum"
	"landlord/db"
	"landlord/pojo"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var RC RoomComponent

type RoomComponent struct {
	//用户玩家当前所在的房间号映射Map <userId, roomId>
	UserRoomMap sync.Map
	//房间号和与该房间所对应的Room对象映射Map <roomId, room>
	//todo 改成db
	RoomMap sync.Map
}

// CreateRoom 创建房间
func (rc *RoomComponent) CreateRoom(user *db.User, title, roomPassword string) (*pojo.Room, string) {
	rid := rc.getUserRoomId(user.Id)
	if rid != "" {
		return nil, fmt.Sprintf("用户已在房间号为 %s 的房间", rid)
	}

	roomId := rc.randRoomId()
	room := &pojo.Room{
		Id:              roomId,
		RoomStatus:      enum.Preparing,
		Multiple:        0,
		PrePlayerId:     0,
		StepNum:         0,
		BiddingPlayerId: 0,
		Title:           title,
		Owner:           user,
	}

	player := &pojo.Player{
		Id:   1,
		User: user,
	}

	room.UserList = append(room.UserList, user)
	room.PlayerList = append(room.PlayerList, player)

	if roomPassword != "" {
		room.Locked = true
		room.Password = roomPassword
	}

	rc.RoomMap.Store(roomId, room)
	rc.setUserRoom(user.Id, roomId)
	return room, ""
}

// JoinRoom 加入房间
func (rc *RoomComponent) JoinRoom(roomId, roomPassword string, user *db.User) string {
	rId := rc.getUserRoomId(user.Id)
	if rId != "" {
		return fmt.Sprintf("用户已在房间号为 %s 的房间", rId)
	}
	if room, ok := rc.RoomMap.Load(roomId); !ok {
		return "该房间不存在，请核实您输入的房间号！"
	} else {
		r := room.(*pojo.Room)
		if r.ContainsUser(user) {
			return "您已经加入此房间，无法重复加入！"
		}
		if r.IsFull() {
			return "该房间已满，请寻找其他房间！"
		}
		if !r.CheckPassword(roomPassword) {
			return "对不起，您输入的房间密码有误！"
		}

		//playerId(座位序号) 可能是1 2 3 这里取不存在的player最小id，以后实现选座位
		playerId := r.GetAvailablePlayerId()
		player := &pojo.Player{Id: playerId, User: user}
		r.UserList = append(r.UserList, user)
		r.PlayerList = append(r.PlayerList, player)

		rc.setUserRoom(user.Id, roomId)

		return ""
	}
}

// ExitRoom 退出房间 return 房间是否已解散
func (rc *RoomComponent) ExitRoom(roomId string, user *db.User) bool {
	rc.UserRoomMap.Delete(user.Id)

	if value, ok := rc.RoomMap.Load(roomId); ok {

		room := value.(*pojo.Room)
		room.RemoveUser(user.Id)
		room.RemovePlayer(user.Id)

		if len(room.PlayerList) == 0 {
			rc.RoomMap.Delete(roomId)
			return true
		}
	}

	return false
}

func (rc *RoomComponent) ListRooms() []*pojo.Room {
	var ret []*pojo.Room

	rc.RoomMap.Range(func(key, value any) bool {
		room := value.(*pojo.Room)
		if room != nil {
			room.SortPlayerList()
			ret = append(ret, room)
		}
		return true
	})
	return ret
}

func (rc *RoomComponent) GetRoom(roomId string) (*pojo.Room, string) {
	if room, ok := rc.RoomMap.Load(roomId); !ok {
		return nil, "该房间不存在，请核实您输入的房间号！"
	} else {
		return room.(*pojo.Room), ""
	}
}

func (rc *RoomComponent) UpdateRoom(new *pojo.Room) string {
	if _, ok := rc.RoomMap.Load(new.Id); !ok {
		return "该房间不存在！"
	} else {
		rc.RoomMap.Store(new.Id, new)
	}
	return ""
}

func (rc *RoomComponent) GetUserCards(userId string) ([]*pojo.Card, string) {
	room := rc.GetUserRoom(userId)
	player := room.GetPlayerByUserId(userId)
	if player == nil {
		return nil, "未找到该玩家！"
	}
	return player.Cards, ""
}

// GetUserRoom 获取当前用户所在的房间对象
func (rc *RoomComponent) GetUserRoom(userId string) *pojo.Room {
	if roomId, ok := rc.UserRoomMap.Load(userId); ok {
		if room, ok := rc.RoomMap.Load(roomId); ok {
			return room.(*pojo.Room)
		} else {
			//todo
			panic(fmt.Sprintf("未找到对应房间:%s", roomId))
		}
	} else {
		panic("玩家还未加入房间内")
	}
}

func (rc *RoomComponent) randRoomId() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(100000)
	if _, ok := rc.RoomMap.Load(n); ok {
		return rc.randRoomId()
	} else {
		return strconv.Itoa(n)
	}
}

func (rc *RoomComponent) setUserRoom(userId, roomId string) {
	rc.UserRoomMap.Store(userId, roomId)
}

func (rc *RoomComponent) getUserRoomId(userId string) string {
	if value, ok := rc.UserRoomMap.Load(userId); ok {
		return value.(string)
	} else {
		return ""
	}
}
