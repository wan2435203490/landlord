package api

import (
	"github.com/gin-gonic/gin"
	"landlord/sdk/api"
	"landlord/svc"
)

var PlayerApi playerApi

type playerApi struct {
	api.Api
	svc.PlayerSvc
}

func (a *playerApi) Cards(c *gin.Context) {
	user := a.User()
	cards := a.GetPlayerCards(user)
	a.OK(cards)
}

func (a *playerApi) Round(c *gin.Context) {
	user := a.User()
	round := a.IsPlayerRound(user)
	a.OK(round)
}

func (a *playerApi) PlayerReady(c *gin.Context) {
	user := a.User()
	ready := a.IsPlayerReady(user)
	a.OK(ready)
}

func (a *playerApi) PlayerPass(c *gin.Context) {
	user := a.User()
	can := a.CanPass(user)
	a.OK(can)
}

func (a *playerApi) Bidding(c *gin.Context) {
	user := a.User()
	can := a.CanBid(user)
	a.OK(can)
}
