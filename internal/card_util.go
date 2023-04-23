package internal

import (
	"landlord/common/enum"
	"landlord/pojo"
	"sort"
)

// 抽象type 获取一个可比较的值出来
func GetCardsType(cards []*pojo.Card) enum.Type {
	if IsSingle(cards) {
		return enum.Single
	}
	if IsPair(cards) {
		return enum.Pair
	}
	if IsThree(cards) {
		return enum.ThreeType
	}
	if IsThreeWithOne(cards) {
		return enum.ThreeWithOne
	}
	if IsThreeWithPair(cards) {
		return enum.ThreeWithPair
	}
	if IsStraight(cards) {
		return enum.Straight
	}
	if IsStraightPair(cards) {
		return enum.StraightPair
	}
	if IsFourWithTwo(cards) {
		return enum.FourWithTwo
	}
	if b, _ := IsFourWithFour(cards); b {
		return enum.FourWithFour
	}
	if IsBomb(cards) {
		return enum.Bomb
	}
	if IsJokerBomb(cards) {
		return enum.JokerBomb
	}
	if IsAircraft(cards) {
		return enum.Aircraft
	}
	if b, _ := IsAircraftWithSingleWing(cards); b {
		return enum.AircraftWithSingleWings
	}
	if b, _ := IsAircraftWithPairWing(cards); b {
		return enum.AircraftWithPairWings
	}
	return -1

}

// 对牌进行从小到大地排序 花色？
func SortCards(cards []*pojo.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Grade < cards[j].Grade
	})
}
