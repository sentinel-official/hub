package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"

	hubtypes "github.com/sentinel-official/hub/types"
)

type (
	Subscription interface {
		proto.Message
		Type() SubscriptionType
		Validate() error
		GetID() uint64
		GetAddress() string
		GetExpiryAt() time.Time
		GetStatus() hubtypes.Status
		GetStatusAt() time.Time
		SetExpiryAt(v time.Time)
		SetStatus(v hubtypes.Status)
		SetStatusAt(v time.Time)
	}
	Subscriptions []Subscription
)

var (
	_ Subscription = (*NodeSubscription)(nil)
	_ Subscription = (*PlanSubscription)(nil)
)

func (s *BaseSubscription) SetExpiryAt(v time.Time)     { s.ExpiryAt = v }
func (s *BaseSubscription) SetStatus(v hubtypes.Status) { s.Status = v }
func (s *BaseSubscription) SetStatusAt(v time.Time)     { s.StatusAt = v }

func (s *BaseSubscription) Validate() error {
	if s.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if s.Address == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(s.Address); err != nil {
		return errors.Wrapf(err, "invalid address %s", s.Address)
	}
	if s.ExpiryAt.IsZero() {
		return fmt.Errorf("expiry_at cannot be zero")
	}
	if !s.Status.IsValid() {
		return fmt.Errorf("status must be valid")
	}
	if s.StatusAt.IsZero() {
		return fmt.Errorf("status_at cannot be zero")
	}

	return nil
}

func (s *NodeSubscription) Type() SubscriptionType {
	return TypeNode
}

func (s *NodeSubscription) Validate() error {
	if s.BaseSubscription == nil {
		return fmt.Errorf("base_subscription cannot be nil")
	}
	if err := s.BaseSubscription.Validate(); err != nil {
		return err
	}
	if s.NodeAddress == "" {
		return fmt.Errorf("node_address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(s.NodeAddress); err != nil {
		return errors.Wrapf(err, "invalid node_address %s", s.NodeAddress)
	}
	if s.Gigabytes == 0 && s.Hours == 0 {
		return fmt.Errorf("[gigabytes, hours] cannot be empty")
	}
	if s.Gigabytes != 0 && s.Hours != 0 {
		return fmt.Errorf("[gigabytes, hours] cannot be non-empty")
	}
	if s.Gigabytes != 0 {
		if s.Gigabytes < 0 {
			return fmt.Errorf("gigabytes cannot be negative")
		}
	}
	if s.Hours != 0 {
		if s.Hours < 0 {
			return fmt.Errorf("hours cannot be negative")
		}
	}
	if s.Deposit.Denom != "" {
		if s.Deposit.IsNil() {
			return fmt.Errorf("deposit cannot be nil")
		}
		if s.Deposit.IsNegative() {
			return fmt.Errorf("deposit cannot be negative")
		}
		if s.Deposit.IsZero() {
			return fmt.Errorf("deposit cannot be zero")
		}
		if !s.Deposit.IsValid() {
			return fmt.Errorf("deposit must be valid")
		}
	}

	return nil
}

func (s *NodeSubscription) GetNodeAddress() hubtypes.NodeAddress {
	if s.NodeAddress == "" {
		return nil
	}

	addr, err := hubtypes.NodeAddressFromBech32(s.NodeAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (s *PlanSubscription) Type() SubscriptionType {
	return TypePlan
}

func (s *PlanSubscription) Validate() error {
	if s.BaseSubscription == nil {
		return fmt.Errorf("base_subscription cannot be nil")
	}
	if err := s.BaseSubscription.Validate(); err != nil {
		return err
	}
	if s.PlanID == 0 {
		return fmt.Errorf("plan_id cannot be zero")
	}
	if s.Denom != "" {
		if err := sdk.ValidateDenom(s.Denom); err != nil {
			return errors.Wrapf(err, "invalid denom %s", s.Denom)
		}
	}

	return nil
}
