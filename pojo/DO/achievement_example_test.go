package DO

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	criteria := &GeneratedCriteria{}
	fmt.Println(criteria.Criteria == nil)
}

func TestAchievementExample_CreateCriteria(t *testing.T) {
	a := &AchievementExample{}
	a.CreateCriteria()
	fmt.Printf("%#v", a)
}
