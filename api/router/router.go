package router

import (
	"github.com/gin-gonic/gin"
	"landlord/api"
)

func InitRouter(engine *gin.Engine) {

	//auth := engine.Group("/auth")
	//{
	engine.POST("/login", api.Login)
	engine.GET("/qqLogin", api.QQLogin)
	engine.GET("/qqLogin/qqCallback", api.QQCallback)
	engine.GET("/401", api.PermissionDenied)
	//}

	users := engine.Group("/users")
	{
		users.GET("/myself", api.Myself)
		users.PUT("", api.UpdateUser)
	}

	achievement := engine.Group("/achievement")
	{
		achievement.GET("/:userId", api.GetAchievementByUserId)
		achievement.GET("/insert/:userId", api.GenerateUser)
		achievement.GET("", api.GetAchievementByUserId)
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
