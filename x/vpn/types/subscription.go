package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
)

type Subscription struct {
	ID                 hub.ID         `json:"id"`
	NodeID             hub.ID         `json:"node_id"`
	Client             sdk.AccAddress `json:"client"`
	PricePerGB         sdk.Coin       `json:"price_per_gb"`
	TotalDeposit       sdk.Coin       `json:"total_deposit"`
	RemainingDeposit   sdk.Coin       `json:"remaining_deposit"`
	RemainingBandwidth hub.Bandwidth  `json:"remaining_bandwidth"`
	Status             string         `json:"status"`
	StatusModifiedAt   int64          `json:"status_modified_at"`
}

func (s Subscription) TotalBandwidth() hub.Bandwidth {
	x := s.TotalDeposit.Amount.
		Mul(hub.MB500).
		Quo(s.PricePerGB.Amount)

	return hub.NewBandwidth(x, x)
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
	ID        hub.ID        `json:"id"`
	Index     uint64        `json:"index"`
	Bandwidth hub.Bandwidth `json:"bandwidth"`
}

func NewBandwidthSignatureData(id hub.ID, index uint64, bandwidth hub.Bandwidth) BandwidthSignatureData {
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
