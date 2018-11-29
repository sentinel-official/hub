package ibc

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codeSpaceIBC = 200

	errCodeEmptyRelayer     = 201
	errCodeInvalidSequence  = 202
	errCodeEmptySrcChainID  = 203
	errCodeEmptyDestChainID = 204

	errMsgEmptyRelayer     = "Empty relayer"
	errMsgInvalidSequence  = "Invalid sequence"
	errMsgEmptySrcChainID  = "Empty source chain ID"
	errMsgEmptyDestChainID = "Empty destination chain ID"
)

func errorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func errorEmptyRelayer() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, errCodeEmptyRelayer, errMsgEmptyRelayer)
}

func errorInvalidIBCSequence() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, sdkTypes.ErrCodeInvalidIBCSequence, sdkTypes.ErrMsgInvalidIBCSequence)
}

func errorEmptySrcChainID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, errCodeEmptySrcChainID, errMsgEmptySrcChainID)
}

func errorEmptyDestChainID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceIBC, errCodeEmptyDestChainID, errMsgEmptyDestChainID)
}
