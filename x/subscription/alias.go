// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/types
package subscription

import (
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/types"
)

const (
	ModuleName              = types.ModuleName
	QuerierRoute            = types.QuerierRoute
	DefaultInactiveDuration = types.DefaultInactiveDuration
)

var (
	// functions aliases
	NewKeeper                                  = keeper.NewKeeper
	NewMsgServiceServer                        = keeper.NewMsgServiceServer
	NewQueryServiceServer                      = keeper.NewQueryServiceServer
	RegisterLegacyAminoCodec                   = types.RegisterLegacyAminoCodec
	RegisterInterfaces                         = types.RegisterInterfaces
	NewGenesisState                            = types.NewGenesisState
	DefaultGenesisState                        = types.DefaultGenesisState
	SubscriptionKey                            = types.SubscriptionKey
	GetSubscriptionForNodeKeyPrefix            = types.GetSubscriptionForNodeKeyPrefix
	SubscriptionForNodeKey                     = types.SubscriptionForNodeKey
	GetSubscriptionForPlanKeyPrefix            = types.GetSubscriptionForPlanKeyPrefix
	SubscriptionForPlanKey                     = types.SubscriptionForPlanKey
	GetActiveSubscriptionForAddressKeyPrefix   = types.GetActiveSubscriptionForAddressKeyPrefix
	ActiveSubscriptionForAddressKey            = types.ActiveSubscriptionForAddressKey
	GetInactiveSubscriptionForAddressKeyPrefix = types.GetInactiveSubscriptionForAddressKeyPrefix
	InactiveSubscriptionForAddressKey          = types.InactiveSubscriptionForAddressKey
	GetInactiveSubscriptionAtKeyPrefix         = types.GetInactiveSubscriptionAtKeyPrefix
	InactiveSubscriptionAtKey                  = types.InactiveSubscriptionAtKey
	GetQuotaKeyPrefix                          = types.GetQuotaKeyPrefix
	QuotaKey                                   = types.QuotaKey
	IDFromSubscriptionForNodeKey               = types.IDFromSubscriptionForNodeKey
	IDFromSubscriptionForPlanKey               = types.IDFromSubscriptionForPlanKey
	IDFromStatusSubscriptionForAddressKey      = types.IDFromStatusSubscriptionForAddressKey
	IDFromInactiveSubscriptionAtKey            = types.IDFromInactiveSubscriptionAtKey
	NewMsgSubscribeToNodeRequest               = types.NewMsgSubscribeToNodeRequest
	NewMsgSubscribeToPlanRequest               = types.NewMsgSubscribeToPlanRequest
	NewMsgCancelRequest                        = types.NewMsgCancelRequest
	NewMsgAddQuotaRequest                      = types.NewMsgAddQuotaRequest
	NewMsgUpdateQuotaRequest                   = types.NewMsgUpdateQuotaRequest
	NewMsgServiceClient                        = types.NewMsgServiceClient
	RegisterMsgServiceServer                   = types.RegisterMsgServiceServer
	NewParams                                  = types.NewParams
	DefaultParams                              = types.DefaultParams
	ParamsKeyTable                             = types.ParamsKeyTable
	NewQuerySubscriptionRequest                = types.NewQuerySubscriptionRequest
	NewQuerySubscriptionsRequest               = types.NewQuerySubscriptionsRequest
	NewQuerySubscriptionsForNodeRequest        = types.NewQuerySubscriptionsForNodeRequest
	NewQuerySubscriptionsForPlanRequest        = types.NewQuerySubscriptionsForPlanRequest
	NewQuerySubscriptionsForAddressRequest     = types.NewQuerySubscriptionsForAddressRequest
	NewQueryQuotaRequest                       = types.NewQueryQuotaRequest
	NewQueryQuotasRequest                      = types.NewQueryQuotasRequest
	NewQueryServiceClient                      = types.NewQueryServiceClient
	RegisterQueryServiceServer                 = types.RegisterQueryServiceServer
	RegisterQueryServiceHandlerServer          = types.RegisterQueryServiceHandlerServer
	RegisterQueryServiceHandlerFromEndpoint    = types.RegisterQueryServiceHandlerFromEndpoint
	RegisterQueryServiceHandler                = types.RegisterQueryServiceHandler
	RegisterQueryServiceHandlerClient          = types.RegisterQueryServiceHandlerClient

	// variable aliases
	ModuleCdc                               = types.ModuleCdc
	ErrorMarshal                            = types.ErrorMarshal
	ErrorUnmarshal                          = types.ErrorUnmarshal
	ErrorUnknownMsgType                     = types.ErrorUnknownMsgType
	ErrorUnknownQueryType                   = types.ErrorUnknownQueryType
	ErrorInvalidFieldFrom                   = types.ErrorInvalidFieldFrom
	ErrorPlanDoesNotExist                   = types.ErrorPlanDoesNotExist
	ErrorNodeDoesNotExist                   = types.ErrorNodeDoesNotExist
	ErrorUnauthorized                       = types.ErrorUnauthorized
	ErrorInvalidPlanStatus                  = types.ErrorInvalidPlanStatus
	ErrorPriceDoesNotExist                  = types.ErrorPriceDoesNotExist
	ErrorInvalidNodeStatus                  = types.ErrorInvalidNodeStatus
	ErrorSubscriptionDoesNotExist           = types.ErrorSubscriptionDoesNotExist
	ErrorInvalidSubscriptionStatus          = types.ErrorInvalidSubscriptionStatus
	ErrorCanNotSubscribe                    = types.ErrorCanNotSubscribe
	ErrorInvalidQuota                       = types.ErrorInvalidQuota
	ErrorDuplicateQuota                     = types.ErrorDuplicateQuota
	ErrorQuotaDoesNotExist                  = types.ErrorQuotaDoesNotExist
	ErrorCanNotAddQuota                     = types.ErrorCanNotAddQuota
	ErrorInvalidFieldId                     = types.ErrorInvalidFieldId
	ErrorInvalidFieldAddress                = types.ErrorInvalidFieldAddress
	ErrorInvalidFieldBytes                  = types.ErrorInvalidFieldBytes
	ErrorInvalidFieldDenom                  = types.ErrorInvalidFieldDenom
	ErrorInvalidFieldDeposit                = types.ErrorInvalidFieldDeposit
	ErrInvalidLengthEvents                  = types.ErrInvalidLengthEvents
	ErrIntOverflowEvents                    = types.ErrIntOverflowEvents
	ErrUnexpectedEndOfGroupEvents           = types.ErrUnexpectedEndOfGroupEvents
	ErrInvalidLengthGenesis                 = types.ErrInvalidLengthGenesis
	ErrIntOverflowGenesis                   = types.ErrIntOverflowGenesis
	ErrUnexpectedEndOfGroupGenesis          = types.ErrUnexpectedEndOfGroupGenesis
	ParamsSubspace                          = types.ParamsSubspace
	RouterKey                               = types.RouterKey
	StoreKey                                = types.StoreKey
	EventModuleName                         = types.EventModuleName
	CountKey                                = types.CountKey
	SubscriptionKeyPrefix                   = types.SubscriptionKeyPrefix
	SubscriptionForNodeKeyPrefix            = types.SubscriptionForNodeKeyPrefix
	SubscriptionForPlanKeyPrefix            = types.SubscriptionForPlanKeyPrefix
	ActiveSubscriptionForAddressKeyPrefix   = types.ActiveSubscriptionForAddressKeyPrefix
	InactiveSubscriptionForAddressKeyPrefix = types.InactiveSubscriptionForAddressKeyPrefix
	InactiveSubscriptionAtKeyPrefix         = types.InactiveSubscriptionAtKeyPrefix
	QuotaKeyPrefix                          = types.QuotaKeyPrefix
	ErrInvalidLengthMsg                     = types.ErrInvalidLengthMsg
	ErrIntOverflowMsg                       = types.ErrIntOverflowMsg
	ErrUnexpectedEndOfGroupMsg              = types.ErrUnexpectedEndOfGroupMsg
	KeyInactiveDuration                     = types.KeyInactiveDuration
	ErrInvalidLengthParams                  = types.ErrInvalidLengthParams
	ErrIntOverflowParams                    = types.ErrIntOverflowParams
	ErrUnexpectedEndOfGroupParams           = types.ErrUnexpectedEndOfGroupParams
	ErrInvalidLengthQuerier                 = types.ErrInvalidLengthQuerier
	ErrIntOverflowQuerier                   = types.ErrIntOverflowQuerier
	ErrUnexpectedEndOfGroupQuerier          = types.ErrUnexpectedEndOfGroupQuerier
	ErrInvalidLengthQuota                   = types.ErrInvalidLengthQuota
	ErrIntOverflowQuota                     = types.ErrIntOverflowQuota
	ErrUnexpectedEndOfGroupQuota            = types.ErrUnexpectedEndOfGroupQuota
	ErrInvalidLengthSubscription            = types.ErrInvalidLengthSubscription
	ErrIntOverflowSubscription              = types.ErrIntOverflowSubscription
	ErrUnexpectedEndOfGroupSubscription     = types.ErrUnexpectedEndOfGroupSubscription
)

