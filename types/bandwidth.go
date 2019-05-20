package types

import (
	"encoding/json"
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var GB = csdkTypes.NewInt(1000000000)

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

func (b Bandwidth) Add(bandwidth Bandwidth) Bandwidth {
	b.Upload = b.Upload.Add(bandwidth.Upload)
	b.Download = b.Download.Add(bandwidth.Download)

	return b
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

func NewBandwidthFromInt64(upload, download int64) Bandwidth {
	return NewBandwidth(csdkTypes.NewInt(upload), csdkTypes.NewInt(download))
}

type BandwidthSignData struct {
	ID        ID
	Bandwidth Bandwidth
	NodeOwner csdkTypes.AccAddress
	Client    csdkTypes.AccAddress
}

func NewBandwidthSignData(id ID, bandwidth Bandwidth, nodeOwner, client csdkTypes.AccAddress) *BandwidthSignData {
	return &BandwidthSignData{
		ID:        id,
		Bandwidth: bandwidth,
		NodeOwner: nodeOwner,
		Client:    client,
	}
}

func (b BandwidthSignData) Bytes() []byte {
	bz, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return bz
}
