package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	"landlord/common/response"
	"landlord/db"
	"landlord/sdk/api"
)

var AchievementApi achievementApi

type achievementApi struct {
	api.Api
	svc biz.AchievementSvc
}

func (a *achievementApi) GetAchievementByUserId(c *gin.Context) {

	userId := c.Param("userId")
	if userId == "" {
		return
	}

	achievement := a.svc.GetAchievementByUserId(userId)
	if achievement == nil || achievement.Id == "" {
		return
	}

	r.Success[*db.Achievement](achievement, c)

}