type (
	Keeper                               = keeper.Keeper
	EventModule                          = types.EventModule
	EventSetSubscriptionCount            = types.EventSetSubscriptionCount
	EventSubscribeToNode                 = types.EventSubscribeToNode
	EventSubscribeToPlan                 = types.EventSubscribeToPlan
	EventCancelSubscription              = types.EventCancelSubscription
	EventAddQuota                        = types.EventAddQuota
	EventUpdateQuota                     = types.EventUpdateQuota
	GenesisSubscriptions                 = types.GenesisSubscriptions
	GenesisSubscription                  = types.GenesisSubscription
	GenesisState                         = types.GenesisState
	MsgSubscribeToNodeRequest            = types.MsgSubscribeToNodeRequest
	MsgSubscribeToPlanRequest            = types.MsgSubscribeToPlanRequest
	MsgCancelRequest                     = types.MsgCancelRequest
	MsgAddQuotaRequest                   = types.MsgAddQuotaRequest
	MsgUpdateQuotaRequest                = types.MsgUpdateQuotaRequest
	MsgSubscribeToNodeResponse           = types.MsgSubscribeToNodeResponse
	MsgSubscribeToPlanResponse           = types.MsgSubscribeToPlanResponse
	MsgCancelResponse                    = types.MsgCancelResponse
	MsgAddQuotaResponse                  = types.MsgAddQuotaResponse
	MsgUpdateQuotaResponse               = types.MsgUpdateQuotaResponse
	MsgServiceClient                     = types.MsgServiceClient
	MsgServiceServer                     = types.MsgServiceServer
	UnimplementedMsgServiceServer        = types.UnimplementedMsgServiceServer
	Params                               = types.Params
	QuerySubscriptionsRequest            = types.QuerySubscriptionsRequest
	QuerySubscriptionsForNodeRequest     = types.QuerySubscriptionsForNodeRequest
	QuerySubscriptionsForPlanRequest     = types.QuerySubscriptionsForPlanRequest
	QuerySubscriptionsForAddressRequest  = types.QuerySubscriptionsForAddressRequest
	QuerySubscriptionRequest             = types.QuerySubscriptionRequest
	QueryQuotaRequest                    = types.QueryQuotaRequest
	QueryQuotasRequest                   = types.QueryQuotasRequest
	QuerySubscriptionsResponse           = types.QuerySubscriptionsResponse
	QuerySubscriptionsForNodeResponse    = types.QuerySubscriptionsForNodeResponse
	QuerySubscriptionsForPlanResponse    = types.QuerySubscriptionsForPlanResponse
	QuerySubscriptionsForAddressResponse = types.QuerySubscriptionsForAddressResponse
	QuerySubscriptionResponse            = types.QuerySubscriptionResponse
	QueryQuotaResponse                   = types.QueryQuotaResponse
	QueryQuotasResponse                  = types.QueryQuotasResponse
	QueryServiceClient                   = types.QueryServiceClient
	QueryServiceServer                   = types.QueryServiceServer
	UnimplementedQueryServiceServer      = types.UnimplementedQueryServiceServer
	Quotas                               = types.Quotas
	Quota                                = types.Quota
	Subscriptions                        = types.Subscriptions
	Subscription                         = types.Subscription
)
