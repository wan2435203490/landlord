package api

import (
	"github.com/gin-contrib/sessions"
	"landlord/db"
	"landlord/middleware"
)

func (a *Api) User() *db.User {
	value, _ := a.Context.Get(middleware.UserSessionKey)
	return value.(*db.User)
}

// SetUser todo 根据type反射set any type
func (a *Api) SetUser(user *db.User) {
	session := sessions.Default(a.Context)
	session.Set(middleware.UserSessionKey, user)
	_ = session.Save()
}
