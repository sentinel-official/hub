package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/node"
	nodesimulation "github.com/sentinel-official/hub/x/node/simulation"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/session/keeper"
	"github.com/sentinel-official/hub/x/session/types"
	"github.com/sentinel-official/hub/x/subscription"
	subscriptionsimulation "github.com/sentinel-official/hub/x/subscription/simulation"
)

func SimulateUpsert(nk node.Keeper, sk subscription.Keeper, k keeper.Keeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			from      = nodesimulation.RandomNode(r, nk.GetNodes(ctx)).Address
			id        = subscriptionsimulation.RandomSubscription(r, sk.GetSubscriptionsForNode(ctx, from)).ID
			address   = simulation.RandomAcc(r, accounts).Address
			duration  = time.Duration(r.Int63n(1e3)) * time.Second
			bandwidth = hub.NewBandwidthFromInt64(r.Int63n(1e6)+1, r.Int63n(1e6)+1)
		)

		msg := types.NewMsgUpsert(from, id, address, duration, bandwidth)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		ok := session.HandleUpsert(ctx, k, msg).IsOK()
		if ok {
			write()
		}

		return simulation.NewOperationMsg(msg, ok, ""), nil, nil
	}
}
