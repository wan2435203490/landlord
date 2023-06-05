package svc

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"landlord/db"
	"landlord/internal/component"
	"landlord/pojo"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
	//"landlord/sdk/service"
	"strings"
)

type AchievementSvc struct {
	//service.Service
}

func (s *AchievementSvc) CountScore(user *db.User, result *pojo.RoundResult) {
	room := component.RC.GetUserRoom(user.Id)
	var resList []*DTO.ResultScore
	var messages []*ws.GameEnd
	multiple := result.Multiple

	for _, u := range room.UserList {
		player := room.GetPlayerByUserId(u.Id)

		isWin := player.Identity == result.WinIdentity

		resultScore := DTO.NewResultScore(u.UserName, multiple, isWin, player.IsLandlord())
		resList = append(resList, resultScore)

		msg := ws.NewGameEnd(result.WinIdentity, isWin)
		messages = append(messages, msg)

		if s.UpdatesUser(u) != nil {
			return
		}

		if s.updateAchievement(u.Id, isWin) != nil {
			return
		}
	}

	for i, user := range room.UserList {
		endMsg := messages[i]
		endMsg.ResList = resList
		component.NC.Send2User(user.Id, endMsg)
	}
}

func (s *AchievementSvc) updateAchievement(userId string, isWinning bool) error {
	achievement := &db.Achievement{UserId: userId}
	if err := s.FindAchievementByUserId(achievement); err != nil {
		return err
	}
	achievement.CalculateScore(isWinning)

	if achievement.Id == "" {
		achievement.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
		return s.CreateAchievement(achievement)
	} else {
		return s.UpdateAchievement(achievement)
	}
}

// gorm 暂时没必要抽象gorm层 目前AddError都在gorm层
func (s *AchievementSvc) UpdateAchievement(achievement *db.Achievement) error {
	return db.MySQL.Updates(achievement).Error
}

func (s *AchievementSvc) ExistUser(userId string) error {
	err := db.MySQL.Find(&db.User{}, "id=?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("用户不存在")
		}
		return err
	}
	return nil
}

func (s *AchievementSvc) UpdatesUser(user *db.User) error {
	return db.MySQL.Updates(user).Error
}

func (s *AchievementSvc) FindAchievementByUserId(achievement *db.Achievement) error {
	err := db.MySQL.Find(&achievement).Where("user_id=?", achievement.UserId).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}

func (s *AchievementSvc) CreateAchievement(achievement *db.Achievement) error {
	return db.MySQL.Create(achievement).Error
}
