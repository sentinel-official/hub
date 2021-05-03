package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
	return fmt.Sprintf("%s:register", ModuleName)
}

func (m *MsgRegisterRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return ErrorInvalidFieldFrom
	}

	// Either provider or price should be empty
	if (m.Provider != "" && m.Price != nil) ||
		(m.Provider == "" && m.Price == nil) {
		return ErrorInvalidFieldProviderOrPrice
	}

	// Provider can be empty. If not, it should be valid
	if m.Provider != "" {
		if _, err := hubtypes.ProvAddressFromBech32(m.Provider); err != nil {
			return ErrorInvalidFieldProvider
		}
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return ErrorInvalidFieldPrice
	}

	// RemoteURL length should be between 1 and 64
	if len(m.RemoteURL) == 0 || len(m.RemoteURL) > 64 {
		return ErrorInvalidFieldRemoteURL
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
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return ErrorInvalidFieldFrom
	}

	if m.Provider != "" && m.Price != nil {
		return ErrorInvalidFieldProviderOrPrice
	}

	// Provider can be empty. If not, it should be valid
	if m.Provider != "" {
		if _, err := hubtypes.ProvAddressFromBech32(m.Provider); err != nil {
			return ErrorInvalidFieldProvider
		}
	}

	// Price can be nil. If not, it should be valid
	if m.Price != nil && !m.Price.IsValid() {
		return ErrorInvalidFieldPrice
	}

	// RemoteURL length should be between 0 and 64
	if len(m.RemoteURL) > 64 {
		return ErrorInvalidFieldRemoteURL
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
	return fmt.Sprintf("%s:set_status", ModuleName)
}

func (m *MsgSetStatusRequest) ValidateBasic() error {
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return ErrorInvalidFieldFrom
	}

	// Status should be either Active or Inactive
	if !m.Status.Equal(hubtypes.StatusActive) && !m.Status.Equal(hubtypes.StatusInactive) {
		return ErrorInvalidFieldStatus
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
