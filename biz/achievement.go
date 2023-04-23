package biz

import (
	"errors"
	"gorm.io/gorm"
	"landlord/common/enum"
	"landlord/db"
	"landlord/db/mysql_model"
	"landlord/internal/component"
	"landlord/pojo"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
	"landlord/sdk/service"
)

type AchievementSvc struct {
	service.Service
}

func (s *AchievementSvc) CountScore(user *db.User, result *pojo.RoundResult) {
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
		s.updateAchievement(user.Id, isWinning)
	}

	for i, user := range room.UserList {
		endMsg := messages[i]
		endMsg.ResList = resList
		component.NC.Send2User(user.Id, endMsg)
	}
}

func (s *AchievementSvc) updateAchievement(userId string, isWinning bool) {
	achievement, _ := s.GetAchievementByUserId(userId)
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

func (s *AchievementSvc) GetAchievementByUserId(userId string) *db.Achievement {
	err := s.Orm.Find(&db.User{}, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.ErrorMsg("用户不存在")
			return nil
		}
		s.Error(err)
		return nil
	}

	achievement := &db.Achievement{}
	s.Orm.Find(&achievement).Where("user_id=?", userId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Error(err)
		return nil
	}

	//if achievement.Id == "" {
	//	newAchievement := db.NewAchievement(userId)
	//	err = s.Orm.Create(newAchievement).Error
	//	if err != nil {
	//		s.Error(err)
	//		return nil
	//	}
	//	return newAchievement
	//}

	return achievement

}

func (s *AchievementSvc) InsertAchievement(achievement *db.Achievement) *db.Achievement {

	err := s.Orm.Create(achievement).Error
	if err != nil {
		s.Error(err)
		return nil
	}
	return achievement

}