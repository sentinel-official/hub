package hub

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codeSpaceHub = csdkTypes.CodespaceType(300)

	errCodeLockerAlreadyExists           = 301
	errCodeLockerNotExists               = 302
	errCodeInvalidLockerOwnerAddress     = 303
	errCodeInvalidLockerStatus           = 304
	errCodeSignatureVerificationFailed   = 305
	errCodeEmptyAddresses                = 306
	errCodeEmptyShares                   = 307
	errCodeAddressesSharesLengthMismatch = 308

	errMsgLockerAlreadyExists           = "Locker already exists"
	errMsgLockerNotExists               = "Locker not exists"
	errMsgInvalidLockerOwnerAddress     = "Invalid locker owner address"
	errMsgInvalidLockerStatus           = "Invalid locker status"
	errMsgSignatureVerificationFailed   = "Signature verification failed"
	errMsgEmptyAddresses                = "Empty addresses"
	errMsgEmptyShares                   = "Empty shares"
	errMsgAddressesSharesLengthMismatch = "Addresses and shares length mismatch"
)

func errorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func errorInvalidIBCSequence() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeInvalidIBCSequence, sdkTypes.ErrMsgInvalidIBCSequence)
}

func errorLockerAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeLockerAlreadyExists, errMsgLockerAlreadyExists)
}

func errorLockerNotExists() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeLockerNotExists, errMsgLockerNotExists)
}

func errorInvalidLockerOwnerAddress() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeInvalidLockerOwnerAddress, errMsgInvalidLockerOwnerAddress)
}

func errorInvalidLockerStatus() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeInvalidLockerStatus, errMsgInvalidLockerStatus)
}

func errorEmptyPubKey() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeEmptyPubKey, sdkTypes.ErrMsgEmptyPubKey)
}

func errorInvalidCoins() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeInvalidCoins, sdkTypes.ErrMsgInvalidCoins)
}

func errorEmptySignature() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeEmptySignature, sdkTypes.ErrMsgEmptySignature)
}

func errorSignatureVerificationFailed() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeSignatureVerificationFailed, errMsgSignatureVerificationFailed)
}

func errorEmptyAddresses() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeEmptyAddresses, errMsgEmptyAddresses)
}

func errorEmptyLockerID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeEmptyLockerID, sdkTypes.ErrMsgEmptyLockerID)
}

func errorEmptyAddress() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeEmptyAddress, sdkTypes.ErrMsgEmptyAddress)
}

func errorEmptyShares() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeEmptyShares, errMsgEmptyShares)
}

func errorAddressesSharesLengthMismatch() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, errCodeAddressesSharesLengthMismatch, errMsgAddressesSharesLengthMismatch)
}
