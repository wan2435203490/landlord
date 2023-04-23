package router

import (
	"github.com/gin-gonic/gin"
	"landlord/api"
	"landlord/sdk"
)

// WithContextDb middleware todo 统一优化
func WithContextDb(c *gin.Context) {
	c.Set("db", sdk.Runtime.GetDbByKey(c.Request.Host).WithContext(c))
	c.Next()
}

// WithUserApi users
func WithUserApi(c *gin.Context) {
	api.UserApi.MakeContext(c)
	c.Next()
}
func WithAuthApi(c *gin.Context) {
	api.AuthApi.MakeContext(c)
	c.Next()
}
func WithAchievementApi(c *gin.Context) {
	api.AuthApi.MakeContext(c)
	c.Next()
}

func InitRouter(engine *gin.Engine) {

	engine.Use(WithContextDb)

	//todo auth放到group
	engine.Use(WithAuthApi)
	engine.POST("/login", api.AuthApi.Login)
	engine.GET("/qqLogin", api.AuthApi.QQLogin)
	engine.GET("/qqLogin/qqCallback", api.AuthApi.QQCallback)
	engine.GET("/401", api.AuthApi.PermissionDenied)

	g1 := engine.Group("/users")
	{
		g1.Use(WithUserApi)
		g1.GET("/myself", api.UserApi.Myself)
		g1.PUT("", api.UserApi.UpdateUser)
	}

	achievement := engine.Group("/achievement")
	{
		g1.Use(WithAchievementApi)
		achievement.GET("/:userId", api.AchievementApi.GetAchievementByUserId)
		achievement.GET("", api.AchievementApi.GetAchievementByUserId)
	}

	chat := engine.Group("/chat")
	{
		chat.POST("", api.Chat)
	}

	games := engine.Group("/games")
	{
		games.POST("/ready", api.ReadyGame)
		games.POST("/unReady", api.UnReady)
		games.POST("/bid", api.Bid)
		games.POST("/play", api.Play)
		games.POST("/pass", api.GamesPass)
	}

	player := engine.Group("/player")
	{
		player.GET("/cards", api.Cards)
		player.GET("/round", api.Round)
		player.GET("/ready", api.PlayerReady)
		player.GET("/pass", api.PlayerPass)
		player.GET("/bidding", api.Bidding)
	}

	rooms := engine.Group("/rooms")
	{
		rooms.GET("", api.Rooms)
		rooms.GET("/:id", api.GetRoomById)
		rooms.POST("", api.CreateRoom)
		rooms.POST("/join", api.Join)
		rooms.POST("/exit", api.Exit)
	}

	test := engine.Group("/test")
	{
		test.Any("/ws-test.html", api.Test)
	}
}
