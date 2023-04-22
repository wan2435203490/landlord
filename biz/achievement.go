package biz

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"landlord/common/enum"
	"landlord/core/component"
	"landlord/db"
	"landlord/db/mysql_model"
	"landlord/pojo"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
)

func CountScore(user *db.User, result *pojo.RoundResult) {
	room := component.RC.GetUserRoom(user.Id)
	var resList []*DTO.ResultScore
	var messages []*ws.GameEnd
	multiple := result.Multiple

	for _, u := range room.UserList {
		msg := ws.EmptyGameEnd()
		var resultScore *DTO.ResultScore
		isWinning := false
		player := room.GetPlayerByUserId(u.Id)
		if result.WinIdentity == enum.Landlord {
			if player.Id == result.LandlordId {
				isWinning = true
				u.Money += float64(multiple * 2)
				resultScore = DTO.NewResultScore(user.UserName, multiple*2, enum.Landlord)
			} else {
				u.Money -= float64(multiple)
				resultScore = DTO.NewResultScore(user.UserName, -multiple, enum.Farmer)
			}
			msg.WinningIdentity = enum.Landlord
		} else {
			if player.Id == result.LandlordId {
				u.Money -= float64(multiple * 2)
				resultScore = DTO.NewResultScore(user.UserName, -2*multiple, enum.Landlord)
			} else {
				isWinning = true
				u.Money += float64(multiple)
				resultScore = DTO.NewResultScore(user.UserName, multiple, enum.Farmer)
			}
			msg.WinningIdentity = enum.Farmer
		}
		resList = append(resList, resultScore)
		msg.IsWinning = isWinning
		messages = append(messages, msg)

		mysql_model.UpdateUser(user)
		updateAchievement(user.Id, isWinning)
	}

	for i, user := range room.UserList {
		endMsg := messages[i]
		endMsg.ResList = resList
		component.NC.Send2User(user.Id, endMsg)
	}
}

func updateAchievement(userId string, isWinning bool) {
	achievement, _ := GetAchievementByUserId(userId)
	if achievement == nil {
		achievement = db.NewAchievement(userId)
		_ = mysql_model.InsertAchievement(achievement)
	}
	if isWinning {
		achievement.IncrWinMatch()
	} else {
		achievement.IncrFailureMatch()
	}
	mysql_model.UpdateAchievement(achievement)
}

func GetAchievementByUserId(userId string) (*db.Achievement, error) {
	_, err := GetUser(userId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("用户不存在")
		} else {
			return nil, err
		}
	}

	achievement, err := mysql_model.GetAchievementByUserId(userId)
	if err != nil {
		return nil, err
	}

	if achievement == nil {
		newAchievement := db.NewAchievement(userId)
		err := mysql_model.InsertAchievement(newAchievement)
		if err != nil {
			return nil, err
		}
		return newAchievement, nil
	}

	return achievement, nil

}
