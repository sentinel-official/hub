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
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "price")
	}

	// Validity should be positive
	if m.Validity <= 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "validity")
	}

	// Bytes should be positive
	if !m.Bytes.IsPositive() {
		return errors.Wrapf(ErrorInvalidField, "%s", "bytes")
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
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Status should be either Active or Inactive
	if !m.Status.Equal(hubtypes.StatusActive) && !m.Status.Equal(hubtypes.StatusInactive) {
		return errors.Wrapf(ErrorInvalidField, "%s", "status")
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
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address shouldn't be nil or empty
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
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
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address should be valid
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
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
