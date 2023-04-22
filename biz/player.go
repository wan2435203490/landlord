package biz

import (
	"landlord/common/enum"
	"landlord/core/component"
	"landlord/db"
	"landlord/pojo"
)

func GetPlayerCards(user *db.User) []*pojo.Card {
	return component.RC.GetUserCards(user.Id)
}

func IsPlayerRound(user *db.User) bool {
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

func IsPlayerReady(user *db.User) bool {
	room := component.RC.GetUserRoom(user.Id)
	if room.RoomStatus == enum.Playing {
		panic("游戏已经开始")
	}
	player := room.GetPlayerByUserId(user.Id)
	return player.Ready
}

func CanPass(user *db.User) bool {
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

func CanBid(user *db.User) bool {
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
