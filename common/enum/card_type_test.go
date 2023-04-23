package enum

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	var s []CardType
	s = []CardType{Spade, Diamond, Heart, Club, SmallJokerType, BigJokerType, Spade}

	bs, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))

	var out []CardType
	err = json.Unmarshal(bs, &out)
	if err != nil {
		panic(err)
	}
}
