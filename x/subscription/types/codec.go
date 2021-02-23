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
	cdc.RegisterConcrete(MsgSubscribeToNodeRequest{}, fmt.Sprintf("x/%s/MsgSubscribeToNode", ModuleName), nil)
	cdc.RegisterConcrete(MsgSubscribeToPlanRequest{}, fmt.Sprintf("x/%s/MsgSubscribeToPlan", ModuleName), nil)
	cdc.RegisterConcrete(MsgCancelRequest{}, fmt.Sprintf("x/%s/MsgCancel", ModuleName), nil)

	cdc.RegisterConcrete(MsgAddQuotaRequest{}, fmt.Sprintf("x/%s/MsgAddQuota", ModuleName), nil)
	cdc.RegisterConcrete(MsgUpdateQuotaRequest{}, fmt.Sprintf("x/%s/MsgUpdateQuota", ModuleName), nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubscribeToNodeRequest{},
		&MsgSubscribeToPlanRequest{},
		&MsgCancelRequest{},
		&MsgAddQuotaRequest{},
		&MsgUpdateQuotaRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
