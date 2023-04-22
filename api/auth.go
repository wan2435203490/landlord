package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"landlord/biz"
	"landlord/common/response"
	"landlord/common/token"
	"landlord/db"
	"landlord/pojo/DTO"
	"net/http"
)

func Login(c *gin.Context) {
	var login DTO.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		//todo
		//log.NewError("0", utils.GetSelfFuncName(), err.Error())
		r.ErrorInternal(err.Error(), c)
		return
	}

	user, err := biz.GetUserByName(login.UserName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//第一次进来 用户为空 自动插入
			newUser := db.NewUser(login.UserName, login.Password)
			user, err = biz.InsertUser(newUser)
			if err != nil {
				r.ErrorInternal(err.Error(), c)
				return
			}
		} else {
			r.ErrorInternal(err.Error(), c)
			return
		}
	} else if user.Password != login.Password {
		r.ErrorInternal("账号或密码错误", c)
		return
	}

	token := saveSession(c, user)

	r.Success(token, c)

}

func QQLogin(c *gin.Context) {
	//todo
	//https://wiki.connect.qq.com/
}

// qq登录成功回调
func QQCallback(c *gin.Context) {

}

func PermissionDenied(c *gin.Context) {
	r.Error(http.StatusUnauthorized, "请先登录", c)
}

// 保存用户登录对象到 Session，并返回客户端Token令牌
func saveSession(c *gin.Context, user *db.User) string {
	session := sessions.Default(c)
	userJson, err := json.Marshal(&user)
	if err != nil {
		panic(err.Error())
	}
	session.Set("curUser", userJson)
	_ = session.Save()
	createToken, err := token.CreateToken(user.Id)
	if err != nil {
		panic(err)
	}
	return createToken
}
