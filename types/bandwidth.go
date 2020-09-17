package types

import (
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	Kilobyte = sdk.NewInt(1000)
	Megabyte = sdk.NewInt(1000).Mul(Kilobyte)
	Gigabyte = sdk.NewInt(1000).Mul(Megabyte)
)

type Bandwidth struct {
	Upload   sdk.Int `json:"upload"`
	Download sdk.Int `json:"download"`
}

func NewBandwidth(upload, download sdk.Int) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}

func (b Bandwidth) String() string {
	return fmt.Sprintf("%s↑, %s↓ bytes", b.Upload, b.Download)
}

func (b Bandwidth) IsAnyZero() bool {
	return b.Upload.IsZero() || b.Download.IsZero()
}

func (b Bandwidth) IsAllZero() bool {
	return b.Upload.IsZero() && b.Download.IsZero()
}

func (b Bandwidth) IsAnyNegative() bool {
	return b.Upload.IsNegative() || b.Download.IsNegative()
}

func (b Bandwidth) IsValid() bool {
	return !b.IsAnyNegative() && !b.IsAnyZero()
}

func (b Bandwidth) Sum() sdk.Int {
	return b.Upload.Add(b.Download)
}

func (b Bandwidth) Add(v Bandwidth) Bandwidth {
	b.Upload = b.Upload.Add(v.Upload)
	b.Download = b.Download.Add(v.Download)

	return b
}

func (b Bandwidth) Sub(v Bandwidth) Bandwidth {
	b.Upload = b.Upload.Sub(v.Upload)
	b.Download = b.Download.Sub(v.Download)

	return b
}

func (b Bandwidth) IsAllLTE(v Bandwidth) bool {
	return b.Upload.LTE(v.Upload) && b.Download.LTE(v.Download)
}

func (b Bandwidth) IsAnyGT(v Bandwidth) bool {
	return b.Upload.GT(v.Upload) || b.Download.GT(v.Download)
}

func (b Bandwidth) CeilTo(precision sdk.Int) Bandwidth {
	if !precision.IsPositive() {
		return b
	}

	v := NewBandwidth(
		precision.Sub(sdk.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Upload.BigInt(), precision.BigInt()))),
		precision.Sub(sdk.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Download.BigInt(), precision.BigInt()))),
	)

	if v.Upload.Equal(precision) {
		v.Upload = sdk.NewInt(0)
	}
	if v.Download.Equal(precision) {
		v.Download = sdk.NewInt(0)
	}

	return b.Add(v)
}

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return Bandwidth{
		Upload:   sdk.NewInt(upload),
		Download: sdk.NewInt(download),
	}
}
