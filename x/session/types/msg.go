package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgUpsertRequest)(nil)
)

func NewMsgUpsertRequest(proof Proof, address sdk.AccAddress, signature []byte) *MsgUpsertRequest {
	return &MsgUpsertRequest{
		Proof:     proof,
		Address:   address.String(),
		Signature: signature,
	}
}

func (m *MsgUpsertRequest) Route() string {
	return RouterKey
}

func (m *MsgUpsertRequest) Type() string {
	return fmt.Sprintf("%s:upsert", ModuleName)
}

func (m *MsgUpsertRequest) ValidateBasic() error {
	// Subscription shouldn't be zero
	if m.Proof.Subscription == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->subscription")
	}

	// Node shouldn't be nil or empty
	if _, err := hubtypes.NodeAddressFromBech32(m.Proof.Node); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->node")
	}

	// Duration shouldn't be negative
	if m.Proof.Duration < 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->duration")
	}

	// Bandwidth shouldn't be negative
	if m.Proof.Bandwidth.IsAnyNegative() {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->bandwidth")
	}

	// Address shouldn't be nil or empty
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Signature can be nil, if not length should be 64
	if m.Signature != nil && len(m.Signature) != 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "signature")
	}

	return nil
}

func (m *MsgUpsertRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgUpsertRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.NodeAddressFromBech32(m.Proof.Node)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
