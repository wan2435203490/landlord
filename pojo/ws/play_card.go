package ws

import (
	"landlord/common/enum"
	"landlord/db"
	"landlord/pojo"
	"sort"
)

type playCard struct {
	Message
	User     *db.User        `json:"user"`
	CardList []*pojo.Card    `json:"cardList"`
	CardType enum.Type       `json:"cardType"`
	Number   enum.CardNumber `json:"number"`
}

func NewPlayCard(user *db.User, cardList []*pojo.Card, cardType enum.Type) *playCard {
	card := &playCard{
		User:     user,
		CardType: cardType,
	}

	card.Type = card.GetMessageType()

	sort.Slice(cardList, func(i, j int) bool {
		return cardList[i].Grade > cardList[j].Grade
	})

	if cardType == enum.Single || cardType == enum.Pair {
		card.Number = cardList[0].Number
	}

	return card
}

func (p *playCard) GetMessageType() string {
	return enum.PlayCard.GetWsMessageType()
}
