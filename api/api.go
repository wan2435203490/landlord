package api

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"landlord/biz"
	"landlord/common/token"
	"landlord/db"
)

func GetUser(c *gin.Context) *db.User {
	session := sessions.Default(c)
	get := session.Get("curUser")
	if get == nil {
		tokenStr := c.GetHeader("Token")
		success, userId := token.GetUserIdFromToken(tokenStr)
		if success {
			user, err := biz.GetUser(userId)
			if err != nil {
				panic(err)
			}
			return user
		} else {
			var user db.User
			if err := c.ShouldBindJSON(&user); err != nil {
				panic("user bind error:" + err.Error())
			}
			return &user
		}
	}

	var user db.User
	err := json.Unmarshal(get.([]byte), &user)
	if err != nil {
		panic(err)
	}

	return &user
}
