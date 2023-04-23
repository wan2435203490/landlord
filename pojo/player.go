package pojo

import (
	"landlord/common/enum"
	"landlord/db"
	"sort"
)

type Player struct {
	Id          int           `json:"id"`
	Identity    enum.Identity `json:"identity"`
	Cards       []*Card       `json:"cards"`
	RecentCards []*Card       `json:"recentCards"`
	User        *db.User      `json:"user"`
	Ready       bool          `json:"ready"`
}

func (p *Player) GetNextPlayerId() int {
	//return utils.IfThen(p.Id == 3, 1, p.Id+1).(int)
	return p.Id%3 + 1
}

func (p *Player) AddCards(cards []*Card) {
	p.Cards = append(p.Cards, cards...)

	sort.SliceStable(p.Cards, func(i, j int) bool {
		return p.Cards[i].Grade > p.Cards[j].Grade
	})
	//sort.Ints(p.Cards)
}

func (p *Player) RemoveCards(cards []*Card) {
	for _, card := range cards {
		for j, old := range p.Cards {
			if old.Equals(card) {
				p.Cards = append(p.Cards[:j], p.Cards[j+1:]...)
				break
			}
		}
	}
}

func (p *Player) ClearRecentCards() {
	p.RecentCards = nil
}

func (p *Player) Reset() {
	p.Cards = nil
	p.Ready = false
	p.Identity = -1
	p.RecentCards = nil
}

func (p *Player) IsLandlord() bool {
	return p.Identity == enum.Landlord
}
