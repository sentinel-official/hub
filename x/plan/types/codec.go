package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	crypto "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	crypto.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgAddRequest{}, fmt.Sprintf("x/%s/MsgAdd", ModuleName), nil)
	cdc.RegisterConcrete(MsgSetStatusRequest{}, fmt.Sprintf("x/%s/MsgSetStatus", ModuleName), nil)
	cdc.RegisterConcrete(MsgAddNodeRequest{}, fmt.Sprintf("x/%s/MsgAddNode", ModuleName), nil)
	cdc.RegisterConcrete(MsgRemoveNodeRequest{}, fmt.Sprintf("x/%s/MsgRemoveNode", ModuleName), nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddRequest{},
		&MsgSetStatusRequest{},
		&MsgAddNodeRequest{},
		&MsgRemoveNodeRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
