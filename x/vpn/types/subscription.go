package types

import (
	"fmt"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Subscription struct {
	ID                 sdkTypes.ID          `json:"id"`
	NodeID             sdkTypes.ID          `json:"node_id"`
	Client             csdkTypes.AccAddress `json:"client"`
	PricePerGB         csdkTypes.Coin       `json:"price_per_gb"`
	TotalDeposit       csdkTypes.Coin       `json:"total_deposit"`
	RemainingDeposit   csdkTypes.Coin       `json:"remaining_deposit"`
	RemainingBandwidth sdkTypes.Bandwidth   `json:"remaining_bandwidth"`
	Status             string               `json:"status"`
	StatusModifiedAt   int64                `json:"status_modified_at"`
}

func (s Subscription) TotalBandwidth() sdkTypes.Bandwidth {
	x := s.TotalDeposit.Amount.
		Mul(sdkTypes.MB500).
		Quo(s.PricePerGB.Amount)

	return sdkTypes.NewBandwidth(x, x)
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
