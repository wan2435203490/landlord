package ws

import "landlord/common/enum"

type bid struct {
	Message
	Score int `json:"score"`
}

func NewBid(score int) *bid {
	//return &bid{&Message{Type: enum.Bid}}
	var v bid
	v.Type = v.GetMessageType()
	v.Score = score
	return &v
}

func (b *bid) GetMessageType() string {
	return enum.Bid.GetWsMessageType()
}
