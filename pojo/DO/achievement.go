package DO

import (
	"github.com/google/uuid"
	"strings"
)

type Achievement struct {
	Id           string `json:"id"`
	WinMatch     int    `json:"winMatch"`
	FailureMatch int    `json:"failureMatch"`
	Sum          int    `json:"sum"`
	UserId       string `json:"userId"`
}

func NewAchievement(userId string) *Achievement {
	return &Achievement{
		Id:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		UserId: userId,
	}
}

func (this *Achievement) IncrWinMatch() {
	this.WinMatch++
	this.Sum++
}

func (this *Achievement) IncrFailureMatch() {
	this.FailureMatch++
	this.Sum++
}
