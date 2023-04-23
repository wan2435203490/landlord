package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"landlord/api/router"
	"landlord/common/utils"
	"landlord/internal/component"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	engine := gin.Default()
	engine.Use(utils.CorsHandler())

	//store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(""))
	store := cookie.NewStore([]byte("qwq"))
	engine.Use(sessions.Sessions("landlord-session", store))
	router.InitRouter(engine)
	engine.GET("/set", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("user", "1111")
		session.Save()
	})
	engine.GET("/get", func(c *gin.Context) {
		session := sessions.Default(c)
		userJson := session.Get("curUser")
		c.Writer.Write(userJson.([]byte))
	})

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	go component.Start()

	log.Fatal(engine.Run(":8080"))

}
