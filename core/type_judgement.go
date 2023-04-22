package core

import (
	"landlord/common/enum"
	"landlord/pojo"
	"sort"
)

var IllegalGradesOfStraight []enum.CardGrade

type TypeJudgement struct {
}

func init() {
	IllegalGradesOfStraight = []enum.CardGrade{enum.Thirteenth, enum.Fourteenth, enum.Fifteenth}
}

func IsSingle(cards []*pojo.Card) bool {
	return len(cards) == 1
}

func IsPair(cards []*pojo.Card) bool {
	if len(cards) == 2 {
		return isAllGradeEqual(cards)
	}
	return false
}

func isAllGradeEqual(cards []*pojo.Card) bool {
	card0 := cards[0]
	for _, card := range cards {
		if !card.EqualsByGrade(card0) {
			return false
		}
	}
	return true
}

func IsThree(cards []*pojo.Card) bool {
	if len(cards) == 3 {
		return isAllGradeEqual(cards)
	}
	return false
}

func IsBomb(cards []*pojo.Card) bool {
	if len(cards) == 4 {
		return isAllGradeEqual(cards)
	}
	return false
}

func IsJokerBomb(cards []*pojo.Card) bool {
	return len(cards) == 2 &&
		(cards[0].Grade == enum.Fourteenth && cards[1].Grade == enum.Fifteenth ||
			cards[1].Grade == enum.Fourteenth && cards[0].Grade == enum.Fifteenth)
}

func IsThreeWithOne(cards []*pojo.Card) bool {
	if len(cards) != 4 {
		return false
	}
	if isAllGradeEqual(cards) {
		return false
	}

	SortCards(cards)

	return EqualsGrade(cards, 0, 1, 2) || EqualsGrade(cards, 1, 2, 3)
}

func IsThreeWithPair(cards []*pojo.Card) bool {
	if len(cards) != 5 {
		return false
	}
	SortCards(cards)

	return EqualsGrade(cards, 0, 1) && EqualsGrade(cards, 2, 3, 4) ||
		EqualsGrade(cards, 0, 1, 2) && EqualsGrade(cards, 3, 4)
}

func IsFourWithTwo(cards []*pojo.Card) bool {
	if len(cards) != 6 {
		return false
	}

	SortCards(cards)

	//不能带王炸
	if cards[5].Grade == enum.Fourteenth {
		return false
	}

	return EqualsGrade(cards, 0, 1, 2, 3) || EqualsGrade(cards, 1, 2, 3, 4) ||
		EqualsGrade(cards, 2, 3, 4, 5) || EqualsGrade(cards, 3, 4, 5, 6)
}

// 返回compare的index
func IsFourWithFour(cards []*pojo.Card) (bool, enum.CardGrade) {
	if len(cards) != 8 {
		return false, -1
	}

	SortCards(cards)

	//不能带王炸
	if cards[7].Grade == enum.Fourteenth {
		return false, -1
	}

	if EqualsGrade(cards, 0, 1, 2, 3) {
		return true, cards[0].Grade
	} else if EqualsGrade(cards, 1, 2, 3, 4) {
		return true, cards[1].Grade
	} else if EqualsGrade(cards, 2, 3, 4, 5) {
		return true, cards[2].Grade
	} else if EqualsGrade(cards, 3, 4, 5, 6) {
		return true, cards[3].Grade
	}

	return false, -1
}

// 是否顺子
func IsStraight(cards []*pojo.Card) bool {
	n := len(cards)
	if n < 5 {
		return false
	}
	SortCards(cards)
	//最后一张牌大于A
	if cards[n-1].Grade > enum.Twelfth {
		return false
	}

	for i := 0; i < n-1; i++ {
		if cards[i].Grade+1 != cards[i+1].Grade {
			return false
		}
	}

	return true
}

// 是否连对
func IsStraightPair(cards []*pojo.Card) bool {
	n := len(cards)
	//最低三连对
	if n < 6 || n%2 != 0 {
		return false
	}
	SortCards(cards)
	//判断是不是两个一样大的顺子
	var cards1, cards2 []*pojo.Card
	for i, card := range cards {
		if i%2 == 0 {
			cards1 = append(cards1, card)
		} else {
			cards2 = append(cards2, card)
		}
	}

	if !IsStraight(cards1) || !IsStraight(cards2) {
		return false
	}

	if cards1[0].EqualsByGrade(cards2[0]) {
		return true
	}

	return false
}

// 是否飞机
func IsAircraft(cards []*pojo.Card) bool {
	n := len(cards)
	//最低双飞
	if n < 6 || n%3 != 0 {
		return false
	}
	SortCards(cards)

	for i := 0; i < n; i += 3 {

		if !EqualsGrade(cards, i, i+1, i+2) {
			return false
		}

		if i+3 < n && cards[i].Grade+1 != cards[i+3].Grade {
			return false
		}
	}

	return true
}

