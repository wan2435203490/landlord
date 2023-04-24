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
		return
	}

	if a.IfError(a.ExistUser(userId)) != nil {
		return
	}

	achievement := &db.Achievement{UserId: userId}
	if a.IfError(a.FindAchievementByUserId(achievement)) != nil {
		return
	}

	if achievement.Id == "" {
		//achievement不存在
		achievement.Id = strings.ReplaceAll(uuid.NewString(), "-", "")
		if a.IfError(a.CreateAchievement(achievement)) != nil {
			return
		}
	}

	a.OK(achievement)
}
