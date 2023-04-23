package enum

import (
	"encoding/json"
	"strings"
)

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

func (c *Type) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToUpper(s) {
	default:
		*c = Single
	case "单张":
		*c = Single
	case "对子":
		*c = Pair
	case "三张":
		*c = ThreeType
	case "三带一":
		*c = ThreeWithOne
	case "三带一对":
		*c = ThreeWithPair
	case "四带二":
		*c = FourWithTwo
	case "四带两对":
		*c = FourWithFour
	case "炸弹":
		*c = Bomb
	case "王炸":
		*c = JokerBomb
	case "顺子":
		*c = Straight
	case "连对":
		*c = StraightPair
	case "飞机":
		*c = Aircraft
	case "飞机带翅膀":
		*c = AircraftWithSingleWings
	case "飞机带大炮":
		*c = AircraftWithPairWings
	}

	return nil
}

func (c Type) MarshalJSON() ([]byte, error) {
	var s string
	switch c {
	default:
		s = "单张"
	case Single:
		s = "单张"
	case Pair:
		s = "对子"
	case ThreeType:
		s = "三张"
	case ThreeWithOne:
		s = "三带一"
	case ThreeWithPair:
		s = "三带一对"
	case FourWithTwo:
		s = "四带二"
	case FourWithFour:
		s = "四带两对"
	case Bomb:
		s = "炸弹"
	case JokerBomb:
		s = "王炸"
	case Straight:
		s = "顺子"
	case StraightPair:
		s = "连对"
	case Aircraft:
		s = "飞机"
	case AircraftWithSingleWings:
		s = "飞机带翅膀"
	case AircraftWithPairWings:
		s = "飞机带大炮"
	}

	return json.Marshal(s)
}
