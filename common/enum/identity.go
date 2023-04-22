package enum

type Identity int

const (
	Landlord Identity = iota
	Farmer
)

func (i Identity) GetIdentity() string {
	if i < 0 {
		return ""
	}
	return []string{"地主", "农民"}[i]
}
