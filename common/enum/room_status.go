package enum

type RoomStatus int

const (
	Preparing RoomStatus = iota
	Playing
)

func (status RoomStatus) GetRoomStatusName() string {
	return []string{"准备中", "游戏中"}[status]
}

func (status RoomStatus) GetRoomStatus() string {
	return []string{"PREPARING", "PLAYING"}[status]
}
