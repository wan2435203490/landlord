package api

import (
	"github.com/gin-gonic/gin"
	"landlord/sdk"
)

func WithContextDb(c *gin.Context) {
	//当前只有一个db
	c.Set("db", sdk.Runtime.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}

// WithUserApi users
func WithUserApi(c *gin.Context) {
	UserApi.MakeContext(c)
	c.Next()
}
func WithAuthApi(c *gin.Context) {
	if AuthApi.MakeContext(c) != nil {
		c.Abort()
		return
	}
	//AuthApi.svc.Orm = AuthApi.Orm
	//AuthApi.svc.Api = &(AuthApi.Api)

	//svc := AuthApi.svc
	//AuthApi.MakeService((&svc.UserSvc{}).Get())
	c.Next()
}
func WithAchievementApi(c *gin.Context) {
	AchievementApi.MakeContext(c)
	c.Next()
}
func WithChatApi(c *gin.Context) {
	ChatApi.MakeContext(c)
	c.Next()
}
func WithGameApi(c *gin.Context) {
	GameApi.MakeContext(c)
	c.Next()
}
func WithPlayerApi(c *gin.Context) {
	PlayerApi.MakeContext(c)
	c.Next()
}
func WithRoomApi(c *gin.Context) {
	RoomApi.MakeContext(c)
	c.Next()
}
