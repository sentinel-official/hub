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
	_ sdk.Msg = (*MsgAddRequest)(nil)
	_ sdk.Msg = (*MsgSetStatusRequest)(nil)
	_ sdk.Msg = (*MsgAddNodeRequest)(nil)
	_ sdk.Msg = (*MsgRemoveNodeRequest)(nil)
)

func NewMsgAddRequest(from string, price sdk.Coins, validity time.Duration, bytes sdk.Int) MsgAddRequest {
	return MsgAddRequest{
		From:     from,
		Price:    price,
		Validity: validity,
		Bytes:    bytes,
	}
}

func (m MsgAddRequest) Route() string {
	return RouterKey
}

func (m MsgAddRequest) Type() string {
	return fmt.Sprintf("%s:add", ModuleName)
}

func (m MsgAddRequest) ValidateBasic() error {
	if _, err := hub.ProvAddressFromBech32(m.From); err != nil {
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

func (m MsgAddRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddRequest) GetSigners() []sdk.AccAddress {
	from, err := hub.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgSetStatusRequest(from string, id uint64, status hub.Status) MsgSetStatusRequest {
	return MsgSetStatusRequest{
		From:   from,
		Id:     id,
		Status: status,
	}
}

func (m MsgSetStatusRequest) Route() string {
	return RouterKey
}

func (m MsgSetStatusRequest) Type() string {
	return fmt.Sprintf("%s:set_status", ModuleName)
}

func (m MsgSetStatusRequest) ValidateBasic() error {
	if _, err := hub.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Status should be either Active or Inactive
	if !m.Status.Equal(hub.StatusActive) && !m.Status.Equal(hub.StatusInactive) {
		return errors.Wrapf(ErrorInvalidField, "%s", "status")
	}

	return nil
}

func (m MsgSetStatusRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgSetStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := hub.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgAddNodeRequest(from string, id uint64, address string) MsgAddNodeRequest {
	return MsgAddNodeRequest{
		From:    from,
		Id:      id,
		Address: address,
	}
}

func (m MsgAddNodeRequest) Route() string {
	return RouterKey
}

func (m MsgAddNodeRequest) Type() string {
	return fmt.Sprintf("%s:add_node", ModuleName)
}

func (m MsgAddNodeRequest) ValidateBasic() error {
	if _, err := hub.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address shouldn't be nil or empty
	if _, err := hub.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	return nil
}

func (m MsgAddNodeRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgAddNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := hub.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgRemoveNodeRequest(from string, id uint64, address string) MsgRemoveNodeRequest {
	return MsgRemoveNodeRequest{
		From:    from,
		Id:      id,
		Address: address,
	}
}

func (m MsgRemoveNodeRequest) Route() string {
	return RouterKey
}

func (m MsgRemoveNodeRequest) Type() string {
	return fmt.Sprintf("%s:remove_node", ModuleName)
}

func (m MsgRemoveNodeRequest) ValidateBasic() error {
	if _, err := hub.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id shouldn't be zero
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address should be valid
	if _, err := hub.NodeAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	return nil
}

func (m MsgRemoveNodeRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRemoveNodeRequest) GetSigners() []sdk.AccAddress {
	from, err := hub.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
