package enum

type CardNumber int

const (
	One CardNumber = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten

	Jack
	Lady
	King

	SmallJoker
	BigJoker
)

func (cn CardNumber) GetCardNumber() string {
	return []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "SmallJoker", "BigJoker"}[cn]
}
