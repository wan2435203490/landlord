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
	if !a.Bind(&chat) {
		return
	}

	user := a.User()

	if !CheckLimit(user) {
		a.ErrorInternal("你说话太快啦～")
		return
	}

	msg := ws.NewChat(&chat, user)

	dimensionType := enum.ToDimensionType(chat.Dimension)

	switch dimensionType {
	case enum.Room:
		room := component.RC.GetUserRoom(user.Id)
		res := component.NC.Send2Room(room.Id, msg)
		a.OK(res)
	case enum.All:
		msg := component.NC.Send2AllUser(msg)
		if msg != "" {
			a.ErrorInternal(msg)
		} else {
			a.OK(true)
		}
	default:
		a.ErrorInternal(fmt.Sprintf("不支持的聊天范围:%s", chat.Dimension))
	}

}

// CheckLimit 限制短时间消息刷屏
func CheckLimit(user *db.User) bool {
	limiter, ok := rateLimiterMap.Load(user.Id)
	if !ok {
		//每秒往桶中放r个令牌, 桶的容量b
		limiter = rate.NewLimiter(1, 4)
		rateLimiterMap.Store(user.Id, limiter)
	}

	return limiter.(*rate.Limiter).Allow()
}
