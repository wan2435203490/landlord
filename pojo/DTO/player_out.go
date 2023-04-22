package DTO

import (
	"landlord/common/enum"
	"landlord/pojo"
)

type PlayerOut struct {
	Id          int           `json:"id"`
	CardSize    int           `json:"cardSize"`
	Identity    enum.Identity `json:"identity"`
	RecentCards []*pojo.Card  `json:"recentCards"`
	Ready       bool          `json:"ready"`
	Online      bool          `json:"online"`
	User        *UserOut      `json:"user"`
}

func (p *PlayerOut) GetIdentityName() string {
	//todo 如何表示空？ 取一个default值？
	return p.Identity.GetIdentity()
}

func ToPlayerOut(p *pojo.Player) *PlayerOut {
	if p == nil {
		return nil
	}

	return &PlayerOut{
		Id:          p.Id,
		CardSize:    len(p.Cards),
		Identity:    p.Identity,
		Ready:       p.Ready,
		User:        ToUserOut(p.User),
		RecentCards: p.RecentCards,
	}
}

func ToPlayerOutList(players []*pojo.Player) []*PlayerOut {
	if players == nil {
		return nil
	}

	var ret []*PlayerOut

	for _, player := range players {
		ret = append(ret, ToPlayerOut(player))
	}

	return ret
}
