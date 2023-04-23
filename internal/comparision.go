package internal

import (
	"landlord/common/enum"
	"landlord/pojo"
)

func CanPlayCards(cards, preCards []*pojo.Card, typ, preType enum.Type) bool {
	if cards == nil || preCards == nil ||
		typ == -1 || preType == -1 ||
		preType == enum.JokerBomb {
		return false
	}

	if typ == enum.JokerBomb ||
		preType != enum.Bomb && typ == enum.Bomb {
		return true
	}
	//当前出的不是炸弹 牌型不一样不能出
	if preType != typ {
		return false
	}

	switch typ {
	case enum.ThreeWithOne:
		return cards[1].CompareTo(preCards[1])
	case enum.ThreeWithPair:
		return cards[2].CompareTo(preCards[2])
	case enum.FourWithTwo:
		return cards[3].CompareTo(preCards[3])
	case enum.FourWithFour:
		_, i0 := IsFourWithFour(cards)
		_, i1 := IsFourWithFour(preCards)
		return i0.CompareGrade(i1)
	case enum.Straight:
		fallthrough
	case enum.StraightPair:
		if len(cards) != len(preCards) {
			return false
		}
		fallthrough
	case enum.Single:
		fallthrough
	case enum.Pair:
		fallthrough
	case enum.ThreeType:
		fallthrough
	case enum.Aircraft:
		return cards[0].CompareTo(preCards[0])
	case enum.AircraftWithSingleWings:
		_, i0 := IsAircraftWithSingleWing(cards)
		_, i1 := IsAircraftWithSingleWing(preCards)
		return i0.CompareGrade(i1)
	case enum.AircraftWithPairWings:
		_, i0 := IsAircraftWithPairWing(cards)
		_, i1 := IsAircraftWithPairWing(preCards)
		return i0.CompareGrade(i1)
	}

	return false
}

// HasPlayCards 判断是否有可以出的牌
func HasPlayCards(cards, preCards []*pojo.Card, typ, preType enum.CardType) bool {
	return false
}
