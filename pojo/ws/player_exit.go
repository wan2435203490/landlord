package ws

import (
	"landlord/common/enum"
	"landlord/db"
)

type playerExit struct {
	Message
	User *db.User `json:"user"`
}

func NewPlayerExit(user *db.User) *playerExit {
	var v playerExit
	v.User = user
	v.Type = v.GetMessageType()
	return &v
}

func (p *playerExit) GetMessageType() string {
	return enum.PlayerExit.GetWsMessageType()
}
