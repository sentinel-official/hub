package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Subscription struct {
	ID                  sdkTypes.ID          `json:"id"`
	NodeID              sdkTypes.ID          `json:"node_id"`
	Client              csdkTypes.AccAddress `json:"client"`
	PricePerGB          csdkTypes.Coin       `json:"price_per_gb"`
	TotalDeposit        csdkTypes.Coin       `json:"total_deposit"`
	TotalBandwidth      sdkTypes.Bandwidth   `json:"total_bandwidth"`
	ConsumedDeposit     csdkTypes.Coin       `json:"consumed_deposit"`
	ConsumedBandwidth   sdkTypes.Bandwidth   `json:"consumed_bandwidth"`
	CalculatedBandwidth sdkTypes.Bandwidth   `json:"calculated_bandwidth"`
	Status              string               `json:"status"`
	StatusModifiedAt    int64                `json:"status_modified_at"`
}

func (s Subscription) String() string {
	return fmt.Sprintf(`Subscription
  ID:                   %d
  NodeID:               %d
  Client Address:       %s
  Price Per GB:         %s
  Total Deposit:        %s
  Total Bandwidth:      %s
  Consumed Deposit:     %s
  Consumed Bandwidth:   %s
  Calculated Bandwidth: %s
  Status:               %s
  Status Modified At:   %d`, s.ID, s.NodeID, s.Client,
		s.PricePerGB, s.TotalDeposit, s.TotalBandwidth, s.ConsumedDeposit, s.ConsumedBandwidth,
		s.CalculatedBandwidth, s.Status, s.StatusModifiedAt)
}

// nolint: gocyclo
func (s Subscription) IsValid() error {
	if s.Client == nil || s.Client.Empty() {
		return fmt.Errorf("invalid client")
	}
	if len(s.PricePerGB.Denom) == 0 || s.PricePerGB.IsZero() {
		return fmt.Errorf("invalid price per gb")
	}
	if len(s.TotalDeposit.Denom) == 0 || s.TotalDeposit.IsZero() {
		return fmt.Errorf("invalid total deposit")
	}
	if s.TotalBandwidth.AnyNil() || !s.TotalBandwidth.AllPositive() {
		return fmt.Errorf("invalid total bandwidth")
	}
	if len(s.ConsumedDeposit.Denom) == 0 || s.TotalDeposit.IsLT(s.ConsumedDeposit) {
		return fmt.Errorf("invalid consumed deposit")
	}
	if s.ConsumedBandwidth.AnyNil() || s.TotalBandwidth.AnyLT(s.ConsumedBandwidth) {
		return fmt.Errorf("invalid total consumed bandwidth")
	}
	if s.CalculatedBandwidth.AnyNil() || s.TotalBandwidth.AnyLT(s.CalculatedBandwidth) {
		return fmt.Errorf("invalid total calculated bandwidth")
	}
	if s.Status != StatusActive && s.Status != StatusInactive {
		return fmt.Errorf("invalid status")
	}

	return nil
}
