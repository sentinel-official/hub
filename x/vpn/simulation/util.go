package simulation

import (
	"math/rand"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	
	hub "github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/vpn/types"
)

var (
	denoms   = []string{"stake", "xxx", "yyy", "zzz"}
	statuses = []string{types.StatusRegistered, types.StatusDeRegistered}
)

func getRandomDenom(r *rand.Rand) string {
	index := r.Intn(len(denoms))
	return denoms[index]
}

func getRandomStatus(r *rand.Rand) string {
	index := r.Intn(len(statuses))
	return statuses[index]
}

func getRandomNodeID(r *rand.Rand) hub.NodeID {
	i := uint64(r.Int63n(10))
	
	return hub.NewNodeID(i)
}
func getRandomSessionID(r *rand.Rand) hub.SessionID {
	i := uint64(r.Int63n(10))
	
	return hub.NewSessionID(i)
}
func getRandomSubscriptionID(r *rand.Rand) hub.SubscriptionID {
	i := uint64(r.Int63n(10))
	
	return hub.NewSubscriptionID(i)
}

func getRandomEncryption(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomType(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomVersion(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomMoniker(r *rand.Rand) string {
	return simulation.RandStringOfLength(r, 10)
}

func getRandomCoin(r *rand.Rand) sdk.Coin {
	denom := getRandomDenom(r)
	amount := simulation.RandIntBetween(r, 1, 1000)
	
	return sdk.NewCoin(denom, sdk.NewInt(int64(amount)))
}

func getRandomCoins(r *rand.Rand) (coins sdk.Coins) {
	coins = append(coins, getRandomCoin(r))
	
	size := r.Intn(2)
	for i := 0; i < size; i++ {
		coin := getRandomCoin(r)
		if coins == nil || coins.AmountOf(coin.Denom).IsZero() {
			coins = append(coins, coin)
		}
	}
	
	return coins.Sort()
}

func getRandomBandwidth(r *rand.Rand) hub.Bandwidth {
	upload := r.Int63n(hub.GB.Int64())
	download := r.Int63n(hub.GB.Int64())
	
	return hub.NewBandwidthFromInt64(upload, download)
}

func GenerateRandomNode(r *rand.Rand) types.Node {
	node := types.Node{
		ID:               getRandomNodeID(r),
		Owner:            nil,
		Deposit:          getRandomCoin(r),
		Type:             getRandomType(r),
		Version:          getRandomVersion(r),
		Moniker:          getRandomMoniker(r),
		PricesPerGB:      getRandomCoins(r),
		InternetSpeed:    getRandomBandwidth(r),
		Encryption:       getRandomEncryption(r),
		Status:           getRandomStatus(r),
		StatusModifiedAt: 0,
	}
	return node
}

func GenerateRandomSubscription(r *rand.Rand, node types.Node) types.Subscription {
	subscription := types.Subscription{
		ID:                 getRandomSubscriptionID(r),
		NodeID:             node.ID,
		Client:             nil,
		PricePerGB:         getRandomCoin(r),
		TotalDeposit:       getRandomCoin(r),
		RemainingDeposit:   getRandomCoin(r),
		RemainingBandwidth: getRandomBandwidth(r),
		Status:             getRandomStatus(r),
		StatusModifiedAt:   0,
	}
	
	return subscription
}

func GenerateRandomSession(r *rand.Rand, id hub.SubscriptionID) types.Session {
	session := types.Session{
		ID:               getRandomSessionID(r),
		SubscriptionID:   id,
		Bandwidth:        getRandomBandwidth(r),
		Status:           getRandomStatus(r),
		StatusModifiedAt: 0,
	}
	return session
}
