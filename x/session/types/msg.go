package types

import (
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
	return TypeMsgStartRequest
}

func (m *MsgStartRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Node == "" {
		return errors.Wrap(ErrorInvalidNode, "node cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.Node); err != nil {
		return errors.Wrapf(ErrorInvalidNode, "%s", err)
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
	return TypeMsgUpdateRequest
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.NodeAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Proof.Id == 0 {
		return errors.Wrap(ErrorInvalidProofId, "proof->id cannot be zero")
	}
	if m.Proof.Duration < 0 {
		return errors.Wrap(ErrorInvalidProofDuration, "proof->duration cannot be negative")
	}
	if m.Proof.Bandwidth.IsAnyNegative() {
		return errors.Wrap(ErrorInvalidProofBandwidth, "proof->bandwidth cannot be negative")
	}
	if m.Signature != nil {
		if len(m.Signature) < 64 {
			return errors.Wrapf(ErrorInvalidSignature, "signature length cannot be less than %d", 64)
		}
		if len(m.Signature) > 64 {
			return errors.Wrapf(ErrorInvalidSignature, "signature length cannot be greater than %d", 64)
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
	return TypeMsgEndRequest
}

func (m *MsgEndRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Id == 0 {
		return errors.Wrap(ErrorInvalidId, "id cannot be zero")
	}
	if m.Rating > 10 {
		return errors.Wrapf(ErrorInvalidRating, "rating cannot be greater than %d", 10)
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
