package pojo

import (
	"landlord/common/enum"
	"landlord/db"
	"sort"
	"sync"
)

var AllPlayerIds = []int{1, 2, 3}

type Room struct {
	Mu sync.Mutex
	//房间号
	Id       string `json:"id"`
	Password string `json:"password"`
	Title    string `json:"title"`
	//房主
	Owner *db.User `json:"owner"`
	//是否设置密码
	Locked bool `json:"locked"`
	//玩家列表
	PlayerList []*Player `json:"playerList"`
	//用户列表
	UserList []*db.User `json:"userList"`
	//房间状态
	RoomStatus enum.RoomStatus `json:"roomStatus"`

	//底分 default 0
	Multiple int `json:"multiple"`
	//每局走的步数，用来控制玩家的出牌回合 -1代表叫牌还未结束
	// 1 2 3
	StepNum int `json:"stepNum"`
	//叫牌的玩家 1 2 3
	BiddingPlayerId int `json:"biddingPlayer"`
	//上一回合玩家打出的牌
	PreCards    []*Card `json:"preCards"`
	PrePlayerId int     `json:"prePlayerId"`
	//上一回合出牌的时间戳
	PrePlayTime  int64             `json:"prePlayTime"`
	Distribution *CardDistribution `json:"distribution"`
	//最后一次叫地主的PlayerId 如果三个人都没叫地主 默认这个人就是地主
	EndBidId int
	//最近一次叫分的PlayerId
	LatestBidId int
}

func (r *Room) ContainsUser(user *db.User) bool {
	for _, u := range r.UserList {
		if user == u {
			return true
		}
	}
	return false
}

func (r *Room) IsFull() bool {
	return len(r.UserList) > 2
}

func (r *Room) CheckPassword(password string) bool {
	return !r.Locked || password == r.Password
}

func (r *Room) GetAvailablePlayerId() int {
	n := len(AllPlayerIds)
	ids, temp := make([]int, 0), make([]int, n)
	copy(temp, AllPlayerIds)
	for _, p := range r.PlayerList {
		ids = append(ids, p.Id)
	}

	for _, i1 := range ids {
		for j0, i0 := range temp {
			if i1 == i0 {
				temp = append(temp[:j0], temp[j0+1:]...)
				break
			}
		}
	}

	if len(temp) == 0 {
		panic("房间已满")
	}

	return temp[0]
}

func (r *Room) Reset() {
	r.Multiple = 0
	r.RoomStatus = enum.Preparing
	r.PreCards = nil
	r.PrePlayerId = 0
	r.StepNum = 0
	r.BiddingPlayerId = 0
	r.PrePlayTime = 0

	for _, player := range r.PlayerList {
		player.Reset()
	}
}

func (r *Room) GetLandlord() *Player {
	for _, player := range r.PlayerList {
		if player.IsLandlord() {
			return player
		}
	}
	return nil
}

func (r *Room) GetFarmers() []*Player {
	var ret []*Player
	for _, player := range r.PlayerList {
		if !player.IsLandlord() {
			ret = append(ret, player)
		}
	}
	return ret
}

func (r *Room) IsLocked() bool {
	return r.Locked && r.Password != ""
}

func (r *Room) RemovePlayer(userId string) {
	for i, player := range r.PlayerList {
		if player.User.Id == userId {
			r.PlayerList = append(r.PlayerList[:i], r.PlayerList[i+1:]...)
			break
		}
	}
}

func (r *Room) RemoveUser(userId string) {
	for i, user := range r.UserList {
		if user.Id == userId {
			r.UserList = append(r.UserList[:i], r.UserList[i+1:]...)
			break
		}
	}
}

func (r *Room) IncrStep() {
	r.StepNum++
}

func (r *Room) IncrBiddingPlayer() {
	if r.BiddingPlayerId == 3 {
		r.BiddingPlayerId = 1
	} else {
		r.BiddingPlayerId++
	}
}

func (r *Room) DoubleMultiple() {
	r.Multiple *= 2
}

func (r *Room) GetPlayerByUserId(userId string) *Player {
	for _, player := range r.PlayerList {
		if player.User.Id == userId {
			return player
		}
	}
	return nil
}

func (r *Room) GetPlayer(playerId int) *Player {
	for _, player := range r.PlayerList {
		if player.Id == playerId {
			return player
		}
	}
	return nil
}

func (r *Room) GetUserByPlayerId(playerId int) *db.User {
	player := r.GetPlayer(playerId)

	if player == nil {
		return nil
	}

	return player.User
}

func (r *Room) IsAllReady() bool {
	if len(r.PlayerList) != 3 {
		return false
	}

	for _, player := range r.PlayerList {
		if !player.Ready {
			return false
		}
	}

	return true
}

func (r *Room) GetUserIds() []string {
	var ret []string
	for _, user := range r.UserList {
		ret = append(ret, user.Id)
	}
	return ret
}

func (r *Room) SortPlayerList() {
	sort.Slice(r.PlayerList, func(i, j int) bool {
		return r.PlayerList[i].Id < r.PlayerList[j].Id
	})
}
