package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgStartRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
	_ sdk.Msg = (*MsgEndRequest)(nil)
)

func NewMsgStartRequest(from sdk.AccAddress, id uint64, node hubtypes.NodeAddress) *MsgStartRequest {
	return &MsgStartRequest{
		From: from.String(),
		Id:   id,
		Node: node.String(),
	}
}

func (m *MsgStartRequest) Route() string {
	return RouterKey
}

func (m *MsgStartRequest) Type() string {
	return fmt.Sprintf("%s:start", ModuleName)
}

func (m *MsgStartRequest) ValidateBasic() error {
	// From shouldn't be nil or empty
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id should be positive
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Address shouldn't be nil or empty
	if _, err := hubtypes.NodeAddressFromBech32(m.Node); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	return nil
}

func (m *MsgStartRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgStartRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateRequest(from hubtypes.NodeAddress, proof Proof, signature []byte) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:      from.String(),
		Proof:     proof,
		Signature: signature,
	}
}

func (m *MsgUpdateRequest) Route() string {
	return RouterKey
}

func (m *MsgUpdateRequest) Type() string {
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	// From shouldn't be nil or empty
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Duration shouldn't be negative
	if m.Proof.Duration < 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->duration")
	}

	// Bandwidth shouldn't be negative
	if m.Proof.Bandwidth.IsAnyNegative() {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->bandwidth")
	}

	// Signature can be nil, if not length should be 64
	if m.Signature != nil && len(m.Signature) != 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "signature")
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

func NewMsgEndRequest(from sdk.AccAddress, id uint64, rating uint64) *MsgEndRequest {
	return &MsgEndRequest{
		From:   from.String(),
		Id:     id,
		Rating: rating,
	}
}

func (m *MsgEndRequest) Route() string {
	return RouterKey
}

func (m *MsgEndRequest) Type() string {
	return fmt.Sprintf("%s:end", ModuleName)
}

func (m *MsgEndRequest) ValidateBasic() error {
	// From shouldn't be nil or empty
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Id should be positive
	if m.Id == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "id")
	}

	// Rating shouldn't be greater than 10
	if m.Rating > 10 {
		return errors.Wrapf(ErrorInvalidField, "%s", "rating")
	}

	return nil
}

func (m *MsgEndRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgEndRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}
