package types

import (
	"net/url"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
)

// The `types` package contains custom message types for the Cosmos SDK.

// The following variables implement the sdk.Msg interface for the provided message types.
// These variables ensure that the corresponding types can be used as messages in the Cosmos SDK.
var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
	_ sdk.Msg = (*MsgUpdateStatusRequest)(nil)
	_ sdk.Msg = (*MsgSubscribeRequest)(nil)
)

// NewMsgRegisterRequest creates a new MsgRegisterRequest instance with the given parameters.
func NewMsgRegisterRequest(from sdk.AccAddress, gigabytePrices, hourlyPrices sdk.Coins, remoteURL string) *MsgRegisterRequest {
	return &MsgRegisterRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

// ValidateBasic performs basic validation checks on the MsgRegisterRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'GigabytePrices' and 'HourlyPrices' fields are valid coins (not empty, not containing nil coins, and having valid coin denominations),
// and if the 'RemoteURL' field is valid (not empty, not longer than 64 characters, and has a valid "https" scheme and non-empty port).
func (m *MsgRegisterRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.GigabytePrices == nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices cannot be nil")
	}
	if m.GigabytePrices.Len() == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices length cannot be zero")
	}
	if m.GigabytePrices.IsAnyNil() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices cannot contain nil")
	}
	if !m.GigabytePrices.IsValid() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices must be valid")
	}
	if m.HourlyPrices == nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices cannot be nil")
	}
	if m.HourlyPrices.Len() == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices length cannot be zero")
	}
	if m.HourlyPrices.IsAnyNil() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices cannot contain nil")
	}
	if !m.HourlyPrices.IsValid() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices must be valid")
	}
	if m.RemoteURL == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "remote_url cannot be empty")
	}
	if len(m.RemoteURL) > 64 {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "remote_url length cannot be greater than %d chars", 64)
	}

	remoteURL, err := url.ParseRequestURI(m.RemoteURL)
	if err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if remoteURL.Scheme != "https" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "remote_url scheme must be https")
	}
	if remoteURL.Port() == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "remote_url port cannot be empty")
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgRegisterRequest.
func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgUpdateDetailsRequest creates a new MsgUpdateDetailsRequest instance with the given parameters.
func NewMsgUpdateDetailsRequest(from hubtypes.NodeAddress, gigabytePrices, hourlyPrices sdk.Coins, remoteURL string) *MsgUpdateDetailsRequest {
	return &MsgUpdateDetailsRequest{
		From:           from.String(),
		GigabytePrices: gigabytePrices,
		HourlyPrices:   hourlyPrices,
		RemoteURL:      remoteURL,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateDetailsRequest fields.
// It checks if the 'From' field is not empty and represents a valid node address,
// if the 'GigabytePrices' and 'HourlyPrices' fields are valid coins (not empty, not containing nil coins, and having valid coin denominations),
// and if the 'RemoteURL' field is valid (not empty, not longer than 64 characters, and has a valid "https" scheme and non-empty port).
func (m *MsgUpdateDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.GigabytePrices != nil {
		if m.GigabytePrices.Len() == 0 {
			return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices length cannot be zero")
		}
		if m.GigabytePrices.IsAnyNil() {
			return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices cannot contain nil")
		}
		if !m.GigabytePrices.IsValid() {
			return sdkerrors.Wrap(ErrorInvalidMessage, "gigabyte_prices must be valid")
		}
	}
	if m.HourlyPrices != nil {
		if m.HourlyPrices.Len() == 0 {
			return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices length cannot be zero")
		}
		if m.HourlyPrices.IsAnyNil() {
			return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices cannot contain nil")
		}
		if !m.HourlyPrices.IsValid() {
			return sdkerrors.Wrap(ErrorInvalidMessage, "hourly_prices must be valid")
		}
	}
	if m.RemoteURL != "" {
		if len(m.RemoteURL) > 64 {
			return sdkerrors.Wrapf(ErrorInvalidMessage, "remote_url length cannot be greater than %d chars", 64)
		}

		remoteURL, err := url.ParseRequestURI(m.RemoteURL)
		if err != nil {
			return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
		}
		if remoteURL.Scheme != "https" {
			return sdkerrors.Wrap(ErrorInvalidMessage, "remote_url scheme must be https")
		}
		if remoteURL.Port() == "" {
			return sdkerrors.Wrap(ErrorInvalidMessage, "remote_url port cannot be empty")
		}
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgUpdateDetailsRequest.
func (m *MsgUpdateDetailsRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgUpdateStatusRequest creates a new MsgUpdateStatusRequest instance with the given parameters.
func NewMsgUpdateStatusRequest(from hubtypes.NodeAddress, status hubtypes.Status) *MsgUpdateStatusRequest {
	return &MsgUpdateStatusRequest{
		From:   from.String(),
		Status: status,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateStatusRequest fields.
// It checks if the 'From' field is not empty and represents a valid node address,
// and if the 'Status' field is one of the allowed values [active, inactive].
func (m *MsgUpdateStatusRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if !m.Status.IsOneOf(hubtypes.StatusActive, hubtypes.StatusInactive) {
		return sdkerrors.Wrap(ErrorInvalidMessage, "status must be one of [active, inactive]")
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgUpdateStatusRequest.
func (m *MsgUpdateStatusRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.NodeAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}

// NewMsgSubscribeRequest creates a new MsgSubscribeRequest instance with the given parameters.
func NewMsgSubscribeRequest(from sdk.AccAddress, addr hubtypes.NodeAddress, gigabytes, hours int64, denom string) *MsgSubscribeRequest {
	return &MsgSubscribeRequest{
		From:        from.String(),
		NodeAddress: addr.String(),
		Gigabytes:   gigabytes,
		Hours:       hours,
		Denom:       denom,
	}
}

// ValidateBasic performs basic validation checks on the MsgSubscribeRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'NodeAddress' field is not empty and represents a valid node address,
// if either 'Gigabytes' or 'Hours' field (but not both) are non-zero and non-negative,
// and if the 'Denom' field is valid according to the Cosmos SDK's ValidateDenom function.
func (m *MsgSubscribeRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.NodeAddress == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "node_address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.NodeAddress); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.Gigabytes == 0 && m.Hours == 0 {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "[gigabytes, hours] cannot be empty")
	}
	if m.Gigabytes != 0 && m.Hours != 0 {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "[gigabytes, hours] cannot be non-empty")
	}
	if m.Gigabytes != 0 {
		if m.Gigabytes < 0 {
			return sdkerrors.Wrap(ErrorInvalidMessage, "gigabytes cannot be negative")
		}
	}
	if m.Hours != 0 {
		if m.Hours < 0 {
			return sdkerrors.Wrap(ErrorInvalidMessage, "hours cannot be negative")
		}
	}
	if m.Denom == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "denom cannot be empty")
	}
	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgSubscribeRequest.
func (m *MsgSubscribeRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
