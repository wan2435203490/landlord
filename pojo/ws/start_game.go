package ws

import (
	"landlord/common/enum"
)

type startGame struct {
	Message
	RoomId string `json:"roomId"`
}

func NewStartGame(roomId string) *startGame {
	var v startGame
	v.Type = v.GetMessageType()
	v.RoomId = roomId
	return &v
}

func (s *startGame) GetMessageType() string {
	return enum.StartGame.GetWsMessageType()
}
