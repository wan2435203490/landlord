package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"landlord/db"
	"landlord/pojo"
	"landlord/pojo/DTO"
	"landlord/sdk/api"
	"landlord/svc"
)

var GameApi gameApi

type gameApi struct {
	api.Api
	svc.GameSvc
}

func (a *gameApi) Ready(c *gin.Context) {
	user := a.User()
	a.OK(a.ReadyGame(user))
}

func (a *gameApi) UnReady(c *gin.Context) {
	user := a.User()
	a.OK(a.UnReadyGame(user))
}

func (a *gameApi) Bid(c *gin.Context) {
	var bid DTO.Bid
	if !a.Bind(&bid) {
		return
	}
	user := a.User()
	if bid.Want {
		if msg := a.Want(user, bid.Score); msg == "" {
			a.OK("已叫地主并分配身份")
		} else {
			a.ErrorInternal(msg)
		}
	} else {
		a.NoWant(user)
		a.OK("已选择不叫地主，并传递给下家")
	}
}

func (a *gameApi) Play(c *gin.Context) {
	user := a.User()
	if !a.validRound(user) {
		return
	}
	var cardList []*pojo.Card
	if !a.Bind(&cardList, binding.JSON) {
		return
	}

	//todo 逻辑优化 指责不单一
	result, msg := a.PlayCard(user, cardList)
	if msg != "" {
		a.ErrorInternal(msg)
		return
	}

	if result == nil {
		a.OK("success")
	} else {
		AchievementApi.CountScore(user, result)
		a.OK(result)
	}
}

func (a *gameApi) Pass(c *gin.Context) {
	user := a.User()

	if !a.validRound(user) {
		return
	}
	a.PassGame(user)
	a.OK("success")
}

func (a *gameApi) Give(c *gin.Context) {
	user := a.User()

	if !a.validRound(user) {
		return
	}
	cards := a.GiveCards(user)
	a.OK(cards)
}

// validRound valid 需要 set response
func (a *gameApi) validRound(user *db.User) bool {
	result := PlayerApi.IsPlayerRound(user)
	if !result {
		a.ErrorInternal("当前不是该玩家出牌回合")
	}
	return result
}
