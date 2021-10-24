package types

import (
	"fmt"
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

func NewMsgAddRequest(from hubtypes.ProvAddress, price sdk.Coins, validity time.Duration, bytes sdk.Int) *MsgAddRequest {
	return &MsgAddRequest{
		From:     from.String(),
		Price:    price,
		Validity: validity,
		Bytes:    bytes,
	}
}

func (m *MsgAddRequest) Route() string {
	return RouterKey
}

func (m *MsgAddRequest) Type() string {
	return fmt.Sprintf("%s:add", ModuleName)
}

func (m *MsgAddRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Price != nil {
		if m.Price.Len() == 0 {
			return errors.Wrap(ErrorInvalidPrice, "price cannot be empty")
		}
		if !m.Price.IsValid() {
			return errors.Wrap(ErrorInvalidPrice, "price must be valid")
		}
	}
	if m.Validity < 0 {
		return errors.Wrap(ErrorInvalidValidity, "validity cannot be negative")
	}
	if m.Validity == 0 {
		return errors.Wrap(ErrorInvalidValidity, "validity cannot be zero")
	}
	if m.Bytes.IsNegative() {
		return errors.Wrap(ErrorInvalidBytes, "bytes cannot be negative")
	}
	if m.Bytes.IsZero() {
		return errors.Wrap(ErrorInvalidBytes, "bytes cannot be zero")
	}

	return nil
}

func (m *MsgAddRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
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
		Id:     id,
		Status: status,
	}
}

func (m *MsgSetStatusRequest) Route() string {
	return RouterKey
}

func (m *MsgSetStatusRequest) Type() string {
	return fmt.Sprintf("%s:set_status", ModuleName)
}

func (m *MsgSetStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if !m.Status.Equal(hubtypes.StatusActive) && !m.Status.Equal(hubtypes.StatusInactive) {
		return errors.Wrap(ErrorInvalidStatus, "status must be either active or inactive")
	}

	return nil
}

func (m *MsgSetStatusRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
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
		Id:      id,
		Address: address.String(),
	}
}

func (m *MsgAddNodeRequest) Route() string {
	return RouterKey
}

func (m *MsgAddNodeRequest) Type() string {
	return fmt.Sprintf("%s:add_node", ModuleName)
}

func (m *MsgAddNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidAddress, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidAddress, "%s", err)
	}

	return nil
}

func (m *MsgAddNodeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
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
		Id:      id,
		Address: address.String(),
	}
}

func (m *MsgRemoveNodeRequest) Route() string {
	return RouterKey
}

func (m *MsgRemoveNodeRequest) Type() string {
	return fmt.Sprintf("%s:remove_node", ModuleName)
}

func (m *MsgRemoveNodeRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Address == "" {
		return errors.Wrap(ErrorInvalidAddress, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidAddress, "%s", err)
	}

	return nil
}

func (m *MsgRemoveNodeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgRemoveNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
