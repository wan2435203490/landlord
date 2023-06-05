package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	. "landlord/api"
	"landlord/common/config"
	"landlord/db"
	. "landlord/middleware"
)

var (
	store = cookie.NewStore([]byte(config.Config.Session.Secret))
	//暂时不用redis
	//store, _ = redis.NewStore(10, "tcp", "localhost:6379", "", []byte(config.Config.Session.Secret))
)

func InitMiddleware(engine *gin.Engine) {
	engine.Use(sessions.Sessions(config.Config.Session.Name, store))
	engine.Use(WithCors, WithSession, WithLimit, WithContextDb)
	//engine.Use(RequestId(TrafficKey), api.SetRequestLogger)
	//engine.Use(CustomError, NoCache)
	//todo log permission
}

func InitRouter(engine *gin.Engine) {

	InitMiddleware(engine)

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//testSession(engine)

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
		game.POST("/give", GameApi.Give)
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

	room := engine.Group("/rooms")
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

func testSession(engine *gin.Engine) {

	engine.GET("/get", func(c *gin.Context) {
		//session, _ := store.Get(c.Request, "user")

		session := sessions.Default(c)
		get := session.Get(config.Config.Session.UserSessionKey).([]byte)
		var user db.User
		_ = json.Unmarshal(get, &user)
		c.JSON(200, user)
		//_, _ = fmt.Fprintf(c.Writer, "name:%s age:%d\n", session.Values["name"], session.Values["age"])
	})
	engine.GET("/set/:value", func(c *gin.Context) {
		value := c.Param("value")
		session := sessions.Default(c)
		user := &db.User{UserName: value}
		userJson, err := json.Marshal(user)
		if err != nil {
			c.Err()
			return
		}
		var u db.User
		_ = json.Unmarshal(userJson, &u)
		fmt.Println(u)
		session.Set(config.Config.Session.UserSessionKey, userJson)
		//session.Set("name", value)
		session.Save()
		s := session.Get(config.Config.Session.UserSessionKey)
		c.JSON(200, s)
	})
}
