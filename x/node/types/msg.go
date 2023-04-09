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

func NewMsgRegisterRequest(from sdk.AccAddress, gigabytePrices, hourlyPrices sdk.Coins, remoteURL string) *MsgRegisterRequest {
	return &MsgRegisterRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

func (m *MsgRegisterRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.GigabytePrices != nil {
		if m.GigabytePrices.Len() == 0 {
			return errors.Wrap(ErrorInvalidMessage, "gigabyte_prices length cannot be zero")
		}
		if !m.GigabytePrices.IsValid() {
			return errors.Wrap(ErrorInvalidMessage, "gigabyte_prices must be valid")
		}
	}
	if m.HourlyPrices != nil {
		if m.HourlyPrices.Len() == 0 {
			return errors.Wrap(ErrorInvalidMessage, "hourly_prices length cannot be zero")
		}
		if !m.HourlyPrices.IsValid() {
			return errors.Wrap(ErrorInvalidMessage, "hourly_prices must be valid")
		}
	}
	if m.RemoteURL == "" {
		return errors.Wrap(ErrorInvalidMessage, "remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return errors.Wrap(ErrorInvalidMessage, "remote_url length cannot be greater than 64 chars")
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if remoteURL.Scheme != "https" {
		return errors.Wrap(ErrorInvalidMessage, "remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return errors.Wrap(ErrorInvalidMessage, "remote_url port cannot be empty")
	}

	return nil
}

func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateRequest(from hubtypes.NodeAddress, gigabytePrices, hourlyPrices sdk.Coins, remoteURL string) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.GigabytePrices != nil {
		if m.GigabytePrices.Len() == 0 {
			return errors.Wrap(ErrorInvalidMessage, "gigabyte_prices length cannot be zero")
		}
		if !m.GigabytePrices.IsValid() {
			return errors.Wrap(ErrorInvalidMessage, "gigabyte_prices must be valid")
		}
	}
	if m.HourlyPrices != nil {
		if m.HourlyPrices.Len() == 0 {
			return errors.Wrap(ErrorInvalidMessage, "hourly_prices length cannot be zero")
		}
		if !m.HourlyPrices.IsValid() {
			return errors.Wrap(ErrorInvalidMessage, "hourly_prices must be valid")
		}
	}
	if m.RemoteURL != "" {
		if len(m.RemoteURL) > 64 {
			return errors.Wrap(ErrorInvalidMessage, "remote_url length cannot be greater than 64 chars")
		}

		remoteURL, err := url.ParseRequestURI(m.RemoteURL)
		if err != nil {
			return errors.Wrap(ErrorInvalidMessage, err.Error())
		}
		if remoteURL.Scheme != "https" {
			return errors.Wrap(ErrorInvalidMessage, "remote_url scheme must be https")
		}
		if remoteURL.Port() == "" {
			return errors.Wrap(ErrorInvalidMessage, "remote_url port cannot be empty")
		}
	}

	return nil
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

func (m *MsgSetStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactive) {
		return errors.Wrap(ErrorInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

func (m *MsgSetStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
