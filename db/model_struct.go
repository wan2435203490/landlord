package db

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type Achievement struct {
	Id           string `gorm:"column:id;primary_key;type:varchar(36)" json:"id"`
	WinMatch     int    `gorm:"column:win_match" json:"winMatch"`
	FailureMatch int    `gorm:"column:failure_match" json:"failureMatch"`
	Sum          int    `gorm:"column:sum" json:"sum"`
	UserId       string `gorm:"column:user_id" json:"userId"`
}

func NewAchievement(userId string) *Achievement {
	return &Achievement{
		Id:     strings.ReplaceAll(uuid.NewString(), "-", ""),
		UserId: userId,
	}
}

func (*Achievement) TableName() string {
	return "achievement"
}

func (a *Achievement) IncrWinMatch() {
	a.WinMatch++
	a.Sum++
}

func (a *Achievement) IncrFailureMatch() {
	a.FailureMatch++
	a.Sum++
}

type User struct {
	Id        string    `gorm:"column:id;primary_key;type:varchar(36)" json:"id"`
	UserName  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	OpenId    string    `gorm:"column:openid" json:"openid"`
	Gender    string    `gorm:"column:gender" json:"gender"`
	Money     float64   `gorm:"column:money" json:"money"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Timestamp time.Time `gorm:"column:timestamp" json:"timestamp"`
}

func NewUser(userName, password string) *User {
	return &User{
		Id:       strings.ReplaceAll(uuid.NewString(), "-", ""),
		UserName: userName,
		Password: password,
	}
}

func (User) TableName() string {
	return "user"
}
