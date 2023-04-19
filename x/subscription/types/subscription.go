package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	hubtypes "github.com/sentinel-official/hub/types"
)

type (
	Subscription interface {
		Type() SubscriptionType
		Validate() error
	}
	Subscriptions []Subscription
)

var (
	_ Subscription = (*NodeSubscription)(nil)
	_ Subscription = (*PlanSubscription)(nil)
)

func (s *BaseSubscription) GetAccountAddress() sdk.AccAddress {
	if s.AccountAddress == "" {
		return nil
	}

	addr, err := sdk.AccAddressFromBech32(s.AccountAddress)
	if err != nil {
		panic(err)
	}

	return addr
}

func (s *BaseSubscription) Validate() error {
	if s.ID == 0 {
		return fmt.Errorf("id cannot be zero")
	}
	if s.AccountAddress == "" {
		return fmt.Errorf("account_address cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(s.AccountAddress); err != nil {
		return errors.Wrapf(err, "invalid account_address %s", s.AccountAddress)
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
	if err := s.BaseSubscription.Validate(); err != nil {
		return err
	}
	if s.NodeAddress == "" {
		return fmt.Errorf("node_address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(s.NodeAddress); err != nil {
		return errors.Wrapf(err, "invalid node_address %s", s.NodeAddress)
	}
	if s.Hours < 0 {
		return fmt.Errorf("hours cannot be negative")
	}
	if s.Hours == 0 {
		return fmt.Errorf("hours cannot be zero")
	}
	if s.Price.Denom != "" {
		if s.Price.IsNegative() {
			return fmt.Errorf("price cannot be negative")
		}
		if s.Price.IsZero() {
			return fmt.Errorf("price cannot be zero")
		}
		if !s.Price.IsValid() {
			return fmt.Errorf("price must be valid")
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
