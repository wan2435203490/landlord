package ws

import (
	"landlord/common/enum"
)

type pleasePlayCard struct {
	Message
}

func NewPleasePlayCard() *pleasePlayCard {
	var v pleasePlayCard
	v.Type = v.GetMessageType()
	return &v
}

func (p *pleasePlayCard) GetMessageType() string {
	return enum.PleasePlayCard.GetWsMessageType()
}
