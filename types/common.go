package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Bandwidth struct {
	Upload   csdkTypes.Int `json:"upload"`
	Download csdkTypes.Int `json:"download"`
}

func NewBandwidth(upload, download csdkTypes.Int) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}
