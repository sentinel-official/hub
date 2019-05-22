package types

import (
	"encoding/json"
	"fmt"
	"math/big"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

// nolint:gochecknoglobals
var (
	GB = csdkTypes.NewInt(1000000000)
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

func (b Bandwidth) String() string {
	return fmt.Sprintf("%d upload, %d download", b.Upload.Int64(), b.Download.Int64())
}

func (b Bandwidth) CeilTo(precision csdkTypes.Int) Bandwidth {
	_b := Bandwidth{
		Upload: precision.Sub(csdkTypes.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Upload.BigInt(), precision.BigInt()))),
		Download: precision.Sub(csdkTypes.NewIntFromBigInt(
			big.NewInt(0).Rem(b.Download.BigInt(), precision.BigInt()))),
	}

	if _b.Upload.Equal(precision) {
		_b.Upload = csdkTypes.NewInt(0)
	}
	if _b.Download.Equal(precision) {
		_b.Download = csdkTypes.NewInt(0)
	}

	return b.Add(_b)
}

func (b Bandwidth) Sum() csdkTypes.Int {
	return b.Upload.Add(b.Download)
}

func (b Bandwidth) Add(bandwidth Bandwidth) Bandwidth {
	b.Upload = b.Upload.Add(bandwidth.Upload)
	b.Download = b.Download.Add(bandwidth.Download)

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
	return b.Upload == csdkTypes.Int{} ||
		b.Download == csdkTypes.Int{}
}

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return NewBandwidth(csdkTypes.NewInt(upload), csdkTypes.NewInt(download))
}

type BandwidthSignData struct {
	SubscriptionID uint64
	SessionIndex   uint64
	Bandwidth      Bandwidth
	NodeOwner      csdkTypes.AccAddress
	Client         csdkTypes.AccAddress
}

func NewBandwidthSignData(subscriptionID, sessionIndex uint64, bandwidth Bandwidth,
	nodeOwner, client csdkTypes.AccAddress) BandwidthSignData {

	return BandwidthSignData{
		SubscriptionID: subscriptionID,
		SessionIndex:   sessionIndex,
		Bandwidth:      bandwidth,
		NodeOwner:      nodeOwner,
		Client:         client,
	}
}

func (b BandwidthSignData) Bytes() []byte {
	bz, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return bz
}
