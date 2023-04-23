package component

import (
	"encoding/json"
	"landlord/pojo/ws"
)

var NC NotifyComponent

type NotifyComponent struct {
}

func (nc *NotifyComponent) SendStr2Room(roomId, content string) bool {
	room := RC.GetRoom(roomId)
	ids := room.GetUserIds()
	return WS.Send2Users(ids, content)
}

func (nc *NotifyComponent) Send2Room(roomId string, msg ws.IMessage) bool {
	if bs, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return nc.SendStr2Room(roomId, string(bs))
	}
}

func (nc *NotifyComponent) Send2User(userId string, msg ws.IMessage) bool {
	if bs, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		return WS.Send2User(userId, string(bs))
	}
}

func (nc *NotifyComponent) Send2AllUser(msg ws.IMessage) {
	if bs, err := json.Marshal(msg); err != nil {
		panic(err)
	} else {
		WS.Send2AllUser(string(bs))
	}
}
