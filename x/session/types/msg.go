package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
)

// The `types` package contains custom message types for the Cosmos SDK.

// The following variables implement the sdk.Msg interface for MsgStartRequest, MsgUpdateDetailsRequest, and MsgEndRequest.
// These variables ensure that the corresponding types can be used as messages in the Cosmos SDK.
var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateDetailsRequest)(nil)
	_ sdk.Msg = (*MsgEndRequest)(nil)
)

// NewMsgStartRequest creates a new MsgStartRequest instance with the given parameters.
func NewMsgStartRequest(from sdk.AccAddress, id uint64, addr hubtypes.NodeAddress) *MsgStartRequest {
	return &MsgStartRequest{
		From:    from.String(),
		ID:      id,
		Address: addr.String(),
	}
}

// ValidateBasic performs basic validation checks on the MsgStartRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'ID' field is not zero,
// if the 'Address' field is not empty and represents a valid node address.
func (m *MsgStartRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Address == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "address cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Address); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgStartRequest.
func (m *MsgStartRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

// NewMsgUpdateDetailsRequest creates a new MsgUpdateDetailsRequest instance with the given parameters.
func NewMsgUpdateDetailsRequest(from hubtypes.NodeAddress, proof Proof, signature []byte) *MsgUpdateDetailsRequest {
	return &MsgUpdateDetailsRequest{
		From:      from.String(),
		Proof:     proof,
		Signature: signature,
	}
}

// ValidateBasic performs basic validation checks on the MsgUpdateDetailsRequest fields.
// It checks if the 'From' field is not empty and represents a valid node address,
// if the 'Proof.ID' field is not zero,
// if the 'Proof.Bandwidth' field does not contain nil or negative values,
// if the 'Proof.Duration' field is not negative,
// and if the 'Signature' field has a length of exactly 64 bytes (if not nil).
func (m *MsgUpdateDetailsRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.Proof.ID == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "proof.id cannot be zero")
	}
	if m.Proof.Bandwidth.IsAnyNil() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "proof.bandwidth cannot contain nil")
	}
	if m.Proof.Bandwidth.IsAnyNegative() {
		return sdkerrors.Wrap(ErrorInvalidMessage, "proof.bandwidth cannot be negative")
	}
	if m.Proof.Duration < 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "proof.duration cannot be negative")
	}
	if m.Signature != nil {
		if len(m.Signature) < 64 {
			return sdkerrors.Wrapf(ErrorInvalidMessage, "signature length cannot be less than %d", 64)
		}
		if len(m.Signature) > 64 {
			return sdkerrors.Wrapf(ErrorInvalidMessage, "signature length cannot be greater than %d", 64)
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

// NewMsgEndRequest creates a new MsgEndRequest instance with the given parameters.
func NewMsgEndRequest(from sdk.AccAddress, id uint64, rating uint64) *MsgEndRequest {
	return &MsgEndRequest{
		From:   from.String(),
		ID:     id,
		Rating: rating,
	}
}

// ValidateBasic performs basic validation checks on the MsgEndRequest fields.
// It checks if the 'From' field is not empty and represents a valid account address,
// if the 'ID' field is not zero,
// and if the 'Rating' field is not greater than 10.
func (m *MsgEndRequest) ValidateBasic() error {
	if m.From == "" {
		return sdkerrors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return sdkerrors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.ID == 0 {
		return sdkerrors.Wrap(ErrorInvalidMessage, "id cannot be zero")
	}
	if m.Rating > 10 {
		return sdkerrors.Wrapf(ErrorInvalidMessage, "rating cannot be greater than %d", 10)
	}

	return nil
}

// GetSigners returns an array containing the signer's account address extracted from the 'From' field of the MsgEndRequest.
func (m *MsgEndRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
