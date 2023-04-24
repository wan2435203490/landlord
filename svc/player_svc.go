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

func (s *PlayerSvc) GetPlayerCards(user *db.User) []*pojo.Card {
	return component.RC.GetUserCards(user.Id)
}

func (s *PlayerSvc) IsPlayerRound(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus != enum.Playing {
		panic("游戏还未开始")
	}
	if room.StepNum == -1 {
		//叫牌未结束
		return false
	}
	player := room.GetPlayerByUserId(user.Id)
	remain := room.StepNum % 3
	if remain == 0 {
		if player.Id != 3 {
			return false
		}
	} else {
		if player.Id != remain {
			return false
		}
	}

	return true
}

func (s *PlayerSvc) IsPlayerReady(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus == enum.Playing {
		panic("游戏已经开始")
	}
	player := room.GetPlayerByUserId(user.Id)
	return player.Ready
}

func (s *PlayerSvc) CanPass(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus != enum.Playing {
		panic("游戏还未开始")
	}
	player := room.GetPlayerByUserId(user.Id)
	if room.PrePlayerId == 0 {
		return player.Identity != enum.Landlord
	}
	return room.PrePlayerId != player.Id
}

func (s *PlayerSvc) CanBid(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus != enum.Playing {
		panic("游戏还未开始")
	}
	if room.StepNum != -1 {
		return false
	}
	player := room.GetPlayerByUserId(user.Id)
	return player.Id == room.BiddingPlayerId
}