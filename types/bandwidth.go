package types

import (
	"fmt"
	"math/big"

	csdk "github.com/cosmos/cosmos-sdk/types"
)

// nolint:gochecknoglobals
var (
	KB    = csdk.NewInt(1000)
	MB    = KB.MulRaw(1000)
	MB500 = MB.MulRaw(500)
	GB    = MB.MulRaw(1000)
)

type Bandwidth struct {
	Upload   csdk.Int `json:"upload"`
	Download csdk.Int `json:"download"`
}

func NewBandwidth(upload, download csdk.Int) Bandwidth {
	return Bandwidth{
		Upload:   upload,
		Download: download,
	}
}

func (b Bandwidth) String() string {
	return fmt.Sprintf("%d upload, %d download", b.Upload.Int64(), b.Download.Int64())
}

func (b Bandwidth) CeilTo(precision csdk.Int) Bandwidth {
	_b := Bandwidth{
		Upload: precision.Sub(csdk.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Upload.BigInt(), precision.BigInt()))),
		Download: precision.Sub(csdk.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Download.BigInt(), precision.BigInt()))),
	}

	if _b.Upload.Equal(precision) {
		_b.Upload = csdk.NewInt(0)
	}
	if _b.Download.Equal(precision) {
		_b.Download = csdk.NewInt(0)
	}

	return b.Add(_b)
}

func (b Bandwidth) Sum() csdk.Int {
	return b.Upload.Add(b.Download)
}

func (b Bandwidth) Add(bandwidth Bandwidth) Bandwidth {
	b.Upload = b.Upload.Add(bandwidth.Upload)
	b.Download = b.Download.Add(bandwidth.Download)

	return b
}

func (b Bandwidth) Sub(bandwidth Bandwidth) Bandwidth {
	b.Upload = b.Upload.Sub(bandwidth.Upload)
	b.Download = b.Download.Sub(bandwidth.Download)

	return b
}

func (b Bandwidth) AllLT(bandwidth Bandwidth) bool {
	return b.Upload.LT(bandwidth.Upload) &&
		b.Download.LT(bandwidth.Download)
}

func (b Bandwidth) AnyLT(bandwidth Bandwidth) bool {
	return b.Upload.LT(bandwidth.Upload) ||
		b.Download.LT(bandwidth.Download)
}

func (b Bandwidth) AllEqual(bandwidth Bandwidth) bool {
	return b.Upload.Equal(bandwidth.Upload) &&
		b.Download.Equal(bandwidth.Download)
}

func (b Bandwidth) AllLTE(bandwidth Bandwidth) bool {
	return b.AllLT(bandwidth) || b.AllEqual(bandwidth)
}

func (b Bandwidth) AnyZero() bool {
	return b.Upload.IsZero() ||
		b.Download.IsZero()
}

func (b Bandwidth) AllPositive() bool {
	return b.Upload.IsPositive() &&
		b.Download.IsPositive()
}

func (b Bandwidth) AnyNegative() bool {
	return b.Upload.IsNegative() ||
		b.Download.IsNegative()
}

func (b Bandwidth) AnyNil() bool {
	return b.Upload == csdk.Int{} ||
		b.Download == csdk.Int{}
}

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return NewBandwidth(csdk.NewInt(upload), csdk.NewInt(download))
}
