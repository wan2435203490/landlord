package ws

import (
	"landlord/common/enum"
	"landlord/db"
	"landlord/pojo/DTO"
	"time"
)

type chat struct {
	Message
	Sender     *DTO.UserOut `json:"sender"`
	Content    string       `json:"content"`
	TypeId     int          `json:"typeId"`
	Dimension  string       `json:"dimension"`
	CreateTime time.Time    `json:"createTime"`
}

func NewChat(c *DTO.Chat, user *db.User) *chat {
	v := &chat{
		Content:    c.Content,
		TypeId:     c.Type,
		Sender:     DTO.ToUserOut(user),
		Dimension:  c.Dimension,
		CreateTime: time.Now(),
	}
	v.Type = v.GetMessageType()
	return v
}

func (c *chat) GetMessageType() string {
	return enum.Chat.GetWsMessageType()
}
