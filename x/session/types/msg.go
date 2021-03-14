package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgUpsert)(nil)
)

// MsgUpsert is for updating or inserting a session of a plan.
type MsgUpsert struct {
	Proof     Proof          `json:"proof"`
	Address   sdk.AccAddress `json:"address"`
	Signature []byte         `json:"signature,omitempty"`
}

func NewMsgUpsert(proof Proof, address sdk.AccAddress, signature []byte) MsgUpsert {
	return MsgUpsert{
		Proof:     proof,
		Address:   address,
		Signature: signature,
	}
}

func (m MsgUpsert) Route() string {
	return RouterKey
}

func (m MsgUpsert) Type() string {
	return fmt.Sprintf("%s:upsert", ModuleName)
}

func (m MsgUpsert) ValidateBasic() error {
	// Subscription shouldn't be zero
	if m.Proof.Subscription == 0 {
		return errors.Wrapf(ErrorInvalidField, "%s", "proof->subscription")
	}

	// Node shouldn't be nil or empty
	if m.Proof.Node == nil || m.Proof.Node.Empty() {
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
	if m.Address == nil || m.Address.Empty() {
		return errors.Wrapf(ErrorInvalidField, "%s", "address")
	}

	// Signature can be nil, if not length should be 64
	if m.Signature != nil && len(m.Signature) != 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "signature")
	}

	return nil
}

func (m MsgUpsert) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgUpsert) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Proof.Node.Bytes()}
}
