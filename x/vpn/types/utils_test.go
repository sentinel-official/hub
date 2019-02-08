package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	CoinPos   = csdkTypes.NewInt64Coin("sent", 10)
	CoinNeg   = csdkTypes.Coin{"sent", csdkTypes.NewInt(-10)}
	CoinZero  = csdkTypes.NewInt64Coin("sent", 0)
	CoinEmpty = csdkTypes.NewInt64Coin("", 0)
	CoinNil   = csdkTypes.Coin{"", csdkTypes.NewInt(0)}

	CoinsPos     = csdkTypes.Coins{CoinPos, csdkTypes.NewInt64Coin("sut", 100)}
	CoinsNeg     = csdkTypes.Coins{CoinNeg, csdkTypes.Coin{"sut", csdkTypes.NewInt(-100)}}
	CoinsZero    = csdkTypes.Coins{CoinZero, csdkTypes.NewInt64Coin("sut", 0)}
	CoinsInvalid = csdkTypes.Coins{csdkTypes.NewInt64Coin("sut", 100), CoinZero}
	CoinsEmpty   = csdkTypes.Coins{}

	Address1     = csdkTypes.AccAddress([]byte("address_1"))
	AddressEmpty = csdkTypes.AccAddress([]byte(""))

	UploadNeg  = csdkTypes.NewInt(-10)
	UploadZero = csdkTypes.NewInt(0)
	UploadPos  = csdkTypes.NewInt(10)

	DownloadNeg  = csdkTypes.NewInt(-10)
	DownloadZero = csdkTypes.NewInt(0)
	DownloadPos  = csdkTypes.NewInt(10)

	APIPortValid   = NewAPIPort(8000)
	APIPortInvalid = NewAPIPort(0)

	EncMethod = "enc_method"
	NodeType  = "node_type"
	Version   = "version"

	NodeIDValid   = NewNodeID("address/count")
	NodeIDInvalid = NewNodeID("invalid")
	NodeIDEmpty   = NewNodeID("")

	StatusInvalid = "invalid"

	SessionIDValid   = NewSessionID("address/count")
	SessionIDInvalid = NewSessionID("invalid")
	SessionIDEmpty   = NewSessionID("")

	ClientSign    = []byte("client_sign")
	NodeOwnerSign = []byte("node_owner_sign")
)
