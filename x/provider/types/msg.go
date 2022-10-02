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

func (m *MsgRegisterRequest) Route() string {
	return RouterKey
}

func (m *MsgRegisterRequest) Type() string {
	return TypeMsgRegisterRequest
}

func (m *MsgRegisterRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := sdk.AccAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if m.Name == "" {
		return errors.Wrap(ErrorInvalidName, "name cannot be empty")
	}
	if len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidName, "name length cannot be greater than %d", 64)
	}
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidIdentity, "identity length cannot be greater than %d", 64)
	}
	if m.Website != "" {
		if len(m.Website) > 64 {
			return errors.Wrapf(ErrorInvalidWebsite, "website length cannot be greater than %d", 64)
		}
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return errors.Wrapf(ErrorInvalidWebsite, "%s", err)
		}
	}
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidDescription, "description length cannot be greater than %d", 256)
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
	return TypeMsgUpdateRequest
}

func (m *MsgUpdateRequest) ValidateBasic() error {
	if m.From == "" {
		return errors.Wrap(ErrorInvalidFrom, "from cannot be empty")
	}
	if _, err := hubtypes.ProvAddressFromBech32(m.From); err != nil {
		return errors.Wrapf(ErrorInvalidFrom, "%s", err)
	}
	if len(m.Name) > 64 {
		return errors.Wrapf(ErrorInvalidName, "name length cannot be greater than %d", 64)
	}
	if len(m.Identity) > 64 {
		return errors.Wrapf(ErrorInvalidIdentity, "identity length cannot be greater than %d", 64)
	}
	if m.Website != "" {
		if len(m.Website) > 64 {
			return errors.Wrapf(ErrorInvalidWebsite, "website length cannot be greater than %d", 64)
		}
		if _, err := url.ParseRequestURI(m.Website); err != nil {
			return errors.Wrapf(ErrorInvalidWebsite, "%s", err)
		}
	}
	if len(m.Description) > 256 {
		return errors.Wrapf(ErrorInvalidDescription, "description length cannot be greater than %d", 256)
	}

	return nil
}

func (m *MsgUpdateRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgUpdateRequest) GetSigners() []sdk.AccAddress {
	from, err := hubtypes.ProvAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{from.Bytes()}
}
