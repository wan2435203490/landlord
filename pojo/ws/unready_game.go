package ws

import "landlord/common/enum"

type unReadyGame struct {
	Message
	UserId string `json:"userId"`
}

func NewUnReadyGame(userId string) *unReadyGame {
	var v unReadyGame
	v.UserId = userId
	v.Type = v.GetMessageType()
	return &v
}

func (u *unReadyGame) GetMessageType() string {
	return enum.UnReadyGame.GetWsMessageType()
}
