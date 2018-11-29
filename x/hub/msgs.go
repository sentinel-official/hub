package hub

import (
	"encoding/json"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

type MsgLockerStatus struct {
	LockerID string `json:"locker_id"`
	Status   string `json:"status"`
}

type MsgLockCoins struct {
	LockerID  string          `json:"locker_id"`
	Coins     csdkTypes.Coins `json:"coins"`
	PubKey    crypto.PubKey   `json:"pub_key"`
	Signature []byte          `json:"signature"`
}

func (msg MsgLockCoins) GetUnSignBytes() []byte {
	bytes, err := json.Marshal(MsgLockCoins{
		LockerID: msg.LockerID,
		Coins:    msg.Coins,
		PubKey:   msg.PubKey,
	})

	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg MsgLockCoins) ValidateBasic() csdkTypes.Error {
	if len(msg.LockerID) == 0 {
		return errorEmptyLockerID()
	}

	if msg.Coins == nil || msg.Coins.Len() == 0 || msg.Coins.IsValid() == false || msg.Coins.IsPositive() == false {
		return errorInvalidCoins()
	}

	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return errorEmptyPubKey()
	}

	if len(msg.Signature) == 0 {
		return errorEmptySignature()
	}

	if msg.PubKey.VerifyBytes(msg.GetUnSignBytes(), msg.Signature) == false {
		return errorSignatureVerificationFailed()
	}

	return nil
}

type MsgReleaseCoins struct {
	LockerID  string        `json:"locker_id"`
	PubKey    crypto.PubKey `json:"pub_key"`
	Signature []byte        `json:"signature"`
}

func (msg MsgReleaseCoins) GetUnSignBytes() []byte {
	bytes, err := json.Marshal(MsgReleaseCoins{
		LockerID: msg.LockerID,
		PubKey:   msg.PubKey,
	})

	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg MsgReleaseCoins) ValidateBasic() csdkTypes.Error {
	if len(msg.LockerID) == 0 {
		return errorEmptyLockerID()
	}

	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return errorEmptyPubKey()
	}

	if len(msg.Signature) == 0 {
		return errorEmptySignature()
	}

	if msg.PubKey.VerifyBytes(msg.GetUnSignBytes(), msg.Signature) == false {
		return errorSignatureVerificationFailed()
	}

	return nil
}

type MsgReleaseCoinsToMany struct {
	LockerID  string                 `json:"locker_id"`
	Addresses []csdkTypes.AccAddress `json:"addresses"`
	Shares    []csdkTypes.Coins      `json:"shares"`
	PubKey    crypto.PubKey          `json:"pub_key"`
	Signature []byte                 `json:"signature"`
}

func (msg MsgReleaseCoinsToMany) GetUnSignBytes() []byte {
	bytes, err := json.Marshal(MsgReleaseCoinsToMany{
		LockerID:  msg.LockerID,
		Addresses: msg.Addresses,
		Shares:    msg.Shares,
		PubKey:    msg.PubKey,
	})

	if err != nil {
		panic(err)
	}

	return bytes
}

func (msg MsgReleaseCoinsToMany) ValidateBasic() csdkTypes.Error {
	if len(msg.LockerID) == 0 {
		return errorEmptyLockerID()
	}

	if len(msg.Addresses) == 0 {
		return errorEmptyAddresses()
	}

	for _, address := range msg.Addresses {
		if address == nil || address.Empty() {
			return errorEmptyAddress()
		}
	}

	if len(msg.Shares) == 0 {
		return errorEmptyShares()
	}

	for _, share := range msg.Shares {
		if share == nil || share.Len() == 0 || share.IsValid() == false || share.IsPositive() == false {
			return errorInvalidCoins()
		}
	}

	if len(msg.Addresses) != len(msg.Shares) {
		return errorAddressesSharesLengthMismatch()
	}

	if msg.PubKey == nil || len(msg.PubKey.Bytes()) == 0 {
		return errorEmptyPubKey()
	}

	if len(msg.Signature) == 0 {
		return errorEmptySignature()
	}

	if msg.PubKey.VerifyBytes(msg.GetUnSignBytes(), msg.Signature) == false {
		return errorSignatureVerificationFailed()
	}

	return nil
}
