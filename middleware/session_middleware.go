package middleware

import (
	json2 "encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	r "landlord/common/response"
	"landlord/common/token"
	"landlord/db"
	"strings"
)

const (
	UserSessionKey = "curUser"
)

func WithSession(c *gin.Context) {

	if strings.HasPrefix(c.Request.RequestURI, "/auth") {
		c.Next()
		return
	}

	var user db.User
	session := sessions.Default(c)
	get := session.Get(UserSessionKey)
	if get == nil {
		//session没有的话 用jwt解析Header.Token
		tokenStr := c.GetHeader("Token")
		ok, userId := token.GetUserIdFromToken(tokenStr)
		if !ok {
			c.Abort()
			r.Error(401, "session记录被清除，请退出重新登录", c)
			return
		}
		user.Id = userId
		if db.MySQL.Find(&user).Error != nil {
			c.Abort()
			r.Error(401, "用户不存在，请重新登录", c)
			return
		}
		c.Set(UserSessionKey, &user)
		c.Next()
		return
	}

	err := json2.Unmarshal(get.([]byte), &user)
	if err != nil {
		c.Abort()
		r.Error(401, "session记录被清除，请退出重新登录", c)
		return
	}

	c.Set(UserSessionKey, &user)
	c.Next()
}
