package enum

type CardGrade int

const (
	// First 3-k
	First CardGrade = iota
	Second
	Third
	Fourth
	Fifth
	Sixth
	Seventh
	Eighth
	Ninth
	Tenth
	Eleventh

	// Twelfth A 2
	Twelfth
	Thirteenth

	// Fourteenth 大小王
	Fourteenth
	Fifteenth
)

func (g0 CardGrade) CompareGrade(g1 CardGrade) bool {
	return g0 > g1
}
