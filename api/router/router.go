package router

import (
	"github.com/gin-gonic/gin"
	. "landlord/api"
	. "landlord/middleware"
)

func InitRouter(engine *gin.Engine) {

	engine.Use(WithSession, WithContextDb)

	auth := engine.Group("/auth")
	{
		auth.Use(WithAuthApi)
		auth.POST("/login", AuthApi.Login)
		auth.GET("/qqLogin", AuthApi.QQLogin)
		auth.GET("/qqLogin/qqCallback", AuthApi.QQCallback)
		auth.GET("/401", AuthApi.PermissionDenied)
	}

	users := engine.Group("/users")
	{
		users.Use(WithUserApi)
		users.GET("/myself", UserApi.Myself)
		users.PUT("", UserApi.Update)
	}

	achievement := engine.Group("/achievement")
	{
		achievement.Use(WithAchievementApi)
		achievement.GET("/:userId", AchievementApi.GetAchievementByUserId)
		achievement.GET("", AchievementApi.GetAchievementByUserId)
	}

	chat := engine.Group("/chat")
	{
		chat.Use(WithChatApi)
		chat.POST("", ChatApi.Chat)
	}

	game := engine.Group("/games")
	{
		game.Use(WithGameApi)
		game.POST("/ready", GameApi.Ready)
		game.POST("/unReady", GameApi.UnReady)
		game.POST("/bid", GameApi.Bid)
		game.POST("/play", GameApi.Play)
		game.POST("/pass", GameApi.Pass)
	}

	player := engine.Group("/player")
	{
		player.Use(WithPlayerApi)
		player.GET("/cards", PlayerApi.Cards)
		player.GET("/round", PlayerApi.Round)
		player.GET("/ready", PlayerApi.PlayerReady)
		player.GET("/pass", PlayerApi.PlayerPass)
		player.GET("/bidding", PlayerApi.Bidding)
	}

	room := engine.Group("/room")
	{
		room.Use(WithRoomApi)
		room.GET("", RoomApi.Rooms)
		room.GET("/:id", RoomApi.GetById)
		room.POST("", RoomApi.Create)
		room.POST("/join", RoomApi.Join)
		room.POST("/exit", RoomApi.Exit)
	}

	test := engine.Group("/test")
	{
		test.Any("/ws-test.html", Test)
	}
}
