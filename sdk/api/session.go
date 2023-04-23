package api

import (
	json2 "encoding/json"
	"github.com/gin-contrib/sessions"
	"landlord/db"
)

func (a *Api) DefaultSession() sessions.Session {
	return sessions.Default(a.Context)
}

func (a *Api) GetUserFromSession() *db.User {
	session := a.DefaultSession()
	get := session.Get(UserSessionKey)
	if get == nil {
		a.ErrorMsg("session记录被清除，请退出重新登录。")
		return nil
	}
	var user db.User
	err := json2.Unmarshal(get.([]byte), &user)
	if err != nil {
		a.Error(err)
		return nil
	}

	if &user == nil || user.Id == "" {
		a.ErrorMsg("session记录被清除，请退出重新登录。")
		return nil
	}

	return &user
}

func (a *Api) SetUserToSession(user *db.User) {
	session := a.DefaultSession()
	session.Set(UserSessionKey, user)
	_ = session.Save()
}
