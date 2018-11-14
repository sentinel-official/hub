package types

import (
	"github.com/tendermint/tendermint/crypto"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"time"
)

type VPNDetails struct {
	ApiPort    string
	VpnPort    string
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
	TotalLockedCoins csdkTypes.Coins
	ReleasedCoins    csdkTypes.Coins
	Counter          int64
	Timestamp        time.Time
	VpnPubKey        crypto.PubKey
	CPubKey          crypto.PubKey
	CAddress         csdkTypes.AccAddress
	Status           uint8
	Locked           bool
}

func GetNewSessionMap(coins csdkTypes.Coins, vpnpub crypto.PubKey, cpub crypto.PubKey, caddr csdkTypes.AccAddress, time time.Time) Session {
	return Session{
		TotalLockedCoins: coins,
		ReleasedCoins:    coins.Minus(coins),
		VpnPubKey:        vpnpub,
		CPubKey:          cpub,
		Timestamp:        time,
		CAddress:         caddr,
		Status:           1,
		Locked:           false,
	}

}
