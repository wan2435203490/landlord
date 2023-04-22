package ws

import "landlord/common/enum"

type bidEnd struct {
	Message
}

func NewBidEnd() *bidEnd {
	var v bidEnd
	v.Type = v.GetMessageType()
	return &v
}

func (b *bidEnd) GetMessageType() string {
	return enum.BidEnd.GetWsMessageType()
}
