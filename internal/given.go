package internal

import (
	"landlord/common/enum"
	"landlord/pojo"
	"sort"
)

// GivePlayCards 推荐出牌 preCards 是排序后的
func GivePlayCards(cards, preCards []*pojo.Card) [][]*pojo.Card {
	var res [][]*pojo.Card

	preType := GetCardsType(preCards...)
	n := len(preCards)
	var preGrade enum.CardGrade

	switch preType {
	case enum.Single:
		preGrade = preCards[0].Grade
		res = GiveSingleCards(cards, preGrade)
	case enum.Pair:
		preGrade = preCards[0].Grade
		res = GivePairCards(cards, preGrade)
	case enum.ThreeType:
		preGrade = preCards[0].Grade
		res = GiveThreeCards(cards, preGrade)
	case enum.ThreeWithOne:
		preGrade = preCards[1].Grade
		res = GiveThreeCards(cards, preGrade)
	case enum.ThreeWithPair:
		preGrade = preCards[2].Grade
		res = GiveThreeCards(cards, preGrade)
	case enum.FourWithTwo:
		preGrade = preCards[3].Grade
		res = GiveFourCards(cards, preGrade)
	case enum.FourWithFour:
		_, preGrade = IsFourWithFour(preCards)
		res = GiveFourCards(cards, preGrade)
	case enum.Bomb:
		preGrade = preCards[0].Grade
		res = GiveFourCards(cards, preGrade)
	case enum.JokerBomb:
	case enum.Straight:
		preGrade = preCards[0].Grade
		res = GiveStraightCards(cards, preGrade, n)
	case enum.StraightPair:
		preGrade = preCards[0].Grade
		res = GiveStraightPairCards(cards, preGrade, n)
	case enum.Aircraft:
		preGrade = preCards[0].Grade
		res = GiveAircraft(cards, preGrade, n)
	case enum.AircraftWithSingleWings:
		_, preGrade = IsAircraftWithSingleWing(cards)
		res = GiveAircraft(cards, preGrade, n)
	case enum.AircraftWithPairWings:
		_, preGrade = IsAircraftWithPairWing(cards)
		res = GiveAircraft(cards, preGrade, n)
	}

	return res
}

func GetCanPlayCards(cards []*pojo.Card, preType enum.Type) {
	switch preType {
	case enum.Single:
	case enum.Pair:
	case enum.ThreeType:
	case enum.ThreeWithOne:
	case enum.ThreeWithPair:
	case enum.FourWithTwo:
	case enum.FourWithFour:
	case enum.Bomb:
	case enum.JokerBomb:
	case enum.Straight:
	case enum.StraightPair:
	case enum.Aircraft:
	case enum.AircraftWithSingleWings:
	case enum.AircraftWithPairWings:

	}
}

//func GetAllTypeCards(cards )

// GiveSingleCards 获取cards中大于card的牌组合
func GiveSingleCards(cards []*pojo.Card, grade enum.CardGrade) [][]*pojo.Card {
	sortCards(cards)

	var res [][]*pojo.Card
	dupGrades := make(map[int]bool)

	//单张 炸弹 王炸
	for _, c := range cards {

		if c.Grade.GreatThanGrade(grade) {
			//单张推荐出牌去重
			if !dupGrades[int(c.Grade)] {
				dupGrades[int(c.Grade)] = true
				res = append(res, []*pojo.Card{c})
			}
		}
	}

	if bbs := findAllBombs(cards); bbs != nil {
		res = append(res, bbs...)
	}

	return res
}

// GivePairCards 获取cards中大于card的牌组合
func GivePairCards(cards []*pojo.Card, grade enum.CardGrade) [][]*pojo.Card {
	if len(cards) < 2 {
		return nil
	}
	sortCards(cards)

	var res [][]*pojo.Card

	for i := 0; i+1 < len(cards); i++ {
		if IsPair(cards[i : i+2]) {
			i++
			if i > 2 && IsPair(cards[i-2:i]) {
				//有炸弹时会进入这里
				continue
			}
			if cards[i].Grade.GreatThanGrade(grade) {
				res = append(res, []*pojo.Card{cards[i-1], cards[i]})
			}
		}
	}

	if bbs := findAllBombs(cards); bbs != nil {
		res = append(res, bbs...)
	}

	return res
}

