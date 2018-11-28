package vpn

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

const (
	codeSpaceVPN = 400

	errCodeLockerIDMismatch        = 401
	errCodeVPNAlreadyExists        = 402
	errCodeSessionAlreadyExists    = 403
	errCodeVPNNotExists            = 404
	errCodeInvalidNodeOwnerAddress = 405
	errCodeInvalidLockStatus       = 406
	errCodeSessionNotExists        = 407

	errMsgLockerIDMismatch        = "Locker ID mismatch"
	errMsgVPNAlreadyExists        = "VPN already exists"
	errMsgSessionAlreadyExists    = "Session already exists"
	errMsgVPNNotExists            = "VPN not exits"
	errMsgInvalidNodeOwnerAddress = "Invalid node owner address"
	errMsgInvalidLockStatus       = "Invalid lock status"
	errMsgSessionNotExists        = "Session not exists"
)

func errorMarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeMarshal, sdkTypes.ErrMsgMarshal)
}

func errorUnmarshal() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeUnmarshal, sdkTypes.ErrMsgUnmarshal)
}

func errorLockerIDMismatch() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeLockerIDMismatch, errMsgLockerIDMismatch)
}

func errorVPNAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeVPNAlreadyExists, errMsgVPNAlreadyExists)
}

func errorSessionAlreadyExists() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeSessionAlreadyExists, errMsgSessionAlreadyExists)
}

func errorVPNNotExists() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeVPNNotExists, errMsgVPNNotExists)
}

func errorInvalidNodeOwnerAddress() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidNodeOwnerAddress, errMsgInvalidNodeOwnerAddress)
}

func errorInvalidLockStatus() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidLockStatus, errMsgInvalidLockStatus)
}

func errorSessionNotExists() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeSessionNotExists, errMsgSessionNotExists)
}
