package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"

	hubtypes "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/session/types"
)

// EndBlock is a function that gets called at the end of every block.
// It processes the inactive sessions and updates their status accordingly.
// The function returns a slice of ValidatorUpdate, but in this case, it always returns nil.
func (k *Keeper) EndBlock(ctx sdk.Context) []abcitypes.ValidatorUpdate {
	// Get the status change delay from the Store.
	statusChangeDelay := k.StatusChangeDelay(ctx)

	// Iterate over all sessions that have become inactive at the current block time.
	k.IterateSessionsForInactiveAt(ctx, ctx.BlockTime(), func(_ int, item types.Session) bool {
		k.Logger(ctx).Info("Found an inactive session", "id", item.ID, "status", item.Status)

		// Delete the session from the InactiveAt index before updating the InactiveAt value.
		k.DeleteSessionForInactiveAt(ctx, item.InactiveAt, item.ID)

		// If the session's status is active, set it to inactive-pending and schedule
		// its next status update based on the status change delay.
		if item.Status.Equal(hubtypes.StatusActive) {
			item.InactiveAt = ctx.BlockTime().Add(statusChangeDelay)
			item.Status = hubtypes.StatusInactivePending
			item.StatusAt = ctx.BlockTime()

			// Save the updated session to the store.
			k.SetSession(ctx, item)
			k.SetSessionForInactiveAt(ctx, item.InactiveAt, item.ID)

			// Emit an event to notify that the session status has been updated.
			ctx.EventManager().EmitTypedEvent(
				&types.EventUpdateStatus{
					Status:         hubtypes.StatusInactivePending,
					Address:        item.Address,
					NodeAddress:    item.NodeAddress,
					ID:             item.ID,
					PlanID:         0,
					SubscriptionID: item.SubscriptionID,
				},
			)

			// Continue the iteration to handle the next session.
			return false
		}

		// If the session's status is not active, we need to end the session and perform necessary cleanup.

		// Get the account address and node address associated with the session.
		var (
			accAddr  = item.GetAddress()
			nodeAddr = item.GetNodeAddress()
		)

		// Call the SessionInactiveHook method of the subscription handler to notify the subscription
		// module that the session has ended. The method handles the necessary logic for payments
		// or other actions related to the session's termination.
		if err := k.subscription.SessionInactiveHook(ctx, item.ID, accAddr, nodeAddr, item.Bandwidth.Sum()); err != nil {
			// If an error occurs during the hook execution, panic to halt the chain.
			// This is done to prevent any inconsistencies or unexpected behavior.
			panic(err)
		}

		// Perform cleanup by deleting the session and its references from the store.
		k.DeleteSession(ctx, item.ID)
		k.DeleteSessionForAccount(ctx, accAddr, item.ID)
		k.DeleteSessionForNode(ctx, nodeAddr, item.ID)
		k.DeleteSessionForSubscription(ctx, item.SubscriptionID, item.ID)
		k.DeleteSessionForAllocation(ctx, item.SubscriptionID, accAddr, item.ID)

		// Emit an event to notify that the session has been terminated.
		ctx.EventManager().EmitTypedEvent(
			&types.EventUpdateStatus{
				Status:         hubtypes.StatusInactive,
				Address:        item.Address,
				NodeAddress:    item.NodeAddress,
				ID:             item.ID,
				PlanID:         0,
				SubscriptionID: item.SubscriptionID,
			},
		)

		// Continue the iteration to handle the next session.
		return false
	})

	// The function always returns nil for ValidatorUpdate slice.
	return nil
}
