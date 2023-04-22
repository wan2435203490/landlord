package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"landlord/biz"
	r "landlord/common/response"
)

func UpdateUser(c *gin.Context) {
	user := GetUser(c)
	record, err := biz.GetUser(user.Id)
	if err != nil {
		panic(err.Error())
	}
	if record == nil {
		panic("用户信息为空")
	}
	dtoUser := GetDTOUser(c)
	record.UserName = dtoUser.UserName
	record.Password = dtoUser.Password
	record.Avatar = dtoUser.Avatar
	record.Gender = dtoUser.Gender
	session := sessions.Default(c)
	session.Set("curUser", record)
	_ = session.Save()
	biz.UpdateUser(record)
	r.Success(true, c)
}

func Myself(c *gin.Context) {
	user := GetUser(c)
	if user.Id == "" {
		r.ErrorInternal("can't find user", c)
	} else {
		r.Success(user, c)
	}
}
