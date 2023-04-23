package internal

import (
	"fmt"
	"landlord/common/enum"
	"landlord/pojo"
	"testing"
)

func TestEq(t *testing.T) {

	//for i := 1; i < 18; i++ {
	//	if i != 10 && i != 15 {
	//		println(i, false)
	//	} else {
	//		println(i, true)
	//	}
	//}

	for i := 0; i < 18; i++ {
		if i != 8 && i != 12 && i != 16 {
			println(i, false)
		} else {
			println(i, true)
		}
	}

}

func TestIsAircraftWithSingleWing(t *testing.T) {
	// 333444 56	    3,4     5,6
	// 333444 55	    3,4		5,5
	// 333444555 667  3,4,5   6,6,7
	// 333444555 999  3,4,5  ,9
	// 33334444
	// 333444555666 7777

	var cards []*pojo.Card
	arr0 := []int{3, 3, 4, 4, 4, 5, 6, 3}
	arr1 := []int{3, 3, 4, 4, 5, 4, 5, 3}
	arr2 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 7}
	arr3 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 9, 9, 9}
	arr4 := []int{3, 3, 3, 3, 4, 4, 4, 4}
	arr5 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 7}

	for _, i := range arr0 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr0:", IsAircraftWithSingleWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr1 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr1:", IsAircraftWithSingleWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr2 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr2:", IsAircraftWithSingleWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr3 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr3:", IsAircraftWithSingleWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr4 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr4:", IsAircraftWithSingleWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr5 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr5:", IsAircraftWithSingleWing(cards))

}

func TestIsAircraftWithPairWing(t *testing.T) {
	// 333444555667788
	// 333444555666677
	// 3334445566
	// 3334445555

	var cards []*pojo.Card
	arr0 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 7, 7, 8, 8}
	arr1 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 8, 8}
	arr2 := []int{3, 3, 3, 4, 4, 4, 5, 5, 6, 6}
	arr3 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 5}
	arr4 := []int{3, 3, 3, 3, 4, 4, 4, 4, 5, 5}
	arr5 := []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7}

	for _, i := range arr0 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr0:", IsAircraftWithPairWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr1 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr1:", IsAircraftWithPairWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr2 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr2:", IsAircraftWithPairWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr3 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr3:", IsAircraftWithPairWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr4 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr4:", IsAircraftWithPairWing(cards))

	cards = make([]*pojo.Card, 0)
	for _, i := range arr5 {
		cards = append(cards, &pojo.Card{Grade: enum.CardGrade(i)})
	}
	fmt.Println("arr5:", IsAircraftWithPairWing(cards))

}
