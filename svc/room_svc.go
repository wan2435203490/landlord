package svc

import (
	"landlord/common/config"
	"landlord/common/enum"
	"landlord/db"
	"landlord/internal/component"
	"landlord/pojo"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
	"landlord/sdk/service"
	"time"
)

type RoomSvc struct {
	service.Service
}

func (s *RoomSvc) GetRoomByUser(user *db.User) (*pojo.Room, string) {
	return component.RC.GetUserRoom(user.Id), ""
}

func (s *RoomSvc) GetRoomOut(user *db.User, roomId string) (*DTO.RoomOut, string) {

	room, msg := component.RC.GetRoom(roomId)
	if msg != "" {
		return nil, msg
	}
	if !s.canVisit(user, room) {
		return nil, "你无权查看本房间的信息"
	}
	result := DTO.ToRoomOut(room)
	for _, player := range result.PlayerList {
		player.Online = component.WS.IsOnline(player.User.Id)
	}
	s.setCountDown(room, result)
	return result, ""
}

func (s *RoomSvc) CreateRoom(user *db.User, title, password string) (*pojo.Room, string) {
	return component.RC.CreateRoom(user, title, password)
}

func (s *RoomSvc) JoinRoom(user *db.User, dtoRoom *DTO.Room) string {
	dtoRoomId := dtoRoom.Id
	room, msg := component.RC.GetRoom(dtoRoomId)
	if msg != "" {
		return msg
	}
	if room.RoomStatus == enum.Playing {
		return "房间正在游戏中，无法加入!"
	}
	password := dtoRoom.Password
	msg = component.RC.JoinRoom(dtoRoomId, password, user)
	if msg == "" {
		component.NC.Send2Room(dtoRoomId, ws.NewPlayerJoin(user))
	}
	return msg
}

func (s *RoomSvc) setCountDown(room *pojo.Room, result *DTO.RoomOut) {
	if room.PrePlayTime == 0 {
		result.CountDown = -1
		return
	}
	gap := (time.Now().UnixMilli() - room.PrePlayTime) / 1000
	countDown := config.Config.Landlords.MaxSecondsForEveryRound - gap
	if countDown <= 0 {
		result.CountDown = 0
	} else {
		result.CountDown = int(countDown)
	}
}

func (s *RoomSvc) canVisit(user *db.User, room *pojo.Room) bool {
	for _, u := range room.UserList {
		if u.Id == user.Id {
			return true
		}
	}
	return false
}

func (s *RoomSvc) ExitRoom(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	dissolved := component.RC.ExitRoom(room.Id, user)
	if !dissolved {
		//如果没解散 通知其他人
		component.NC.Send2Room(room.Id, ws.NewPlayerExit(user))
	}
	return dissolved
}