// 333444555667788
// 333444555666677
// 3334445566
// 3334445555
// 第二个返回值是grade
func IsAircraftWithPairWing(cards []*pojo.Card) (bool, enum.CardGrade) {
	n := len(cards)
	//最低双飞带两对 32*2 10, 32*3 15
	if n != 10 && n != 15 {
		return false, -1
	}
	SortCards(cards)

	//分离飞机和翅膀
	allGrades := make(map[enum.CardGrade]int)

	// airGrades: 3,4,5  wingGrades: 6,7
	var airGrades, wingGrades []enum.CardGrade

	for _, card := range cards {
		allGrades[card.Grade] += 1
	}

	//333444555667788
	//[3,3] [4,3] [5,3] [6,2] [7,2] [8,2]
	//air: 3,4,5 wing: 6,7,8
	for g, c := range allGrades {
		switch c {
		case 4:
			//带4个的情况 加2个翅膀
			wingGrades = append(wingGrades, g)
			fallthrough
		case 2:
			wingGrades = append(wingGrades, g)
		case 3:
			airGrades = append(airGrades, g)
		default:
			return false, -1
		}
	}

	na := len(airGrades)
	if na < 2 {
		return false, -1
	}

	//解析air的牌型
	if isContinuous, _ := ResolveAirGrades(airGrades); !isContinuous {
		//不连续
		return false, -1
	}

	//飞机和翅膀长度要相等
	if len(airGrades) != len(wingGrades) {
		return false, -1
	}

	return true, airGrades[0]
}

// IsAircraftWithSingleWing
// 333444 56	    3,4     5,6
// 333444 55	    3,4		5,5
// 333444555 667  3,4,5   6,6,7
// 333444555 999  3,4,5  ,9
// 33334444
// 333444555666 7777
func IsAircraftWithSingleWing(cards []*pojo.Card) (bool, enum.CardGrade) {
	n := len(cards)
	//31*2 8,31*3 12, 31*4 16
	if n != 8 && n != 12 && n != 16 {
		return false, -1
	}
	SortCards(cards)

	//分离飞机和翅膀
	allGrades := make(map[enum.CardGrade]int)

	// airGrades: 3,4,5(去重)  wingGrades: 6,6,7(不去重)
	var airGrades, wingGrades []enum.CardGrade

	for _, card := range cards {
		if allGrades[card.Grade] == 3 {
			wingGrades = append(wingGrades, card.Grade)
		} else {
			allGrades[card.Grade] += 1
		}
	}

	//333444555667 这种情况会同时进入case 1 2
	//[3,3] [4,3] [5,3] [6,2] [7,1]
	//air: 3,4,5 wing: 6,7
	for g, c := range allGrades {
		switch c {
		case 3:
			airGrades = append(airGrades, g)
		case 2:
			wingGrades = append(wingGrades, g)
			fallthrough
		case 1:
			wingGrades = append(wingGrades, g)
		default:
			//不允许33334444这种情况
			return false, -1
		}
	}

	if len(airGrades) < 2 {
		return false, -1
	}

	//解析air的牌型
	if isContinuous, grade := ResolveAirGrades(airGrades); !isContinuous {
		//不连续
		return false, -1
	} else {
		if grade != -1 {
			//不是单纯的连续飞机 是3457这种情况 将grade: 7插入到wingGrades
			wingGrades = append(wingGrades, grade, grade, grade)
			for i, g := range airGrades {
				if g == grade {
					airGrades = append(airGrades[:i], airGrades[i+1:]...)
				}
			}
		}
	}

	//飞机和翅膀长度要相等
	if len(airGrades) != len(wingGrades) && len(airGrades)-len(wingGrades) != 4 {
		return false, -1
	}

	return true, airGrades[0]
}

// 3457 3567 345
// 返回是否连续 上面三种情况都算连续 同时返回形如3457中的7
func ResolveAirGrades(grades []enum.CardGrade) (bool, enum.CardGrade) {
	n := len(grades)
	if n == 0 {
		return false, -1
	}
	//3457 3567 345
	sort.Slice(grades, func(i, j int) bool {
		return grades[i] < grades[j]
	})

	flag := false
	var wingGrade enum.CardGrade
	for i := 0; i < n-1; i++ {
		if grades[i]+1 != grades[i+1] {
			//i: 2, i+1: 3
			//i: 0, i+1: 1
			if flag {
				//第二次不连续 只能是3 5 7 9
				return false, -1
			}
			if i == 0 {
				//第一个grade不连续时 可能是3 567，也可能是3 5 7 9
				flag = true
				//记录不连续的3
				wingGrade = grades[0]
			} else if i == n-2 {
				//最后一个不连续 345 7
				return true, grades[n-1]
			}

		}
	}

	if flag {
		return true, wingGrade
	}

	return true, -1
}

func EqualsGrade(cards []*pojo.Card, ii ...int) bool {
	card0 := cards[0]
	for _, i := range ii {
		if !card0.EqualsByGrade(cards[i]) {
			return false
		}
	}
	return true
}
