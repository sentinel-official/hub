package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var GB = csdkTypes.NewInt(1000000000)

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

func (b Bandwidth) IsZero() bool {
	return b.Upload.IsZero() ||
		b.Download.IsZero()
}

func (b Bandwidth) IsPositive() bool {
	return b.Upload.IsPositive() &&
		b.Download.IsPositive()
}

func (b Bandwidth) IsNegative() bool {
	return b.Upload.IsNegative() ||
		b.Download.IsNegative()
}

func (b Bandwidth) IsNil() bool {
	return b.Upload == csdkTypes.Int{} ||
		b.Download == csdkTypes.Int{}
}

func NewBandwidth(upload, download csdkTypes.Int) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}
