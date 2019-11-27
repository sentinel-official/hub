package types

import (
	"encoding/json"
	"fmt"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	KB    = sdk.NewInt(1000)
	MB    = KB.MulRaw(1000)
	MB500 = MB.MulRaw(500)
	GB    = MB.MulRaw(1000)
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
	return fmt.Sprintf("%d upload, %d download", b.Upload.Int64(), b.Download.Int64())
}

func (b Bandwidth) CeilTo(precision sdk.Int) Bandwidth {
	_b := Bandwidth{
		Upload: precision.Sub(sdk.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Upload.BigInt(), precision.BigInt()))),
		Download: precision.Sub(sdk.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Download.BigInt(), precision.BigInt()))),
	}

	if _b.Upload.Equal(precision) {
		_b.Upload = sdk.NewInt(0)
	}

	if _b.Download.Equal(precision) {
		_b.Download = sdk.NewInt(0)
	}

	return b.Add(_b)
}

func (b Bandwidth) Sum() sdk.Int {
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
	return b.Upload == sdk.Int{} ||
		b.Download == sdk.Int{}
}

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return NewBandwidth(sdk.NewInt(upload), sdk.NewInt(download))
}

type BandwidthSignatureData struct {
	ID        SubscriptionID `json:"id"`
	Index     uint64         `json:"index"`
	Bandwidth Bandwidth      `json:"bandwidth"`
}

func NewBandwidthSignatureData(id SubscriptionID, index uint64, bandwidth Bandwidth) BandwidthSignatureData {
	return BandwidthSignatureData{
		ID:        id,
		Index:     index,
		Bandwidth: bandwidth,
	}
}

func (b BandwidthSignatureData) Bytes() []byte {
	bz, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return bz
}
