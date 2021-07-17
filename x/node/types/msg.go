package types

import (
	"net/url"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
	_ sdk.Msg = (*MsgSetStatusRequest)(nil)
)

func NewMsgRegisterRequest(from sdk.AccAddress, provider hubtypes.ProvAddress, price sdk.Coins, remoteURL string) *MsgRegisterRequest {
	return &MsgRegisterRequest{
		From:      from.String(),
		Provider:  provider.String(),
		Price:     price,
		RemoteURL: remoteURL,
	}
}

func (m *MsgRegisterRequest) Route() string {
	return RouterKey
}

func (m *MsgRegisterRequest) Type() string {
	return TypeMsgRegisterRequest
}

func (m *MsgRegisterRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Provider == "" && m.Price == nil {
		return errors.Wrap(ErrorInvalidField, "both provider and price cannot be empty")
	}
	if m.Provider != "" && m.Price != nil {
		return errors.Wrap(ErrorInvalidField, "either provider or price must be empty")
	}
	if m.Provider != "" {
		if _, err := hubtypes.ProvAddressFromBech32(m.Provider); err != nil {
			return errors.Wrapf(ErrorInvalidProvider, "%s", err)
		}
	}
	if m.Price != nil {
		if m.Price.Len() == 0 {
			return errors.Wrap(ErrorInvalidPrice, "price cannot be empty")
		}
		if !m.Price.IsValid() {
			return errors.Wrap(ErrorInvalidPrice, "price must be valid")
		}
	}
	if m.RemoteURL == "" {
		return errors.Wrap(ErrorInvalidRemoteURL, "remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return errors.Wrapf(ErrorInvalidRemoteURL, "remote_url length cannot be greater than %d", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return errors.Wrapf(ErrorInvalidRemoteURL, "%s", err)
	}
	if remoteURL.Scheme != "https" {
		return errors.Wrap(ErrorInvalidRemoteURL, "remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return errors.Wrap(ErrorInvalidRemoteURL, "remote_url port cannot be empty")
	}

	return nil
}

func (m *MsgRegisterRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateRequest(from hubtypes.NodeAddress, provider hubtypes.ProvAddress, price sdk.Coins, remoteURL string) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:      from.String(),
		Provider:  provider.String(),
		Price:     price,
		RemoteURL: remoteURL,
	}
}

func (m *MsgUpdateRequest) Route() string {
	return RouterKey
}

func (m *MsgUpdateRequest) Type() string {
	return TypeMsgUpdateRequest
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Provider != "" && m.Price != nil {
		return errors.Wrap(ErrorInvalidField, "either provider or price must be empty")
	}
	if m.Provider != "" {
		if _, err := hubtypes.ProvAddressFromBech32(m.Provider); err != nil {
			return errors.Wrapf(ErrorInvalidProvider, "%s", err)
		}
	}
	if m.Price != nil {
		if m.Price.Len() == 0 {
			return errors.Wrap(ErrorInvalidPrice, "price cannot be empty")
		}
		if !m.Price.IsValid() {
			return errors.Wrap(ErrorInvalidPrice, "price must be valid")
		}
	}
	if m.RemoteURL != "" {
		if len(m.RemoteURL) > 64 {
			return errors.Wrapf(ErrorInvalidRemoteURL, "remote_url length cannot be greater than %d", 64)
		}

		remoteURL, err := url.ParseRequestURI(m.RemoteURL)
		if err != nil {
			return errors.Wrapf(ErrorInvalidRemoteURL, "%s", err)
		}
		if remoteURL.Scheme != "https" {
			return errors.Wrap(ErrorInvalidRemoteURL, "remote_url scheme must be https")
		}
		if remoteURL.Port() == "" {
			return errors.Wrap(ErrorInvalidRemoteURL, "remote_url port cannot be empty")
		}
	}

	return nil
}

func (m *MsgUpdateRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

func NewMsgSetStatusRequest(from hubtypes.NodeAddress, status hubtypes.Status) *MsgSetStatusRequest {
	return &MsgSetStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

func (m *MsgSetStatusRequest) Route() string {
	return RouterKey
}

func (m *MsgSetStatusRequest) Type() string {
	return TypeMsgSetStatusRequest
}

func (m *MsgSetStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
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
	from, err := hubtypes.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
