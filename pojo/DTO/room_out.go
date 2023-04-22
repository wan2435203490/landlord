package DTO

import (
	"landlord/common/enum"
	"landlord/pojo"
)

type RoomOut struct {
	Id         string          `json:"id"`
	Title      string          `json:"title"`
	Owner      *UserOut        `json:"owner"`
	PlayerList []*PlayerOut    `json:"playerList"`
	RoomStatus enum.RoomStatus `json:"roomStatus"`
	Status     string          `json:"status"`
	Multiple   int             `json:"multiple"`
	TopCards   []*pojo.Card    `json:"topCards"`
	StepNum    int             `json:"stepNum"`
	CountDown  int             `json:"countdown"`
}

func ToRoomOut(r *pojo.Room) *RoomOut {
	out := &RoomOut{
		Id:         r.Id,
		Title:      r.Title,
		Owner:      ToUserOut(r.Owner),
		PlayerList: ToPlayerOutList(r.PlayerList),
		RoomStatus: r.RoomStatus,
		Status:     r.RoomStatus.GetRoomStatus(),
		Multiple:   r.Multiple,
		StepNum:    r.StepNum,
	}

	if r.Distribution != nil {
		out.TopCards = r.Distribution.TopCards
	}

	return out
}
