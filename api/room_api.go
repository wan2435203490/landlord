package api

import (
	"github.com/gin-gonic/gin"
	"landlord/internal/component"
	"landlord/pojo/DTO"
	"landlord/sdk/api"
	"landlord/svc"
)

var RoomApi roomApi

type roomApi struct {
	api.Api
	svc.RoomSvc
}

func (a *roomApi) Rooms(c *gin.Context) {
	rooms := component.RC.ListRooms()
	if len(rooms) != 0 {
		var roomsOut []*DTO.RoomListOutput
		for _, room := range rooms {
			roomsOut = append(roomsOut, DTO.RoomListOutputFromRoom(room))
		}
		a.OK(roomsOut)
		return
	}
	a.OK(rooms)
}

func (a *roomApi) GetById(c *gin.Context) {

	roomId := a.Param("id")
	if roomId == "" {
		return
	}
	user := a.User()
	outRoom := a.GetRoomOut(user, roomId)
	a.OK(outRoom)
}

func (a *roomApi) Create(c *gin.Context) {

	var createRoom DTO.CreateRoom
	if a.Bind(&createRoom) != nil {
		return
	}
	if len(createRoom.Title) == 0 {
		a.ErrorInternal("房间名称不能为空")
		return
	}
	user := a.User()
	room := a.CreateRoom(user, createRoom.Title, createRoom.Password)
	a.OK(room)
}

func (a *roomApi) Join(c *gin.Context) {
	var room DTO.Room
	if a.Bind(&room) != nil {
		return
	}
	user := a.User()
	msg := a.JoinRoom(user, &room)
	a.OK(msg)
}

func (a *roomApi) Exit(c *gin.Context) {
	user := a.User()
	dissolved := a.ExitRoom(user)
	a.OK(dissolved)
}