// GiveThreeCards 三张 三带一 三带二 都用这个 不提示带的牌
func GiveThreeCards(cards []*pojo.Card, grade enum.CardGrade) [][]*pojo.Card {
	if len(cards) < 3 {
		return nil
	}
	sortCards(cards)

	var res [][]*pojo.Card

	for i := 0; i+2 < len(cards); i++ {
		if IsThree(cards[i : i+3]) {
			if cards[i].Grade.GreatThanGrade(grade) {
				res = append(res, []*pojo.Card{cards[i], cards[i+1], cards[i+2]})
			}
			i += 2
		}
	}

	if bbs := findAllBombs(cards); bbs != nil {
		res = append(res, bbs...)
	}

	return res
}

// GiveFourCards 炸弹 四带二 四带两对 都用这个 不提示带的牌
func GiveFourCards(cards []*pojo.Card, grade enum.CardGrade) [][]*pojo.Card {
	if len(cards) < 4 {
		return nil
	}
	sortCards(cards)

	var res [][]*pojo.Card

	if bombs := GiveBombs(cards, grade); bombs != nil {
		res = append(res, bombs...)
	}
	if bomb := findJokerBomb(cards); bomb != nil {
		res = append(res, bomb)
	}

	return res
}

// GiveStraightCards 顺子 grade为顺子最小值 n为长度
func GiveStraightCards(cards []*pojo.Card, grade enum.CardGrade, n int) [][]*pojo.Card {
	if len(cards) < 5 {
		return nil
	}
	sortCards(cards)

	var res [][]*pojo.Card

	cardsLen := len(cards)

	for i := 0; i+n-1 < cardsLen; i++ {
		//过滤重复的顺子
		if res != nil && cards[i].Grade == res[len(res)-1][0].Grade {
			continue
		}
		if cards[i].Grade.GreatThanGrade(grade) {
			cur := []*pojo.Card{cards[i]}
			//从cards[i]开始寻找一个顺子
			ii := i
		out:
			//获取长度为n的顺子 顺子里的牌不能大于A
			for j := 1; j < n && ii+1 < cardsLen && cards[ii+1].Grade <= enum.Twelfth; {
				switch cards[ii+1].Grade - cur[len(cur)-1].Grade {
				case 0:
					//连续两张牌同级 遍历下一张牌
					ii++
				case 1:
					//连续两张牌差一级 加入顺子 遍历下一张牌
					cur = append(cur, cards[ii+1])
					ii++
					j++
				default:
					//连续两张牌差两级及以上 说明没有顺子
					break out
				}

			}
			if len(cur) == n {
				res = append(res, cur)
				//有顺子ii回到最初到index
				ii = i
			}
			//ii之前都没顺子
			i = ii
		}
	}

	if bbs := findAllBombs(cards); bbs != nil {
		res = append(res, bbs...)
	}

	return res
}

// GiveStraightPairCards 连对334455 grade为连对最小值3 n为长度 6
func GiveStraightPairCards(cards []*pojo.Card, grade enum.CardGrade, n int) [][]*pojo.Card {
	if len(cards) < 6 {
		return nil
	}

	sortCards(cards)

	var res [][]*pojo.Card

	cardsLen := len(cards)

	for i := 0; i+n-1 < cardsLen; i++ {
		//过滤重复的连对
		if res != nil && cards[i].Grade == res[len(res)-1][0].Grade {
			continue
		}
		if !IsPair(cards[i : i+2]) {
			continue
		}
		if cards[i].Grade.GreatThanGrade(grade) {
			cur := []*pojo.Card{cards[i], cards[i+1]}
			//从cards[i]开始寻找一个连对
			ii := i
		out:
			//获取长度为n的连对 连对里的牌不能大于A
			//校验下一对存不存在 ii+2和ii+3
			for j := 1; j < n/2 && ii+3 < cardsLen && cards[ii+3].Grade <= enum.Twelfth; {
				if !IsPair(cards[ii+2 : ii+4]) {
					ii++
					continue
				}
				switch cards[ii+2].Grade - cur[len(cur)-1].Grade {
				case 0:
					//连续两对同级 遍历下两张牌
					ii += 2
				case 1:
					//连续两对差一级 加入连对 遍历下一张牌
					cur = append(cur, cards[ii+2], cards[ii+3])
					ii += 2
					j++
				default:
					//连续两对差两级及以上 说明没有连对
					break out
				}

			}
			if len(cur) == n {
				res = append(res, cur)
				//有连对ii回到最初到index
				ii = i
			}
			//ii之前都没连对
			i = ii
		}
	}

	if bbs := findAllBombs(cards); bbs != nil {
		res = append(res, bbs...)
	}

	return res
}

