package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/simulation"
	"github.com/sentinel-official/hub/x/session/expected"
	"github.com/sentinel-official/hub/x/session/types"
	subscription "github.com/sentinel-official/hub/x/subscription/simulation"
)

func SimulateUpsert(ak expected.AccountKeeper, pk expected.PlanKeeper, sk expected.SubscriptionKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			toAccount, _  = simulation.RandomAcc(r, accounts)
			subscriptions = sk.GetActiveSubscriptionsForAddress(ctx, toAccount.Address, 0, 0)
		)

		rSubscription := subscription.RandomSubscription(r, subscriptions)
		if rSubscription.Plan == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		nodes := pk.GetNodesForPlan(ctx, rSubscription.Plan, 0, 0)
		if len(nodes) == 0 {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		rNode := node.RandomNode(r, nodes)

		rAccount, found := simulation.FindAccount(accounts, rNode.Address)
		if !found {
			return simulation.NoOpMsg(types.ModuleName), nil, nil
		}

		var (
			account   = ak.GetAccount(ctx, rAccount.Address)
			duration  = time.Duration(r.Int63n(1e3)+1) * time.Second
			bandwidth = hub.NewBandwidthFromInt64(r.Int63n(1e6)+1, r.Int63n(1e6)+1)
			proof     = types.Proof{
				Identity:  rSubscription.ID,
				Channel:   0,
				Address:   rNode.Address,
				Duration:  duration,
				Bandwidth: bandwidth,
			}
		)

		msg := types.NewMsgUpsert(proof, toAccount.Address, nil)
		if msg.ValidateBasic() != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		tx := helpers.GenTx(
			[]sdk.Msg{msg},
			nil,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			rAccount.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}
