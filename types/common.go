package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Bandwidth struct {
	Upload   csdkTypes.Int `json:"upload"`
	Download csdkTypes.Int `json:"download"`
}

func (b Bandwidth) LT(bandwidth Bandwidth) bool {
	return b.Upload.LT(bandwidth.Upload) &&
		b.Download.LT(bandwidth.Download)
}

func (b Bandwidth) Equal(bandwidth Bandwidth) bool {
	return b.Upload.Equal(bandwidth.Upload) &&
		b.Download.Equal(bandwidth.Download)
}

func (b Bandwidth) LTE(bandwidth Bandwidth) bool {
	return b.LT(bandwidth) || b.Equal(bandwidth)
}

func NewBandwidth(upload, download csdkTypes.Int) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}
