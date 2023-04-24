package middleware

import (
	json2 "encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	r "landlord/common/response"
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
		//_ = c.AbortWithError(401, errors.New("session记录被清除，请退出重新登录。"))
		c.Abort()
		r.Error(401, "session记录被清除，请退出重新登录", c)
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
