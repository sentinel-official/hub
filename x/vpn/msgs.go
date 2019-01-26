package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type MsgRegisterNode struct {
	From csdkTypes.AccAddress `json:"from"`

	Owner        csdkTypes.AccAddress `json:"owner"`
	AmountToLock csdkTypes.Coins      `json:"amount_to_lock"`
	APIPort      uint16               `json:"api_port"`
	NetSpeed     sdkTypes.Bandwidth   `json:"net_speed"`
	EncMethod    string               `json:"enc_method"`
	PerGBAmount  csdkTypes.Coins      `json:"per_gb_amount"`
	Version      string               `json:"version"`
}

func (msg MsgRegisterNode) Type() string {
	return "msg_register_node"
}

func (msg MsgRegisterNode) ValidateBasic() csdkTypes.Error {
	if msg.From == nil || msg.From.Empty() {
		return errorInvalidField("from")
	}
	if msg.Owner == nil || msg.Owner.Empty() {
		return errorInvalidField("owner")
	}
	if msg.AmountToLock == nil || msg.AmountToLock.Len() == 0 ||
		msg.AmountToLock.IsValid() == false || msg.AmountToLock.IsPositive() == false {
		return errorInvalidField("amount_to_lock")
	}
	if msg.APIPort <= 0 || msg.APIPort > 65535 {
		return errorInvalidField("api_port")
	}
	if msg.NetSpeed.Download <= 0 || msg.NetSpeed.Upload <= 0 {
		return errorInvalidField("net_speed")
	}
	if len(msg.EncMethod) == 0 {
		return errorInvalidField("enc_method")
	}
	if msg.PerGBAmount == nil || msg.PerGBAmount.Len() == 0 ||
		msg.PerGBAmount.IsValid() == false || msg.PerGBAmount.IsPositive() == false {
		return errorInvalidField("per_gb_amount")
	}
	if len(msg.Version) == 0 {
		return errorInvalidField("version")
	}

	return nil
}

func (msg MsgRegisterNode) GetSignBytes() []byte {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	return msgBytes
}

func (msg MsgRegisterNode) GetSigners() []csdkTypes.AccAddress {
	return []csdkTypes.AccAddress{msg.From}
}

func (msg MsgRegisterNode) Route() string {
	return sdkTypes.StoreKeyVPNNode
}

func NewMsgRegisterNode(from csdkTypes.AccAddress,
	owner csdkTypes.AccAddress, apiPort uint16,
	upload uint64, download uint64,
	encMethod string, perGBAmount csdkTypes.Coins, version string,
	amountToLock csdkTypes.Coins) *MsgRegisterNode {

	return &MsgRegisterNode{
		From:         from,
		Owner:        owner,
		AmountToLock: amountToLock,
		APIPort:      apiPort,
		NetSpeed: sdkTypes.Bandwidth{
			Upload:   upload,
			Download: download,
		},
		EncMethod:   encMethod,
		PerGBAmount: perGBAmount,
		Version:     version,
	}
}
