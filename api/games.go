package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	r "landlord/common/response"
	"landlord/db"
	"landlord/pojo"
	"landlord/pojo/DTO"
)

func ReadyGame(c *gin.Context) {
	user := GetUser(c)
	r.Success(biz.ReadyGame(user), c)
}

func UnReady(c *gin.Context) {
	user := GetUser(c)
	biz.UnReadyGame(user)
	r.Success("success", c)
}

func Bid(c *gin.Context) {
	var bid *DTO.Bid
	if err := c.ShouldBindJSON(bid); err != nil {
		panic(err.Error())
	}
	user := GetUser(c)
	if bid.Want {
		biz.Want(user, bid.Score)
		r.Success("已叫地主并分配身份", c)
		return
	} else {
		biz.NoWant(user)
		r.Success("已选择不叫地主，并传递给下家", c)
	}
}

func Play(c *gin.Context) {
	user := GetUser(c)
	validRound(user)
	//todo
	var cardList []*pojo.Card
	result := biz.PlayCard(user, cardList)
	if result == nil {
		r.Success("success", c)
	} else {
		biz.CountScore(user, result)
		r.Success(result, c)
	}
}

func GamesPass(c *gin.Context) {
	user := GetUser(c)
	validRound(user)
	biz.Pass(user)
	r.Success("success", c)
}

func validRound(user *db.User) {
	result := biz.IsPlayerRound(user)
	if !result {
		panic("当前不是该玩家出牌回合")
	}
}
