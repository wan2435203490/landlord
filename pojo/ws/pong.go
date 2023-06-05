package ws

import (
	"landlord/common/enum"
	"time"
)

type pong struct {
	Message
	TimeStamp time.Time `json:"timeStamp"`
}

func NewPong() *pong {
	var v pong
	v.TimeStamp = time.Now()
	v.Type = v.GetMessageType()
	return &v
}

func (p *pong) GetMessageType() string {
	return enum.Pong.GetWsMessageType()
}
