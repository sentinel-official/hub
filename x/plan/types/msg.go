package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgAddRequest)(nil)
	_ sdk.Msg = (*MsgSetStatusRequest)(nil)
	_ sdk.Msg = (*MsgAddNodeRequest)(nil)
	_ sdk.Msg = (*MsgRemoveNodeRequest)(nil)
)

func NewMsgAddRequest(from hubtypes.ProvAddress, prices sdk.Coins, validity time.Duration, bytes sdk.Int) *MsgAddRequest {
	return &MsgAddRequest{
		From:     from.String(),
		Prices:   prices,
		Validity: validity,
		Bytes:    bytes,
	}
}

func (m *MsgAddRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.Prices != nil {
		if m.Prices.Len() == 0 {
			return errors.Wrap(ErrorInvalidMessage, "prices cannot be empty")
		}
		if !m.Prices.IsValid() {
			return errors.Wrap(ErrorInvalidMessage, "prices must be valid")
		}
	}
	if m.Validity < 0 {
		return errors.Wrap(ErrorInvalidMessage, "validity cannot be negative")
	}
	if m.Validity == 0 {
		return errors.Wrap(ErrorInvalidMessage, "validity cannot be zero")
	}
	if m.Bytes.IsNegative() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be negative")
	}
	if m.Bytes.IsZero() {
		return errors.Wrap(ErrorInvalidMessage, "bytes cannot be zero")
	}

	return nil
}

func (m *MsgAddRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgSetStatusRequest(from hubtypes.ProvAddress, id uint64, status hubtypes.Status) *MsgSetStatusRequest {
	return &MsgSetStatusRequest{
		From:   from.String(),
		ID:     id,
		Status: status,
	}
}

func (m *MsgSetStatusRequest) ValidateBasic() error {
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

func (m *MsgSetStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgAddNodeRequest(from hubtypes.ProvAddress, id uint64, address hubtypes.NodeAddress) *MsgAddNodeRequest {
	return &MsgAddNodeRequest{
		From:    from.String(),
		ID:      id,
		Address: address.String(),
	}
}

func (m *MsgAddNodeRequest) ValidateBasic() error {
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

func (m *MsgAddNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRemoveNodeRequest(from sdk.AccAddress, id uint64, address hubtypes.NodeAddress) *MsgRemoveNodeRequest {
	return &MsgRemoveNodeRequest{
		From:    from.String(),
		ID:      id,
		Address: address.String(),
	}
}

func (m *MsgRemoveNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
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

func (m *MsgRemoveNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
