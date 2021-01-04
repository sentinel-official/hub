package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/exported"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	hub "github.com/sentinel-official/hub/types"
	node "github.com/sentinel-official/hub/x/node/simulation"
	"github.com/sentinel-official/hub/x/session/expected"
	"github.com/sentinel-official/hub/x/session/types"
	subscription "github.com/sentinel-official/hub/x/subscription/simulation"
)

func SimulateUpsert(ak expected.AccountKeeper, nk expected.NodeKeeper, sk expected.SubscriptionKeeper) simulation.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accounts []simulation.Account, chainID string) (
		simulation.OperationMsg, []simulation.FutureOperation, error) {
		var (
			rNode         = node.RandomNode(r, nk.GetNodes(ctx, 0, 0))
			from, account = func() (simulation.Account, exported.Account) {
				from, _ := simulation.FindAccount(accounts, rNode.Address)
				return from, ak.GetAccount(ctx, from.Address)
			}()

			id         = subscription.RandomSubscription(r, sk.GetSubscriptionsForNode(ctx, rNode.Address, 0, 0)).ID
			address, _ = simulation.RandomAcc(r, accounts)
			duration   = time.Duration(r.Int63n(1e3)+1) * time.Second
			bandwidth  = hub.NewBandwidthFromInt64(r.Int63n(1e6)+1, r.Int63n(1e6)+1)
		)

		msg := types.NewMsgUpsert(rNode.Address, id, address.Address, duration, bandwidth)
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
			from.PrivKey,
		)

		_, _, err := app.Deliver(tx)
		if err != nil {
			return simulation.NoOpMsg(types.ModuleName), nil, err
		}

		return simulation.NewOperationMsg(msg, true, ""), nil, nil
	}
}
