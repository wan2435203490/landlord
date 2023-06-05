package component

import (
	"encoding/json"
	"landlord/pojo/ws"
)

var NC NotifyComponent

type NotifyComponent struct {
}

func (nc *NotifyComponent) SendStr2Room(roomId, content string) string {
	room, msg := RC.GetRoom(roomId)
	if msg != "" {
		return msg
	}
	ids := room.GetUserIds()
	return WS.Send2Users(ids, content)
}

func (nc *NotifyComponent) Send2Room(roomId string, msg ws.IMessage) string {
	if bs, err := json.Marshal(msg); err != nil {
		return err.Error()
	} else {
		return nc.SendStr2Room(roomId, string(bs))
	}
}

func (nc *NotifyComponent) Send2User(userId string, msg ws.IMessage) string {
	if bs, err := json.Marshal(msg); err != nil {
		return err.Error()
	} else {
		return WS.Send2User(userId, string(bs))
	}
}

func (nc *NotifyComponent) Send2AllUser(msg ws.IMessage) string {
	if bs, err := json.Marshal(msg); err != nil {
		return err.Error()
	} else {
		return WS.Send2AllUser(string(bs))
	}
}
