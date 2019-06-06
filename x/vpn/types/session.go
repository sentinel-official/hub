package types

import (
	"fmt"

	sdk "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Session struct {
	ID               sdk.ID        `json:"id"`
	SubscriptionID   sdk.ID        `json:"subscription_id"`
	Bandwidth        sdk.Bandwidth `json:"bandwidth"`
	Status           string        `json:"status"`
	StatusModifiedAt int64         `json:"status_modified_at"`
}

func (s Session) String() string {
	return fmt.Sprintf(`Session
  ID:                   %s
  Subscription ID:      %s
  Bandwidth:            %s
  Status:               %s
  Status Modified At:   %d`, s.ID, s.SubscriptionID, s.Bandwidth, s.Status, s.StatusModifiedAt)
}

func (s Session) IsValid() error {
	if s.Bandwidth.AnyNil() {
		return fmt.Errorf("invalid bandwidth")
	}
	if s.Status != StatusActive && s.Status != StatusInactive {
		return fmt.Errorf("invalid status")
	}

	return nil
}
