package ws

import (
	"landlord/common/enum"
)

type readyGame struct {
	Message
	UserId string `json:"userId"`
}

func NewReadyGame(userId string) *readyGame {
	var v readyGame
	v.UserId = userId
	v.Type = v.GetMessageType()
	return &v
}

func (r *readyGame) GetMessageType() string {
	return enum.ReadyGame.GetWsMessageType()
}
