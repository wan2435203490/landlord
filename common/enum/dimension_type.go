package enum

type DimensionType int

const (
	All DimensionType = iota
	Room
)

func (d DimensionType) DimensionType() string {
	return [2]string{"ALL", "ROOM"}[d]
}
