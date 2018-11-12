package types

type VPNDetails struct {
	Ip         string
	Port       string
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
