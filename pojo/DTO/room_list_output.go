package DTO

import (
	"landlord/common/enum"
	"landlord/pojo"
)

type RoomListOutput struct {
	Id         string          `json:"id"`
	Title      string          `json:"title"`
	Owner      *UserOut        `json:"owner"`
	Locked     bool            `json:"locked"`
	UserList   []*UserOut      `json:"userList"`
	RoomStatus enum.RoomStatus `json:"roomStatus"`
	Status     string          `json:"status"`
}

func RoomListOutputFromRoom(room *pojo.Room) *RoomListOutput {
	return &RoomListOutput{
		Id:         room.Id,
		Title:      room.Title,
		Owner:      ToUserOut(room.Owner),
		Locked:     room.Locked,
		UserList:   ToUserOutList(room.UserList),
		RoomStatus: room.RoomStatus,
		//json enum问题
		Status: room.RoomStatus.GetRoomStatus(),
	}
}
