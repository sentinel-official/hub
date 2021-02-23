package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hub "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
	_ sdk.Msg = (*MsgSetStatusRequest)(nil)
)

func NewMsgRegisterRequest(from, provider string, price sdk.Coins, remoteUrl string) MsgRegisterRequest {
	return MsgRegisterRequest{
		From:      from,
		Provider:  provider,
		Price:     price,
		RemoteUrl: remoteUrl,
	}
}

func (m MsgRegisterRequest) Route() string {
	return RouterKey
}

func (m MsgRegisterRequest) Type() string {
	return fmt.Sprintf("%s:register", ModuleName)
}

func (m MsgRegisterRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Either provider or price should be empty
	if (m.Provider != "" && m.Price != nil) ||
		(m.Provider == "" && m.Price == nil) {
		return errors.Wrapf(ErrorInvalidField, "%s", "provider or price")
	}

	// Provider can be empty. If not, it should be valid
	if m.Provider != "" {
		if _, err := hub.ProvAddressFromBech32(m.Provider); err != nil {
			return errors.Wrapf(ErrorInvalidField, "%s", "provider")
		}
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "price")
	}

	// RemoteUrl length should be between 1 and 64
	if len(m.RemoteUrl) == 0 || len(m.RemoteUrl) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "remote_url")
	}

	return nil
}

func (m MsgRegisterRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegisterRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateRequest(from, provider string, price sdk.Coins, remoteUrl string) MsgUpdateRequest {
	return MsgUpdateRequest{
		From:      from,
		Provider:  provider,
		Price:     price,
		RemoteUrl: remoteUrl,
	}
}

func (m MsgUpdateRequest) Route() string {
	return RouterKey
}

func (m MsgUpdateRequest) Type() string {
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m MsgUpdateRequest) ValidateBasic() error {
	if _, err := hub.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	if m.Provider != "" && m.Price != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "provider or price")
	}

	// Provider can be empty. If not, it should be valid
	if m.Provider != "" {
		if _, err := hub.ProvAddressFromBech32(m.Provider); err != nil {
			return errors.Wrapf(ErrorInvalidField, "%s", "provider")
		}
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return errors.Wrapf(ErrorInvalidField, "%s", "price")
	}

	// RemoteUrl length should be between 0 and 64
	if len(m.RemoteUrl) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "remote_url")
	}

	return nil
}

func (m MsgUpdateRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := hub.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgSetStatusRequest(from string, status hub.Status) MsgSetStatusRequest {
	return MsgSetStatusRequest{
		From:   from,
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
	if _, err := hub.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
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
	from, err := hub.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
