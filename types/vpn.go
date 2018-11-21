package types

import (
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type NetSpeed struct {
	Upload   int64 `json:"upload"`
	Download int64 `json:"download"`
}

type Location struct {
	Latitude  int64  `json:"latitude"`
	Longitude int64  `json:"longitude"`
	City      string `json:"city"`
	Country   string `json:"country"`
}

type Info struct {
	Status      bool
	BlockHeight int64
}

type VPNDetails struct {
	Address    csdkTypes.AccAddress
	APIPort    string
	Location   Location
	NetSpeed   NetSpeed
	EncMethod  string
	PricePerGB int64
	Version    string
	Info       Info
	LockerID   string
}

type SessionDetails struct {
	VPNID         string
	ClientAddress csdkTypes.AccAddress
	GBToProvide   int64
	PricePerGB    int64
	Status        bool
	Upload        int64
	Download      int64
	StartTime     *time.Time
	EndTime       *time.Time
	Locked        bool
}

type ActiveSessions []string
