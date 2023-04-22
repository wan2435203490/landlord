package DO

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	OpenId    string    `json:"openid"`
	Gender    string    `json:"gender"`
	Money     float64   `json:"money"`
	Avatar    string    `json:"avatar"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewUser(userId string) *User {
	return &User{
		Id:        strings.ReplaceAll(uuid.NewString(), "-", ""),
		TimeStamp: time.Now(),
		Money:     0.0,
	}
}

func (u *User) IncrMoney(value float64) {
	u.Money += value
}

func (u *User) DescMoney(value float64) {
	u.Money -= value
}
