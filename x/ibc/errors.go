package ibc

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codeSpaceIBC = 200
)

func errorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}
