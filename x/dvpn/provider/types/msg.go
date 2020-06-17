package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = (*MsgRegisterProvider)(nil)

type MsgRegisterProvider struct {
	From        sdk.AccAddress `json:"from"`
	Name        string         `json:"name"`
	Website     string         `json:"website"`
	Description string         `json:"description"`
}

func NewMsgRegisterProvider(from sdk.AccAddress, name, website, description string) MsgRegisterProvider {
	return MsgRegisterProvider{
		From:        from,
		Name:        name,
		Website:     website,
		Description: description,
	}
}

func (m MsgRegisterProvider) Route() string {
	return RouterKey
}

func (m MsgRegisterProvider) Type() string {
	return "register_provider"
}

func (m MsgRegisterProvider) ValidateBasic() sdk.Error {
	return nil
}

func (m MsgRegisterProvider) GetSignBytes() []byte {
	bytes, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (m MsgRegisterProvider) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.From}
}
