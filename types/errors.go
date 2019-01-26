package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

const (
	codespaceSDK = csdkTypes.CodespaceType(100)

	ErrCodeMarshal   = 101
	ErrCodeUnmarshal = 102

	ErrMsgMarshal   = "Error occurred while marshalling data"
	ErrMsgUnmarshal = "Error occurred while unmarshalling data"
)

func ErrorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespaceSDK, ErrCodeMarshal, ErrMsgMarshal)
}

func ErrorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codespaceSDK, ErrCodeUnmarshal, ErrMsgUnmarshal)
}
