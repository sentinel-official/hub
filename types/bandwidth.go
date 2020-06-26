package types

import (
	"fmt"

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

func (n Bandwidth) IsAnyZero() bool {
	return n.Upload.IsZero() || n.Download.IsZero()
}

func (n Bandwidth) IsAllZero() bool {
	return n.Upload.IsZero() && n.Download.IsZero()
}

func (n Bandwidth) IsAnyNegative() bool {
	return n.Upload.IsNegative() || n.Download.IsNegative()
}

func (n Bandwidth) IsValid() bool {
	return !n.IsAnyNegative() && !n.IsAnyZero()
}

func (n Bandwidth) String() string {
	return fmt.Sprintf("%s↑, %s↓ bytes", n.Upload, n.Download)
}

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return Bandwidth{
		Upload:   sdk.NewInt(upload),
		Download: sdk.NewInt(download),
	}
}
