package ws

type Message struct {
	//messageType
	Type string `json:"type"`
}

type IMessage interface {
	GetMessageType() string
}
