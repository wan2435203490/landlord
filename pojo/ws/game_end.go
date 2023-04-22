package ws

import (
	"landlord/common/enum"
	"landlord/pojo/DTO"
)

type GameEnd struct {
	Message
	WinningIdentity enum.Identity      `json:"winingIdentity"`
	IsWinning       bool               `json:"winning"`
	ResList         []*DTO.ResultScore `json:"resList"`
}

func EmptyGameEnd() *GameEnd {
	v := &GameEnd{}
	v.Type = v.GetMessageType()
	return v
}

func NewGameEnd(winningIdentity enum.Identity, isWinning bool,
	resList []*DTO.ResultScore) *GameEnd {
	v := &GameEnd{
		WinningIdentity: winningIdentity,
		IsWinning:       isWinning,
		ResList:         resList,
	}
	v.Type = v.GetMessageType()
	return v
}

func (g *GameEnd) GetMessageType() string {
	return enum.GameEnd.GetWsMessageType()
}
