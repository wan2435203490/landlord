package ws

import (
	"landlord/common/enum"
	"landlord/db"
)

type pass struct {
	Message
	User *db.User `json:"user"`
}

func NewPass(user *db.User) *pass {
	var v pass
	v.User = user
	v.Type = v.GetMessageType()
	return &v
}

func (p *pass) GetMessageType() string {
	return enum.PassType.GetWsMessageType()
}
