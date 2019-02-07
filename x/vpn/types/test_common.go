package types

import (
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

var (
	Coin     = csdkTypes.Coin{"sent", csdkTypes.NewInt(10)}
	CoinNeg  = csdkTypes.Coin{"sent", csdkTypes.NewInt(-10)}
	CoinZero = csdkTypes.Coin{"sent", csdkTypes.NewInt(0)}

	Coins        = csdkTypes.Coins{Coin, csdkTypes.Coin{"sut", csdkTypes.NewInt(100)}}
	CoinsNeg     = csdkTypes.Coins{CoinNeg, csdkTypes.Coin{"sut", csdkTypes.NewInt(-100)}}
	CoinsInvalid = csdkTypes.Coins{CoinZero, csdkTypes.Coin{"sut", csdkTypes.NewInt(0)}}

	ClientAddress1 = csdkTypes.AccAddress([]byte("clientAddress-1"))
	ClientAddress2 = csdkTypes.AccAddress([]byte("clientAddress-1"))

	NodeAddress1 = csdkTypes.AccAddress([]byte("nodeAddress-1"))
	NodeAddress2 = csdkTypes.AccAddress([]byte("nodeAddress-2"))

	UploadNeg  = csdkTypes.NewInt(-10)
	UploadZero = csdkTypes.NewInt(0)
	UploadPos  = csdkTypes.NewInt(100)

	DownloadNeg  = csdkTypes.NewInt(-10)
	DownloadZero = csdkTypes.NewInt(0)
	DownloadPos  = csdkTypes.NewInt(100)
)
