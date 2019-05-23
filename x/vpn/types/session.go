package types

import (
	"fmt"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Session struct {
	ID                  sdkTypes.ID        `json:"id"`
	SubscriptionID      sdkTypes.ID        `json:"subscription_id"`
	Bandwidth           sdkTypes.Bandwidth `json:"bandwidth"`
	CalculatedBandwidth sdkTypes.Bandwidth `json:"calculated_bandwidth"`
	Status              string             `json:"status"`
	StatusModifiedAt    int64              `json:"status_modified_at"`
}

func (s Session) String() string {
	return fmt.Sprintf(`Session
  ID:                   %d
  Subscription ID:      %d
  Bandwidth:            %s
  Client Signature:     %s
  Status:               %s
  Status Modified At:   %d`, s.ID, s.SubscriptionID, s.Bandwidth,
		s.CalculatedBandwidth, s.Status, s.StatusModifiedAt)
}

func (s Session) IsValid() error {
	if s.Bandwidth.AnyNil() {
		return fmt.Errorf("invalid bandwidth")
	}
	if s.CalculatedBandwidth.AnyNil() {
		return fmt.Errorf("invalid calculated bandwidth")
	}
	if s.Status != StatusActive && s.Status != StatusInactive {
		return fmt.Errorf("invalid status")
	}

	return nil
}
