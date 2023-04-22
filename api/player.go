package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	r "landlord/common/response"
)

func Cards(c *gin.Context) {
	user := GetUser(c)
	cards := biz.GetPlayerCards(user)
	r.Success(cards, c)
}

func Round(c *gin.Context) {
	user := GetUser(c)
	round := biz.IsPlayerRound(user)
	r.Success(round, c)
}

func PlayerReady(c *gin.Context) {
	user := GetUser(c)
	ready := biz.IsPlayerReady(user)
	r.Success(ready, c)
}

func PlayerPass(c *gin.Context) {
	user := GetUser(c)
	can := biz.CanPass(user)
	r.Success(can, c)
}

func Bidding(c *gin.Context) {
	user := GetUser(c)
	can := biz.CanBid(user)
	r.Success(can, c)
}
