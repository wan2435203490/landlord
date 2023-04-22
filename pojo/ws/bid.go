package ws

import "landlord/common/enum"

type bid struct {
	Message
}

func NewBid() *bid {
	//return &bid{&Message{Type: enum.Bid}}
	var v bid
	v.Type = v.GetMessageType()
	return &v
}

func (b *bid) GetMessageType() string {
	return enum.Bid.GetWsMessageType()
}
