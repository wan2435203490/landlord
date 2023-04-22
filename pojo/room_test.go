package pojo

import (
	"fmt"
	"landlord/common/enum"
	"testing"
)

func TestReset(t *testing.T) {
	pl := make([]*Player, 0, 8)
	for i := 0; i < 8; i++ {
		pl = append(pl, &Player{
			Identity: enum.Identity(i),
		})
	}

	r := &Room{
		PlayerList: pl,
	}

	r.Reset()

	fmt.Printf("%#v", r)

}

func TestFor(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}

	for i, i2 := range s {
		fmt.Println(i, &i2)
	}
}
