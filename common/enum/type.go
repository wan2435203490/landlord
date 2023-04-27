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

func (t *Type) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToUpper(s) {
	default:
		*t = Single
	case "单张":
		*t = Single
	case "对子":
		*t = Pair
	case "三张":
		*t = ThreeType
	case "三带一":
		*t = ThreeWithOne
	case "三带一对":
		*t = ThreeWithPair
	case "四带二":
		*t = FourWithTwo
	case "四带两对":
		*t = FourWithFour
	case "炸弹":
		*t = Bomb
	case "王炸":
		*t = JokerBomb
	case "顺子":
		*t = Straight
	case "连对":
		*t = StraightPair
	case "飞机":
		*t = Aircraft
	case "飞机带翅膀":
		*t = AircraftWithSingleWings
	case "飞机带大炮":
		*t = AircraftWithPairWings
	}

	return nil
}

func (t Type) MarshalJSON() ([]byte, error) {
	var s string
	switch t {
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
