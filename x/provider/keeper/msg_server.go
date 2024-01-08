package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/v12/types"
	"github.com/sentinel-official/hub/v12/x/provider/types"
)

// The following line asserts that the `msgServer` type implements the `types.MsgServiceServer` interface.
var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

// msgServer is a message server that implements the `types.MsgServiceServer` interface.
type msgServer struct {
	Keeper // Keeper is an instance of the main keeper for the module.
}

// NewMsgServiceServer creates a new instance of `types.MsgServiceServer` using the provided Keeper.
func NewMsgServiceServer(k Keeper) types.MsgServiceServer {
	return &msgServer{k}
}

// MsgRegister registers a new provider with the provided details and stores it in the Store.
// It validates the registration request, checks for provider existence, and assigns a unique address to the provider.
func (k *msgServer) MsgRegister(c context.Context, msg *types.MsgRegisterRequest) (*types.MsgRegisterResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to an `sdk.AccAddress`.
	accAddr, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Convert the `accAddr` to a `hubtypes.ProvAddress` to represent the provider address.
	provAddr := hubtypes.ProvAddress(accAddr.Bytes())

	// Check if the provider with the given address exists in the network. If yes, return an error.
	if k.HasProvider(ctx, provAddr) {
		return nil, types.NewErrorDuplicateProvider(provAddr)
	}

	// Get the deposit value from the Store and fund the community pool with the deposit from the provider.
	deposit := k.Deposit(ctx)
	if err = k.FundCommunityPool(ctx, accAddr, deposit); err != nil {
		return nil, err
	}

	// Create a new provider with the provided details and set its status as `Inactive`.
	provider := types.Provider{
		Address:     provAddr.String(),
		Name:        msg.Name,
		Identity:    msg.Identity,
		Website:     msg.Website,
		Description: msg.Description,
		Status:      hubtypes.StatusInactive,
		StatusAt:    ctx.BlockTime(),
	}

	// Save the new provider in the Store.
	k.SetProvider(ctx, provider)

	// Emit an event to notify that a new provider has been registered.
	ctx.EventManager().EmitTypedEvent(
		&types.EventRegister{
			Address: provider.Address,
		},
	)

	return &types.MsgRegisterResponse{}, nil
}

// MsgUpdate updates the details of a provider.
// It validates the update request, checks for provider existence, and updates the provider's details and status.
func (k *msgServer) MsgUpdate(c context.Context, msg *types.MsgUpdateRequest) (*types.MsgUpdateResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	// Convert the `msg.From` address (in Bech32 format) to a `hubtypes.ProvAddress`.
	provAddr, err := hubtypes.ProvAddressFromBech32(msg.From)
	if err != nil {
		return nil, err
	}

	// Get the provider from the Store based on the provided provider address.
	provider, found := k.GetProvider(ctx, provAddr)
	if !found {
		return nil, types.NewErrorProviderNotFound(provAddr)
	}

	// Update the provider's details (name, identity, website, description) if they are provided in the message.
	if len(msg.Name) > 0 {
		provider.Name = msg.Name
	}
	provider.Identity = msg.Identity
	provider.Website = msg.Website
	provider.Description = msg.Description

	// If the status is provided in the message and it is not `StatusUnspecified`, update the provider's status.
	if !msg.Status.Equal(hubtypes.StatusUnspecified) {
		// If the current status of the provider is `Active`, handle the necessary updates for changing to `Inactive` status.
		if provider.Status.Equal(hubtypes.StatusActive) {
			if msg.Status.Equal(hubtypes.StatusInactive) {
				k.DeleteActiveProvider(ctx, provAddr)
			}
		}
		// If the current status of the provider is `Inactive`, handle the necessary updates for changing to `Active` status.
		if provider.Status.Equal(hubtypes.StatusInactive) {
			if msg.Status.Equal(hubtypes.StatusActive) {
				k.DeleteInactiveProvider(ctx, provAddr)
			}
		}

		// Update the provider's status and status timestamp with the provided new status and current block time.
		provider.Status = msg.Status
		provider.StatusAt = ctx.BlockTime()
	}

	// Save the updated provider in the Store.
	k.SetProvider(ctx, provider)

	// Emit an event to notify that the provider details have been updated.
	ctx.EventManager().EmitTypedEvent(
		&types.EventUpdate{
			Address: provider.Address,
		},
	)

	return &types.MsgUpdateResponse{}, nil
}
