package types

import (
	"encoding/json"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgUpsertRequest)(nil)
)

func NewMsgUpsertRequest(from string, id uint64, address string, duration time.Duration, bandwidth hub.Bandwidth) MsgUpsertRequest {
	return MsgUpsertRequest{
		From:      from,
		Id:        id,
		Address:   address,
		Duration:  duration,
		Bandwidth: bandwidth,
		Signature: signature,
	}
}

func (m MsgUpsertRequest) Route() string {
	return RouterKey
}

func (m MsgUpsertRequest) Type() string {
	return fmt.Sprintf("%s:upsert", ModuleName)
}

func (m MsgUpsertRequest) ValidateBasic() error {
	if _, err := hub.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address should be valid
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Duration shouldn't be negative
	if m.Duration < 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "duration")
	}

	// Bandwidth shouldn't be negative
	if m.Bandwidth.IsAnyNegative() {
		return errors.Wrapf(ErrorInvalidField, "%s", "bandwidth")
	}

	// Signature can be nil, if not length should be 64
	if m.Signature != nil && len(m.Signature) != 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "signature")
	}

	return nil
}

func (m MsgUpsertRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpsertRequest) GetSigners() []sdk.AccAddress {
	from, err := hub.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
