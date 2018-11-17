package types

import (
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type VPNDetails struct {
	ApiPort    string
	VPNPort    string
	Pubkey     crypto.PubKey
	Address    csdkTypes.AccAddress
	NetSpeed   NetSpeed
	PricePerGb int64
	EncMethod  string
	Location   Location
	Version    string
	Info       Info
}

type NetSpeed struct {
	Upload   int64
	Download int64
}

type Location struct {
	Latitude  int64
	Longitude int64
	City      string
	Country   string
}

type Info struct {
	Status      bool
	BlockHeight int64
}

type Session struct {
	VPNID        string
	ClienAddress csdkTypes.AccAddress
	GbToProvide  int64
	PricePerGb   int64
	Status       bool
	Upload       int64
	Download     int64
	StartTime    *time.Time
	EndTime      *time.Time
	Locked       bool
}

func GetNewSession(vpnID string, clientAddress csdkTypes.AccAddress, gbToProvide int64, pricePerGb int64) Session {
	return Session{
		VPNID:        vpnID,
		ClienAddress: clientAddress,
		GbToProvide:  gbToProvide,
		PricePerGb:   pricePerGb,
		Status:       false,
		Locked:       false,
	}

}

type ActiveSessions []string
