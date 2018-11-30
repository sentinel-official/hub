package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	codeSpaceSDK = csdkTypes.CodespaceType(100)

	ErrCodeMarshal            = 101
	ErrCodeUnmarshal          = 102
	ErrCodeInvalidIBCSequence = 103
	ErrCodeEmptyAddress       = 104
	ErrCodeEmptyLockerID      = 105
	ErrCodeInvalidCoins       = 106
	ErrCodeEmptyPubKey        = 107
	ErrCodeEmptySignature     = 108

	ErrMsgMarshal            = "Error occurred while marshalling data"
	ErrMsgUnmarshal          = "Error occurred while unmarshalling data"
	ErrMsgInvalidIBCSequence = "Invalid IBC sequence"
	ErrMsgEmptyAddress       = "Empty address"
	ErrMsgEmptyLockerID      = "Empty locker ID"
	ErrMsgInvalidCoins       = "Invalid coins"
	ErrMsgEmptyPubKey        = "Empty public key"
	ErrMsgEmptySignature     = "Empty signature"
)

func ErrorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceSDK, ErrCodeMarshal, ErrMsgMarshal)
}

func ErrorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceSDK, ErrCodeUnmarshal, ErrMsgUnmarshal)
}
