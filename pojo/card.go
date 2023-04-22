package pojo

import "landlord/common/enum"

type Card struct {
	Id     int             `json:"id"`
	Type   enum.CardType   `json:"type"`
	Number enum.CardNumber `json:"number"`
	Grade  enum.CardGrade  `json:"grade"`
}

func (c *Card) CompareTo(c2 *Card) bool {
	return c.Grade.CompareGrade(c2.Grade)
}

//func CompareGrade(i, j int) bool {
//	return p.Cards[i].Grade > p.Cards[j].Grade
//}

func (c *Card) EqualsByGrade(c2 *Card) bool {
	return c.Grade == c2.Grade
}

func (c *Card) Equals(c2 *Card) bool {
	return c.Number == c2.Number
}
