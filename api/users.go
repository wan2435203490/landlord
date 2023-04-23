package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	"landlord/pojo/DTO"
	"landlord/sdk/api"
)

var UserApi userApi

type userApi struct {
	api.Api
	biz.UserSvc
}

func (a *userApi) UpdateUser(c *gin.Context) {
	var dtoUser DTO.User
	if a.Bind(&dtoUser).Err != nil {
		return
	}

	user := a.GetUserFromSession()
	if user == nil {
		return
	}

	user.UserName = dtoUser.UserName
	user.Password = dtoUser.Password
	user.Avatar = dtoUser.Avatar
	user.Gender = dtoUser.Gender

	a.UpdateUserBiz(user)
	a.SetUserToSession(user)
	a.Success(true)
}

func (a *userApi) Myself(c *gin.Context) {

	user := a.GetUserFromSession()
	if user == nil {
		return
	}

	a.Success(user)
}
