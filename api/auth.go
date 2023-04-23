package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	"landlord/common/response"
	"landlord/common/token"
	"landlord/db"
	"landlord/pojo/DTO"
	"landlord/sdk/api"
	"net/http"
)

var AuthApi authApi

type authApi struct {
	api.Api
	biz.UserSvc
}

func (a authApi) Login(c *gin.Context) {
	var login DTO.Login
	if a.Bind(&login).Err != nil {
		return
	}

	user := &db.User{UserName: login.UserName, Password: login.Password}
	if !a.GetOrInsertUser(user) {
		return
	}

	if user.Password != login.Password {
		a.ErrorMsg("账号或密码错误")
		return
	}

	a.SetUserToSession(user)
	tokenStr, err := token.CreateToken(user.Id)
	if err != nil {
		a.ErrorMsg("token创建失败:" + err.Error())
		return
	}

	a.Success(tokenStr)
}

func (a authApi) QQLogin(c *gin.Context) {
	//todo
	//https://wiki.connect.qq.com/
}

// qq登录成功回调
func (a authApi) QQCallback(c *gin.Context) {

}

func (a authApi) PermissionDenied(c *gin.Context) {
	r.Error(http.StatusUnauthorized, "请先登录", c)
}
