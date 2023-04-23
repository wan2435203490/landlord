package api

import (
	"github.com/gin-gonic/gin"
	"landlord/biz"
	r "landlord/common/response"
	"landlord/internal/component"
	"landlord/pojo/DTO"
)

func Rooms(c *gin.Context) {
	rooms := component.RC.ListRooms()
	if len(rooms) != 0 {
		var roomsOut []*DTO.RoomListOutput
		for _, room := range rooms {
			roomsOut = append(roomsOut, DTO.RoomListOutputFromRoom(room))
		}
		r.Success(roomsOut, c)
		return
	}
	r.Success(rooms, c)
}

func GetRoomById(c *gin.Context) {
	roomId := c.Param("id")
	user := GetUser(c)
	outRoom := biz.GetRoomOut(user, roomId)
	r.Success(outRoom, c)
}

func CreateRoom(c *gin.Context) {

	var createRoom DTO.CreateRoom
	if err := c.ShouldBindJSON(&createRoom); err != nil {
		panic(err.Error())
	}
	if len(createRoom.Title) == 0 {
		panic("房间名称不能为空")
	}
	user := GetUser(c)
	room := biz.CreateRoom(user, createRoom.Title, createRoom.Password)
	r.Success(room, c)
}

func Join(c *gin.Context) {
	var room DTO.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		panic(err.Error())
	}
	user := GetUser(c)
	msg := biz.JoinRoom(user, &room)
	r.Success(msg, c)
}

func Exit(c *gin.Context) {
	user := GetUser(c)
	dissolved := biz.ExitRoom(user)
	r.Success(dissolved, c)
}
