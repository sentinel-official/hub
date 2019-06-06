package types

import (
	"encoding/json"
	"fmt"

	csdk "github.com/cosmos/cosmos-sdk/types"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Subscription struct {
	ID                 sdk.ID          `json:"id"`
	NodeID             sdk.ID          `json:"node_id"`
	Client             csdk.AccAddress `json:"client"`
	PricePerGB         csdk.Coin       `json:"price_per_gb"`
	TotalDeposit       csdk.Coin       `json:"total_deposit"`
	RemainingDeposit   csdk.Coin       `json:"remaining_deposit"`
	RemainingBandwidth sdk.Bandwidth   `json:"remaining_bandwidth"`
	Status             string          `json:"status"`
	StatusModifiedAt   int64           `json:"status_modified_at"`
}

func (s Subscription) TotalBandwidth() sdk.Bandwidth {
	x := s.TotalDeposit.Amount.
		Mul(sdk.MB500).
		Quo(s.PricePerGB.Amount)

	return sdk.NewBandwidth(x, x)
}

func (s Subscription) String() string {
	return fmt.Sprintf(`Subscription
  ID:                  %s
  Node ID:             %s
  Client Address:      %s
  Price Per GB:        %s
  Total Deposit:       %s
  Total Bandwidth:     %s
  Remaining Deposit:   %s
  Remaining Bandwidth: %s
  Status:              %s
  Status Modified At:  %d`, s.ID, s.NodeID, s.Client,
		s.PricePerGB, s.TotalDeposit, s.TotalBandwidth(),
		s.RemainingDeposit, s.RemainingBandwidth, s.Status, s.StatusModifiedAt)
}

// nolint: gocyclo
func (s Subscription) IsValid() error {
	if s.Client == nil || s.Client.Empty() {
		return fmt.Errorf("invalid client")
	}
	if s.PricePerGB.Denom == "" || s.PricePerGB.IsZero() {
		return fmt.Errorf("invalid price per gb")
	}
	if s.TotalDeposit.Denom != s.PricePerGB.Denom || s.TotalDeposit.IsZero() {
		return fmt.Errorf("invalid total deposit")
	}
	if s.RemainingDeposit.Denom != s.TotalDeposit.Denom || s.TotalDeposit.IsLT(s.RemainingDeposit) {
		return fmt.Errorf("invalid remaining deposit")
	}
	if s.RemainingBandwidth.AnyNil() || s.TotalBandwidth().AnyLT(s.RemainingBandwidth) {
		return fmt.Errorf("invalid total remaining bandwidth")
	}
	if s.Status != StatusActive && s.Status != StatusInactive {
		return fmt.Errorf("invalid status")
	}

	return nil
}

type BandwidthSignatureData struct {
	ID        sdk.ID        `json:"id"`
	Index     uint64        `json:"index"`
	Bandwidth sdk.Bandwidth `json:"bandwidth"`
}

func NewBandwidthSignatureData(id sdk.ID, index uint64, bandwidth sdk.Bandwidth) BandwidthSignatureData {
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
