package DTO

import (
	"landlord/db"
)

type UserOut struct {
	Id       string  `json:"id"`
	UserName string  `json:"username"`
	Gender   string  `json:"gender"`
	Avatar   string  `json:"avatar"`
	Money    float64 `json:"money"`
}

func ToUserOut(u *db.User) *UserOut {
	if u == nil {
		return nil
	}

	return &UserOut{
		Id:       u.Id,
		UserName: u.UserName,
		Avatar:   u.Avatar,
		Gender:   u.Gender,
		Money:    u.Money,
	}
}

func ToUserOutList(users []*db.User) []*UserOut {
	if users == nil {
		return nil
	}

	var ret []*UserOut

	for _, u := range users {
		user := ToUserOut(u)
		ret = append(ret, user)
	}

	return ret
}
