package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterInterface(
		"sentinel.subscription.v2.Subscription",
		(*Subscription)(nil),
		&NodeSubscription{},
		&PlanSubscription{},
	)

	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgCancelRequest{},
		&MsgShareRequest{},
		&MsgUpdateQuotaRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_MsgService_serviceDesc)
}
