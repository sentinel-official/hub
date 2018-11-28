package hub

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codeSpaceHub = 300

	errCodeLockerAlreadyExists       = 301
	errCodeLockerNotExists           = 302
	errCodeInvalidLockerOwnerAddress = 303
	errCodeInvalidLockerStatus       = 304

	errMsgLockerAlreadyExists       = "Locker already exists"
	errMsgLockerNotExists           = "Locker not exists"
	errMsgInvalidLockerOwnerAddress = "Invalid locker owner address"
	errMsgInvalidLockerStatus       = "Invalid locker status"
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

func errorIBCPacketMsgVerificationFailed() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceHub, sdkTypes.ErrCodeIBCPacketMsgVerificationFailed, sdkTypes.ErrMsgIBCPacketMsgVerificationFailed)
}

func errorCodeLockerAlreadyExists() csdkTypes.Error {
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
