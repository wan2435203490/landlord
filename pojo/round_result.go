package pojo

import "landlord/common/enum"

type RoundResult struct {
	WinIdentity enum.Identity `json:"winIdentity"`
	LandlordId  int           `json:"landlord"`
	Multiple    int           `json:"multiple"`
}
