package types

import sdkTypes "github.com/cosmos/cosmos-sdk/types"

type Registervpn struct {
	Ip         string
	Port       string
	NetSpeed   NetSpeed
	PricePerGb int64
	EncMethod  string
	Location   Location
	NodeType   string
	Status     bool
	Version    string
}

type NetSpeed struct {
	UploadSpeed   int64
	DownloadSpeed int64
}
type Location struct {
	Latitude  int64
	Longitude int64
	City      string
	Country   string
}

type VpnIBCPacket struct {
	VpnId     string
	Address   sdkTypes.AccAddress
	Coin      sdkTypes.Coin
	DestChain string
}
