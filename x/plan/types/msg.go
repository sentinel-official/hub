package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgCreateRequest)(nil)
	_ sdk.Msg = (*MsgUpdateStatusRequest)(nil)
	_ sdk.Msg = (*MsgLinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgUnlinkNodeRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeRequest)(nil)
)

func NewMsgCreateRequest(from hubtypes.ProvAddress, bytes sdk.Int, duration time.Duration, prices sdk.Coins) *MsgCreateRequest {
	return &MsgCreateRequest{
		From:     from.String(),
		Bytes:    bytes,
		Duration: duration,
		Prices:   prices,
	}
}

func (m *MsgCreateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.Bytes.IsNil() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be nil")
	}
	if m.Bytes.IsNegative() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be negative")
	}
	if m.Bytes.IsZero() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be zero")
	}
	if m.Duration < 0 {
		return errors.Wrap(ErrorInvalidMessage, "duration cannot be negative")
	}
	if m.Duration == 0 {
		return errors.Wrap(ErrorInvalidMessage, "duration cannot be zero")
	}
	if m.Prices != nil {
		if m.Prices.Len() == 0 {
			return errors.Wrap(ErrorInvalidMessage, "prices cannot be empty")
		}
		if m.Prices.IsAnyNil() {
			return errors.Wrap(ErrorInvalidMessage, "prices cannot contain nil")
		}
		if !m.Prices.IsValid() {
			return errors.Wrap(ErrorInvalidMessage, "prices must be valid")
		}
	}

	return nil
}

func (m *MsgCreateRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUpdateStatusRequest(from hubtypes.ProvAddress, id uint64, status hubtypes.Status) *MsgUpdateStatusRequest {
	return &MsgUpdateStatusRequest{
		From:   from.String(),
		ID:     id,
		Status: status,
	}
}

func (m *MsgUpdateStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactive) {
		return errors.Wrap(ErrorInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgUpdateStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgLinkNodeRequest(from hubtypes.ProvAddress, id uint64, addr hubtypes.NodeAddress) *MsgLinkNodeRequest {
	return &MsgLinkNodeRequest{
		From:    from.String(),
		ID:      id,
		Address: addr.String(),
	}
}

func (m *MsgLinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidMessage, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgLinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgUnlinkNodeRequest(from hubtypes.ProvAddress, id uint64, addr hubtypes.NodeAddress) *MsgUnlinkNodeRequest {
	return &MsgUnlinkNodeRequest{
		From:    from.String(),
		ID:      id,
		Address: addr.String(),
	}
}

func (m *MsgUnlinkNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidMessage, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}

	return nil
}

func (m *MsgUnlinkNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgSubscribeRequest(from sdk.AccAddress, id uint64, denom string) *MsgSubscribeRequest {
	return &MsgSubscribeRequest{
		From:  from.String(),
		ID:    id,
		Denom: denom,
	}
}

func (m *MsgSubscribeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return errors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Denom != "" {
		if err := sdk.ValidateDenom(m.Denom); err != nil {
			return errors.Wrap(ErrorInvalidMessage, err.Error())
		}
	}

	return nil
}

func (m *MsgSubscribeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
