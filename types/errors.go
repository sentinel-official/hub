package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	codeSpaceSDK = 100

	ErrCodeMarshal                        = 101
	ErrCodeUnmarshal                      = 102
	ErrCodeInvalidIBCSequence             = 103
	ErrCodeIBCPacketMsgVerificationFailed = 104

	ErrMsgMarshal                        = "Error occurred while marshalling data"
	ErrMsgUnmarshal                      = "Error occurred while unmarshalling data"
	ErrMsgInvalidIBCSequence             = "Invalid IBC sequence"
	ErrMsgIBCPacketMsgVerificationFailed = "IBC packet message verification failed"
)

func ErrorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceSDK, ErrCodeMarshal, ErrMsgMarshal)
}

func ErrorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceSDK, ErrCodeUnmarshal, ErrMsgUnmarshal)
}
