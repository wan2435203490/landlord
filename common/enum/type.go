package enum

type Type int

const (
	Single Type = iota
	Pair
	ThreeType
	ThreeWithOne
	ThreeWithPair

	FourWithTwo
	FourWithFour

	Bomb
	JokerBomb

	Straight
	StraightPair

	Aircraft
	AircraftWithSingleWings
	AircraftWithPairWings
)

func (t Type) GetType() string {
	return []string{"单张", "对子", "三张", "三带一", "三带一对", "四带二", "四带两对", "炸弹", "王炸", "顺子",
		"连对", "飞机", "飞机带翅膀", "飞机带大炮"}[t]
}
