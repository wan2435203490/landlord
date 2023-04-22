package DTO

import (
	"landlord/common/enum"
	"strconv"
)

type ResultScore struct {
	UserName     string `json:"username"`
	MoneyChange  string `json:"moneyChange"`
	IdentityName string `json:"identityName"`
}

func NewResultScore(userName string, moneyChange int, identity enum.Identity) *ResultScore {
	return &ResultScore{UserName: userName, MoneyChange: strconv.Itoa(moneyChange),
		IdentityName: identity.GetIdentity()}
}
