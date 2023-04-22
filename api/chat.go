package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"landlord/common/enum"
	r "landlord/common/response"
	"landlord/core/component"
	"landlord/db"
	"landlord/pojo/DTO"
	"landlord/pojo/ws"
	"sync"
)

// var cl ConcurrentLimiter
//
//	type ConcurrentLimiter struct {
//		mu             sync.RWMutex
//		rateLimiterMap map[string]*rate.Limiter
//	}
var rateLimiterMap sync.Map

func Chat(c *gin.Context) {

	//ShouldBindJson只能用一次
	//如果上面使用了 第二次使用不会Bind成功
	var chat DTO.Chat
	if err := c.ShouldBindJSON(&chat); err != nil {
		r.ErrorInternal("bind chat error", c)
		return
	}
	user := GetUser(c)

	if !CheckLimit(user) {
		r.ErrorInternal("你说话太快啦～", c)
		return
	}

	msg := ws.NewChat(&chat, user)

	switch chat.Dimension {
	case enum.Room.DimensionType():
		room := component.RC.GetUserRoom(user.Id)
		msg.Dimension = enum.Room.DimensionType()
		res := component.NC.Send2Room(room.Id, msg)
		r.Success(res, c)
	case enum.All.DimensionType():
		component.NC.Send2AllUser(msg)
		r.Success(true, c)
	default:
		r.ErrorInternal(fmt.Sprintf("不支持的聊天范围:%s", chat.Dimension), c)
	}

}

func CheckLimit(user *db.User) bool {
	//var limiter *rate.Limiter
	if limiter, ok := rateLimiterMap.Load(user.Id); ok {
		return limiter.(*rate.Limiter).Allow()
	} else {
		//每秒往桶中放r个令牌, 桶的容量b
		limiter = rate.NewLimiter(.5, 4)
		rateLimiterMap.Store(user.Id, limiter)
		return limiter.(*rate.Limiter).Allow()
	}
}

//func CheckLimit(user *db.User) bool {
//
//	cl.mu.RLock()
//	limiter := cl.rateLimiterMap[user.Id]
//	if limiter == nil {
//		cl.mu.RUnlock()
//		cl.mu.Lock()
//		//每秒往桶中放r个令牌, 桶的容量b
//		limiter = rate.NewLimiter(.5, 4)
//		cl.rateLimiterMap[user.Id] = limiter
//		cl.mu.Unlock()
//	}
//	cl.mu.RUnlock()
//
//	return limiter.Allow()
//}
