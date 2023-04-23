package pojo

import (
	"landlord/common/enum"
)

func GetCard(id int) *Card {
	t := enum.CardType(id % 4)
	n := id / 4
	var g enum.CardGrade

	if n > 12 {
		if t == 0 {
			t = enum.SmallJokerType
			g = enum.Fourteenth
		} else {
			t = enum.BigJokerType
			g = enum.Fifteenth
		}
	}

	card := &Card{
		Type:   t,
		Number: enum.CardNumber(n + 1),
	}

	if g != 0 {
		card.Grade = g
	} else {
		card.Grade = ConvertNum2Grade(n)
	}

	return card
}

func ConvertNum2Grade(n int) enum.CardGrade {
	//n: 0,1,2...13
	//n:0, Grade:A Twelfth(11)
	//n:1, Grade:2 Thirteenth(12)
	//n:2, Grade:3 First(0)
	//(过滤掉n:13了)n:13, Grade:小王 Fourteenth(14) ｜ 大王 Fifteenth(15)
	if n == 13 {
		panic("你咋进来的呢")
	}

	return enum.CardGrade((n + 11) % 13)
}
