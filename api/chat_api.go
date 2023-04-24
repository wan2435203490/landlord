package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"landlord/common/enum"
	"landlord/db"
	"landlord/internal/component"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
	"landlord/sdk/api"
	"sync"
)

var ChatApi chatApi

type chatApi struct {
	api.Api
}

var rateLimiterMap sync.Map

func (a *chatApi) Chat(c *gin.Context) {

	var chat DTO.Chat
	if a.Bind(&chat) != nil {
		return
	}

	user := a.User()

	if !CheckLimit(user) {
		a.ErrorInternal("你说话太快啦～")
		return
	}

	msg := ws.NewChat(&chat, user)

	switch chat.Dimension {
	case enum.Room.DimensionType():
		room := component.RC.GetUserRoom(user.Id)
		res := component.NC.Send2Room(room.Id, msg)
		a.OK(res)
	case enum.All.DimensionType():
		component.NC.Send2AllUser(msg)
		a.OK(true)
	default:
		a.ErrorInternal(fmt.Sprintf("不支持的聊天范围:%s", chat.Dimension))
	}

}

func CheckLimit(user *db.User) bool {
	if limiter, ok := rateLimiterMap.Load(user.Id); ok {
		return limiter.(*rate.Limiter).Allow()
	} else {
		//每秒往桶中放r个令牌, 桶的容量b
		limiter = rate.NewLimiter(.5, 4)
		rateLimiterMap.Store(user.Id, limiter)
		return limiter.(*rate.Limiter).Allow()
	}
}
