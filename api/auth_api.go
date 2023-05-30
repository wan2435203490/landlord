package api

import (
	"github.com/gin-gonic/gin"
	"landlord/common/response"
	"landlord/common/token"
	"landlord/db"
	"landlord/pojo/DTO"
	"landlord/sdk/api"
	"landlord/svc"
	"net/http"
)

var AuthApi authApi

type authApi struct {
	api.Api
}

func (a *authApi) Login(c *gin.Context) {

	var login DTO.Login
	if !a.Bind(&login) {
		return
	}

	user := &db.User{UserName: login.UserName, Password: login.Password}
	var s svc.UserSvc
	if a.IsError(s.GetOrInsertUser(user)) {
		return
	}

	if user.Password != login.Password {
		a.ErrorInternal("账号或密码错误")
		return
	}

	a.SetUser(user)
	tokenStr, err := token.CreateToken(user.Id)
	if err != nil {
		a.ErrorInternal("token创建失败:" + err.Error())
		return
	}

	a.OK(tokenStr)
}

func (a *authApi) QQLogin(c *gin.Context) {
	//todo
	//https://wiki.connect.qq.com/
}

// qq登录成功回调
func (a *authApi) QQCallback(c *gin.Context) {

}

func (a *authApi) PermissionDenied(c *gin.Context) {
	r.Error(http.StatusUnauthorized, "请先登录", c)
}
