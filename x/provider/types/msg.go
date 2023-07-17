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

func (m *MsgRegisterRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if m.Name == "" {
		return errors.Wrap(ErrorInvalidMessage, "name cannot be empty")
	}
	if len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidMessage, "name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidMessage, "identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return errors.Wrapf(ErrorInvalidMessage, "website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return errors.Wrap(ErrorInvalidMessage, err.Error())
		}
	}
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidMessage, "description length cannot be greater than %d chars", 256)
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

func NewMsgUpdateRequest(from hubtypes.ProvAddress, name, identity, website, description string, status hubtypes.Status) *MsgUpdateRequest {
	return &MsgUpdateRequest{
		From:        from.String(),
		Name:        name,
		Identity:    identity,
		Website:     website,
		Description: description,
		Status:      status,
	}
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidMessage, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrap(ErrorInvalidMessage, err.Error())
	}
	if len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidMessage, "name length cannot be greater than %d chars", 64)
	}
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidMessage, "identity length cannot be greater than %d chars", 64)
	}
	if len(m.Website) > 64 {
		return errors.Wrapf(ErrorInvalidMessage, "website length cannot be greater than %d chars", 64)
	}
	if m.Website != "" {
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return errors.Wrap(ErrorInvalidMessage, err.Error())
		}
	}
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidMessage, "description length cannot be greater than %d chars", 256)
	}
	if !m.Status.IsOneOf(hubtypes.StatusUnspecified, hubtypes.StatusActive, hubtypes.StatusInactive) {
		return errors.Wrap(ErrorInvalidMessage, "status must be one of [unspecified, active, inactive]")
	}

	return nil
}

func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
