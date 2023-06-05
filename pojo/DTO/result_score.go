package DTO

import (
	"landlord/common/enum"
	"landlord/common/utils"
	"strconv"
)

type ResultScore struct {
	UserName     string `json:"username"`
	MoneyChange  string `json:"moneyChange"`
	IdentityName string `json:"identityName"`
}

func NewResultScore(userName string, multiple int, isWin, isLandlord bool) *ResultScore {
	if isLandlord {
		return NewLandlordScore(userName, multiple, isWin)
	} else {
		return NewFarmerScore(userName, multiple, isWin)
	}
}

func NewFarmerScore(userName string, multiple int, isWin bool) *ResultScore {
	return &ResultScore{
		UserName:     userName,
		MoneyChange:  GetMoneyChange(multiple, isWin),
		IdentityName: enum.Farmer.GetIdentity(),
	}
}

func NewLandlordScore(userName string, multiple int, isWin bool) *ResultScore {
	return &ResultScore{
		UserName:     userName,
		MoneyChange:  GetMoneyChange(multiple*2, isWin),
		IdentityName: enum.Landlord.GetIdentity(),
	}
}

func GetMoneyChange(moneyChange int, isWin bool) string {
	moneyChange = utils.IfThen(isWin, moneyChange, -moneyChange).(int)
	return strconv.Itoa(moneyChange)
}