// GiveAircraft 飞机333444555 grade为连对最小值3 n为长度 9
func GiveAircraft(cards []*pojo.Card, grade enum.CardGrade, n int) [][]*pojo.Card {
	if len(cards) < 6 {
		return nil
	}
	sortCards(cards)

	var res [][]*pojo.Card

	cardsLen := len(cards)

	for i := 0; i+n-1 < cardsLen; i++ {
		if !IsThree(cards[i : i+3]) {
			continue
		}
		//过滤重复的顺子
		if res != nil && cards[i].Grade == res[len(res)-1][0].Grade {
			continue
		}
		if cards[i].Grade.GreatThanGrade(grade) {
			cur := []*pojo.Card{cards[i], cards[i+1], cards[i+2]}
			//从cards[i]开始寻找一个飞机
			ii := i
		out:
			//获取总长度为n的飞机 飞机里的牌不能大于A
			//校验下一个三张存不存在 ii+2，ii+3和ii+4
			for j := 1; j < n/3 && ii+5 < cardsLen && cards[ii+5].Grade <= enum.Twelfth; {
				if !IsThree(cards[ii+3 : ii+6]) {
					ii++
					continue
				}
				switch cards[ii+3].Grade - cur[len(cur)-1].Grade {
				case 1:
					//连续两个三张差一级 加入飞机 遍历下一张牌
					cur = append(cur, cards[ii+3], cards[ii+4], cards[ii+5])
					ii += 3
					j++
				default:
					//连续两个三张差两级及以上 说明没有飞机
					break out
				}

			}
			if len(cur) == n {
				res = append(res, cur)
				//有飞机ii回到最初到index
				ii = i
			}
			//ii之前都没飞机
			i = ii
		}
	}

	if bbs := findAllBombs(cards); bbs != nil {
		res = append(res, bbs...)
	}

	return res
}

// 从小到大
func sortCards(cards []*pojo.Card) {
	sort.Slice(cards, func(i, j int) bool {
		return cards[j].Grade.GreatThanGrade(cards[i].Grade)
	})
}

// GiveBombs 查找炸弹  cards需已排序
func GiveBombs(cards []*pojo.Card, grade enum.CardGrade) [][]*pojo.Card {
	var bombs [][]*pojo.Card
	n := len(cards) - 3

	for i := 0; i < n; i++ {
		if cards[i].Grade.GreatThanGrade(grade) && IsBomb(cards[i:i+4]) {
			bombs = append(bombs, []*pojo.Card{cards[i], cards[i+1], cards[i+2], cards[i+3]})
			i += 3
		}
	}
	return bombs
}

// findJokerBomb 查找王炸 cards已排序
func findJokerBomb(cards []*pojo.Card) []*pojo.Card {
	n := len(cards) - 1
	if cards[n-1].Grade == enum.Fourteenth {
		return []*pojo.Card{cards[n-1], cards[n]}
	}
	return nil
}

func findAllBombs(cards []*pojo.Card) [][]*pojo.Card {
	var bb [][]*pojo.Card
	if bombs := GiveBombs(cards, -1); bombs != nil {
		bb = append(bb, bombs...)
	}
	if bomb := findJokerBomb(cards); bomb != nil {
		bb = append(bb, bomb)
	}
	return bb
}
