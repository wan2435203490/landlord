package api

import (
	"github.com/gin-gonic/gin"
	"landlord/pojo/DTO"
	"landlord/sdk/api"
	"landlord/svc"
)

var UserApi userApi

type userApi struct {
	api.Api
	svc.UserSvc
}

func (a *userApi) Update(c *gin.Context) {
	var dtoUser DTO.User
	if !a.Bind(&dtoUser) {
		return
	}
	user := a.User()

	user.UserName = dtoUser.UserName
	user.Password = dtoUser.Password
	user.Avatar = dtoUser.Avatar
	user.Gender = dtoUser.Gender

	if a.UpdateUser(user) != nil {
		return
	}
	a.SetUser(user)
	a.OK(true)
}

func (a *userApi) Myself(c *gin.Context) {
	user := a.User()
	a.OK(user)
}
