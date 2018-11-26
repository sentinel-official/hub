package vpn

import (
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper struct {
	VPNStoreKey     csdkTypes.StoreKey
	SessionStoreKey csdkTypes.StoreKey
	cdc             *codec.Codec
}

func NewKeeper(cdc *codec.Codec, vpnKey, sessionKey csdkTypes.StoreKey) Keeper {
	return Keeper{
		VPNStoreKey:     vpnKey,
		SessionStoreKey: sessionKey,
		cdc:             cdc,
	}
}

func (k Keeper) SetVPNDetails(ctx csdkTypes.Context, vpnID string, vpnDetails *sdkTypes.VPNDetails) {
	vpnStore := ctx.KVStore(k.VPNStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(vpnID)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(vpnDetails)

	if err != nil {
		panic(err)
	}

	vpnStore.Set(keyBytes, valueBytes)

	k.SetVPNsCount(ctx, k.GetVPNsCount(ctx)+1)
}

func (k Keeper) GetVPNDetails(ctx csdkTypes.Context, vpnID string) *sdkTypes.VPNDetails {
	store := ctx.KVStore(k.VPNStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(vpnID)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil
	}

	var vpnDetails sdkTypes.VPNDetails

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &vpnDetails); err != nil {
		panic(err)
	}

	return &vpnDetails
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, nodeIDs []string) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("ACTIVE_NODE_IDS")

	if err != nil {
		panic(err)
	}

	sort.Strings(nodeIDs)
	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(nodeIDs)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetActiveNodeIDs(ctx csdkTypes.Context) []string {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("ACTIVE_NODE_IDS")

	if err != nil {
		panic(err)
	}

	var nodeIDs []string
	store := ctx.KVStore(k.VPNStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nodeIDs
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &nodeIDs); err != nil {
		panic(err)
	}

	return nodeIDs
}

func (k Keeper) SetVPNsCount(ctx csdkTypes.Context, count uint64) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("VPNS_COUNT")

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(count)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetVPNsCount(ctx csdkTypes.Context) uint64 {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("VPNS_COUNT")

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.VPNStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0
	}

	var count uint64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &count); err != nil {
		panic(err)
	}

	return count
}

func (k Keeper) SetSessionDetails(ctx csdkTypes.Context, sessionID string, sessionDetails *sdkTypes.SessionDetails) {
	store := ctx.KVStore(k.SessionStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionID)

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionDetails)

	if err != nil {
		panic(err)
	}

	store.Set(keyBytes, valueBytes)

	k.SetSessionsCount(ctx, k.GetSessionsCount(ctx)+1)
}

func (k Keeper) GetSessionDetails(ctx csdkTypes.Context, sessionID string) *sdkTypes.SessionDetails {
	store := ctx.KVStore(k.SessionStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionID)

	if err != nil {
		panic(err)
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil
	}

	var sessionDetails sdkTypes.SessionDetails

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &sessionDetails); err != nil {
		panic(err)
	}

	return &sessionDetails
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, sessionIDs []string) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("ACTIVE_SESSION_IDS")

	if err != nil {
		panic(err)
	}

	sort.Strings(sessionIDs)
	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionIDs)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetActiveSessionIDs(ctx csdkTypes.Context) []string {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("ACTIVE_SESSION_IDS")

	if err != nil {
		panic(err)
	}

	var sessionIDs []string
	store := ctx.KVStore(k.SessionStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return sessionIDs
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &sessionIDs); err != nil {
		panic(err)
	}

	return sessionIDs
}

func (k Keeper) SetSessionsCount(ctx csdkTypes.Context, count uint64) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("SESSIONS_COUNT")

	if err != nil {
		panic(err)
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(count)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(keyBytes, valueBytes)
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context) uint64 {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed("SESSIONS_COUNT")

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.SessionStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0
	}

	var count uint64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &count); err != nil {
		panic(err)
	}

	return count
}

func (k Keeper) SetVPNStatus(ctx csdkTypes.Context, vpnID string, status string) {
	vpnDetails := k.GetVPNDetails(ctx, vpnID)
	vpnDetails.Status = status

	k.SetVPNDetails(ctx, vpnID, vpnDetails)
}

func (k Keeper) AddActiveNodeID(ctx csdkTypes.Context, nodeID string) {
	nodeIDs := k.GetActiveNodeIDs(ctx)
	nodeIDs = append(nodeIDs, nodeID)
	k.SetActiveSessionIDs(ctx, nodeIDs)
}

func (k Keeper) RemoveActiveNodeID(ctx csdkTypes.Context, nodeID string) {
	oldNodeIDs := k.GetActiveNodeIDs(ctx)
	var nodeIDs []string

	for _, id := range oldNodeIDs {
		if id != nodeID {
			nodeIDs = append(nodeIDs, id)
		}
	}

	k.SetActiveNodeIDs(ctx, nodeIDs)
}

func (k Keeper) SetSessionStatus(ctx csdkTypes.Context, sessionID string, status string) {
	sessionDetails := k.GetSessionDetails(ctx, sessionID)
	sessionDetails.Status = status

	k.SetSessionDetails(ctx, sessionID, sessionDetails)
}

func (k Keeper) AddActiveSessionID(ctx csdkTypes.Context, sessionID string) {
	sessionIDs := k.GetActiveSessionIDs(ctx)
	sessionIDs = append(sessionIDs, sessionID)
	k.SetActiveSessionIDs(ctx, sessionIDs)
}

func (k Keeper) RemoveActiveSessionID(ctx csdkTypes.Context, sessionID string) {
	oldSessionIDs := k.GetActiveSessionIDs(ctx)
	var sessionIDs []string

	for _, id := range oldSessionIDs {
		if id != sessionID {
			sessionIDs = append(sessionIDs, id)
		}
	}

	k.SetActiveSessionIDs(ctx, sessionIDs)
}
