package biz

import (
	"landlord/common/config"
	"landlord/common/enum"
	"landlord/core/component"
	"landlord/db"
	"landlord/pojo"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
	"time"
)

var RoomBiz Room

type Room struct {
}

//func (r *Room) Func(args ...any) {
//	//todo
//	//method := args[0]
//	r.ExitRoom(args[0].(*db.User))
//}
//
//func init() {
//	component.Register(RoomBiz)
//}

func GetRoomByUser(user *db.User) *pojo.Room {
	return component.RC.GetUserRoom(user.Id)
}

func GetRoomOut(user *db.User, roomId string) *DTO.RoomOut {
	room := component.RC.GetRoom(roomId)
	if !canVisit(user, room) {
		panic("你无权查看本房间的信息")
	}
	result := DTO.ToRoomOut(room)
	for _, player := range result.PlayerList {
		player.Online = component.WS.IsOnline(player.User.Id)
	}
	setCountDown(room, result)
	return result
}

func CreateRoom(user *db.User, title, password string) *pojo.Room {
	return component.RC.CreateRoom(user, title, password)
}

func JoinRoom(user *db.User, dtoRoom *DTO.Room) string {
	dtoRoomId := dtoRoom.Id
	room := component.RC.GetRoom(dtoRoomId)
	if room.RoomStatus == enum.Playing {
		panic("房间正在游戏中，无法加入!")
	}
	password := dtoRoom.Password
	msg := component.RC.JoinRoom(dtoRoomId, password, user)
	component.NC.Send2Room(dtoRoomId, ws.NewPlayerJoin(user))
	return msg
}

func setCountDown(room *pojo.Room, result *DTO.RoomOut) {
	if room.PrePlayTime == 0 {
		result.CountDown = -1
		return
	}
	gap := time.Now().UnixMilli() - room.PrePlayTime/1000
	countDown := config.Config.Landlords.MaxSecondsForEveryRound - gap
	if countDown <= 0 {
		result.CountDown = 0
	} else {
		result.CountDown = int(countDown)
	}
}

func canVisit(user *db.User, room *pojo.Room) bool {
	for _, u := range room.UserList {
		if u.Id == user.Id {
			return true
		}
	}
	return false
}

func ExitRoom(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	dissolved := component.RC.ExitRoom(room.Id, user)
	if !dissolved {
		//如果没解散 通知其他人
		component.NC.Send2Room(room.Id, ws.NewPlayerExit(user))
	}
	return dissolved
}
