package svc

import (
	"landlord/common/enum"
	"landlord/db"
	"landlord/internal/component"
	"landlord/pojo"
	"landlord/sdk/service"
)

type PlayerSvc struct {
	service.Service
}

func (s *PlayerSvc) GetPlayerCards(user *db.User) ([]*pojo.Card, string) {
	return component.RC.GetUserCards(user.Id)
}

func (s *PlayerSvc) IsPlayerRound(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus != enum.Playing {
		//todo 处理错误时的情况
		return false
	}
	if room.StepNum == 0 {
		//叫牌未结束
		return false
	}
	player := room.GetPlayerByUserId(user.Id)
	remain := room.StepNum % 3
	if remain == 0 {
		remain = 3
	}

	return player.Id == remain
}

func (s *PlayerSvc) IsPlayerReady(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus == enum.Playing {
		//游戏已经开始
		//todo 处理错误时的情况
		return false
	}
	player := room.GetPlayerByUserId(user.Id)
	return player.Ready
}

func (s *PlayerSvc) CanPass(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus != enum.Playing {
		//todo 处理错误时的情况
		return false
	}
	player := room.GetPlayerByUserId(user.Id)
	if room.PrePlayerId == 0 {
		return player.Identity != enum.Landlord
	}
	return room.PrePlayerId != player.Id
}

func (s *PlayerSvc) CanBid(user *db.User) int {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus != enum.Playing {
		//todo 处理错误时的情况
		return -1
	}
	if room.StepNum != 0 {
		return -1
	}
	player := room.GetPlayerByUserId(user.Id)
	if player.Id != room.BiddingPlayerId {
		return -1
	}
	return room.Multiple
}
