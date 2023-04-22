package enum

type CardType int

const (
	//♠️
	Spade CardType = iota
	//♥️
	Heart
	//♣️
	Club
	//♦️
	Diamond
	//🃏小王
	SmallJokerType
	//大王
	BigJokerType
)

func (c CardType) GetCardType() string {
	return []string{"黑桃", "红桃", "梅花", "方块", "小王", "大王"}[c]
}
