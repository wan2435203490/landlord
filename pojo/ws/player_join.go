package ws

import (
	"landlord/common/enum"
	"landlord/db"
)

type playerJoin struct {
	Message
	User *db.User `json:"user"`
}

func NewPlayerJoin(user *db.User) *playerJoin {
	var v playerJoin
	v.User = user
	v.Type = v.GetMessageType()
	return &v
}

func (p *playerJoin) GetMessageType() string {
	return enum.PlayerJoin.GetWsMessageType()
}
