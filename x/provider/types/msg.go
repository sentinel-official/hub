package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	hubtypes "github.com/sentinel-official/hub/types"
)

var (
	_ sdk.Msg = (*MsgRegisterRequest)(nil)
	_ sdk.Msg = (*MsgUpdateRequest)(nil)
)

func NewMsgRegisterRequest(from sdk.AccAddress, name, identity, website, description string) *MsgRegisterRequest {
	return &MsgRegisterRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
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
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Name length should be between 1 and 64
	if len(m.Name) == 0 || len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "name")
	}

	// Identity length should be between 0 and 64
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "identity")
	}

	// Website length should be between 0 and 64
	if len(m.Website) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "website")
	}

	// Description length should be between 0 and 256
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidField, "%s", "description")
	}

	return nil
}

func (m *MsgRegisterRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m *MsgRegisterRequest) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from}
}

func NewMsgUpdateRequest(from hubtypes.ProvAddress, name, identity, website, description string) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
	}
}

func (m *MsgUpdateRequest) Route() string {
	return RouterKey
}

func (m *MsgUpdateRequest) Type() string {
	return fmt.Sprintf("%s:update", ModuleName)
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidField, "%s", "from")
	}

	// Name length should be between 0 and 64
	if len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "name")
	}

	// Identity length should be between 0 and 64
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "identity")
	}

	// Website length should be between 0 and 64
	if len(m.Website) > 64 {
		return errors.Wrapf(ErrorInvalidField, "%s", "website")
	}

	// Description length should be between 0 and 256
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidField, "%s", "description")
	}

	return nil
}

func (m *MsgUpdateRequest) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
