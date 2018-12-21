package ibc

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codeSpaceIBC = csdkTypes.CodespaceType(200)

	errCodeInvalidRelayer   = 201
	errCodeEmptySrcChainID  = 202
	errCodeEmptyDestChainID = 203

	errMsgInvalidRelayer   = "Invalid relayer"
	errMsgEmptySrcChainID  = "Empty source chain ID"
	errMsgEmptyDestChainID = "Empty destination chain ID"
)

func errorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func errorInvalidRelayer() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, errCodeInvalidRelayer, errMsgInvalidRelayer)
}

func errorEmptySrcChainID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, errCodeEmptySrcChainID, errMsgEmptySrcChainID)
}

func errorEmptyDestChainID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, errCodeEmptyDestChainID, errMsgEmptyDestChainID)
}
