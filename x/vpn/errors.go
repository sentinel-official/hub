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
	errCodeInvalidAPIPort          = 408
	errCodeInvalidLocation         = 409
	errCodeInvalidNetSpeed         = 410
	errCodeEmptyEncMethod          = 411
	errCodeInvalidPricePerGB       = 412
	errCodeEmptyVersion            = 413
	errCodeEmptyVPNID              = 414
	errCodeEmptySessionID          = 415
	errCodeInvalidSessionStatus    = 416

	errMsgLockerIDMismatch        = "Locker ID mismatch"
	errMsgVPNAlreadyExists        = "VPN already exists"
	errMsgSessionAlreadyExists    = "Session already exists"
	errMsgVPNNotExists            = "VPN not exits"
	errMsgInvalidNodeOwnerAddress = "Invalid node owner address"
	errMsgInvalidLockStatus       = "Invalid lock status"
	errMsgSessionNotExists        = "Session not exists"
	errMsgInvalidAPIPort          = "Invalid API port"
	errMsgInvalidLocation         = "Invalid location"
	errMsgInvalidNetSpeed         = "Invalid net speed"
	errMsgEmptyEncMethod          = "Empty encryption method"
	errMsgInvalidPricePerGB       = "Invalid price per GB"
	errMsgEmptyVersion            = "Empty version"
	errMsgEmptyVPNID              = "Empty VPN ID"
	errMsgEmptySessionID          = "Empty session ID"
	errMsgInvalidSessionStatus    = "Invalid session status"
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

func errorEmptyAddress() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeEmptyAddress, sdkTypes.ErrMsgEmptyAddress)
}

func errorInvalidAPIPort() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidAPIPort, errMsgInvalidAPIPort)
}

func errorInvalidLocation() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidLocation, errMsgInvalidLocation)
}

func errorInvalidNetSpeed() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidNetSpeed, errMsgInvalidNetSpeed)
}

func errorEmptyEncMethod() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeEmptyEncMethod, errMsgEmptyEncMethod)
}

func errorInvalidPricePerGB() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidPricePerGB, errMsgInvalidPricePerGB)
}

func errorEmptyVersion() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeEmptyVersion, errMsgEmptyVersion)
}

func errorEmptyLockerID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeEmptyLockerID, sdkTypes.ErrMsgEmptyLockerID)
}

func errorInvalidCoins() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeInvalidCoins, sdkTypes.ErrMsgInvalidCoins)
}

func errorEmptyPubKey() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeEmptyPubKey, sdkTypes.ErrMsgEmptyPubKey)
}

func errorEmptySignature() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeEmptySignature, sdkTypes.ErrMsgEmptySignature)
}

func errorEmptyVPNID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeEmptyVPNID, errMsgEmptyVPNID)
}

func errorEmptySessionID() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeEmptySessionID, errMsgEmptySessionID)
}

func errorInvalidSessionStatus() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, errCodeInvalidSessionStatus, errMsgInvalidSessionStatus)
}

func errorInvalidIBCSequence() csdkTypes.Error {
	return csdkTypes.NewError(codeSpaceVPN, sdkTypes.ErrCodeInvalidIBCSequence, sdkTypes.ErrMsgInvalidIBCSequence)
}
