package main

import (
	"github.com/gin-gonic/gin"
	"landlord/api/router"
	"landlord/internal/component"
	"log"
)

func main() {

	engine := gin.Default()

	router.InitRouter(engine)

	//ws
	go component.Start()

	//log.Fatal(engine.RunTLS(":8080", config.Config.TLS.Cert, config.Config.TLS.Key))
	log.Fatal(engine.Run(":8080"))

}
