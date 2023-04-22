package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	"landlord/common/response"
	"landlord/db"
	"net/http"
)

func GenerateUser(c *gin.Context) {
	userId := c.Param("userId")
	user, err := biz.GenerateUser(userId)
	if err != nil {
		r.Error(http.StatusInternalServerError, err.Error(), c)
		return
	}
	r.Success(user, c)
}

func GetAchievementByUserId(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		r.Error(http.StatusBadRequest, "userId is empty", c)
		return
	}

	achievement, err := biz.GetAchievementByUserId(userId)

	if err != nil {
		r.Error(http.StatusInternalServerError, err.Error(), c)
		return
	}

	r.Success[*db.Achievement](achievement, c)

}
