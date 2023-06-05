package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"landlord/db"
	"landlord/sdk/api"
	"landlord/svc"
	"strings"
)

var AchievementApi achievementApi

type achievementApi struct {
	api.Api
	svc.AchievementSvc //本名的svc放这里 其他的通过实例调用
}

func (a *achievementApi) GetAchievementByUserId(c *gin.Context) {

	userId := a.Param("userId")
	if userId == "" {
		a.ErrorInternal("userId is empty")
		return
	}

	if a.IsError(a.ExistUser(userId)) {
		return
	}

	achievement := &db.Achievement{UserId: userId}
	if a.IsError(a.FindAchievementByUserId(achievement)) {
		return
	}

	if achievement.Id == "" {
		//achievement不存在
		achievement.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
		if a.IsError(a.CreateAchievement(achievement)) {
			return
		}
	}

	a.OK(achievement)
}
