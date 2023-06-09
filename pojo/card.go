package pojo

import (
	"encoding/json"
	"landlord/common/enum"
)

type Card struct {
	Id     int             `json:"id"`
	Type   enum.CardType   `json:"type"`
	Number enum.CardNumber `json:"number"`
	Grade  enum.CardGrade  `json:"grade"`
}

func (c *Card) CompareTo(c2 *Card) bool {
	_, _ = json.Marshal(c2)
	return c.Grade.GreatThanGrade(c2.Grade)
}

//func GreatThanGrade(i, j int) bool {
//	return p.Cards[i].Grade > p.Cards[j].Grade
//}

func (c *Card) EqualsByGrade(c2 *Card) bool {
	return c.Grade == c2.Grade
}

func (c *Card) Equals(c2 *Card) bool {
	return c.Grade == c2.Grade && c.Type == c2.Type
}
