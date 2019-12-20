package types

import (
	"fmt"
	
	hub "github.com/sentinel-official/hub/types"
)

type Session struct {
	ID               hub.SessionID      `json:"id"`
	SubscriptionID   hub.SubscriptionID `json:"subscription_id"`
	Bandwidth        hub.Bandwidth      `json:"bandwidth"`
	Status           string             `json:"status"`
	StatusModifiedAt int64              `json:"status_modified_at"`
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
	if s.Status != StatusRegistered && s.Status != StatusDeRegistered {
		return fmt.Errorf("invalid status")
	}
	
	return nil
}
