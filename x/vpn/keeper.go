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

func (k Keeper) SetVPNDetails(ctx csdkTypes.Context, vpnID string, vpnDetails *sdkTypes.VPNDetails) csdkTypes.Error {
	vpnStore := ctx.KVStore(k.VPNStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(vpnID)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(vpnDetails)

	if err != nil {
		return errorMarshal()
	}

	vpnStore.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetVPNDetails(ctx csdkTypes.Context, vpnID string) (*sdkTypes.VPNDetails, csdkTypes.Error) {
	store := ctx.KVStore(k.VPNStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(vpnID)

	if err != nil {
		return nil, errorMarshal()
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil, nil
	}

	var vpnDetails sdkTypes.VPNDetails

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &vpnDetails); err != nil {
		return nil, errorUnmarshal()
	}

	return &vpnDetails, nil
}

func (k Keeper) SetActiveNodeIDs(ctx csdkTypes.Context, nodeIDs []string) csdkTypes.Error {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveNodeIDs)

	if err != nil {
		return errorMarshal()
	}

	sort.Strings(nodeIDs)
	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(nodeIDs)

	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetActiveNodeIDs(ctx csdkTypes.Context) ([]string, csdkTypes.Error) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveNodeIDs)

	if err != nil {
		return nil, errorMarshal()
	}

	var nodeIDs []string
	store := ctx.KVStore(k.VPNStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nodeIDs, nil
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &nodeIDs); err != nil {
		return nil, errorUnmarshal()
	}

	return nodeIDs, nil
}

func (k Keeper) SetVPNsCount(ctx csdkTypes.Context, count uint64) csdkTypes.Error {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyVPNsCount)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(count)

	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetVPNsCount(ctx csdkTypes.Context) (uint64, csdkTypes.Error) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyVPNsCount)

	if err != nil {
		return 0, errorMarshal()
	}

	store := ctx.KVStore(k.VPNStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0, nil
	}

	var count uint64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &count); err != nil {
		return 0, errorUnmarshal()
	}

	return count, nil
}

func (k Keeper) SetSessionDetails(ctx csdkTypes.Context, sessionID string, sessionDetails *sdkTypes.SessionDetails) csdkTypes.Error {
	store := ctx.KVStore(k.SessionStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionID)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(*sessionDetails)

	if err != nil {
		return errorMarshal()
	}

	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetSessionDetails(ctx csdkTypes.Context, sessionID string) (*sdkTypes.SessionDetails, csdkTypes.Error) {
	store := ctx.KVStore(k.SessionStoreKey)
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionID)

	if err != nil {
		return nil, errorMarshal()
	}

	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return nil, nil
	}

	var sessionDetails sdkTypes.SessionDetails

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &sessionDetails); err != nil {
		return nil, errorUnmarshal()
	}

	return &sessionDetails, nil
}

func (k Keeper) SetActiveSessionIDs(ctx csdkTypes.Context, sessionIDs []string) csdkTypes.Error {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveSessionIDs)

	if err != nil {
		return errorMarshal()
	}

	sort.Strings(sessionIDs)
	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sessionIDs)

	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetActiveSessionIDs(ctx csdkTypes.Context) ([]string, csdkTypes.Error) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeyActiveSessionIDs)

	if err != nil {
		return nil, errorMarshal()
	}

	var sessionIDs []string
	store := ctx.KVStore(k.SessionStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return sessionIDs, nil
	}

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &sessionIDs); err != nil {
		return nil, errorUnmarshal()
	}

	return sessionIDs, nil
}

func (k Keeper) SetSessionsCount(ctx csdkTypes.Context, count uint64) csdkTypes.Error {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeySessionsCount)

	if err != nil {
		return errorMarshal()
	}

	valueBytes, err := k.cdc.MarshalBinaryLengthPrefixed(count)

	if err != nil {
		return errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	store.Set(keyBytes, valueBytes)

	return nil
}

func (k Keeper) GetSessionsCount(ctx csdkTypes.Context) (uint64, csdkTypes.Error) {
	keyBytes, err := k.cdc.MarshalBinaryLengthPrefixed(sdkTypes.KeySessionsCount)

	if err != nil {
		return 0, errorMarshal()
	}

	store := ctx.KVStore(k.SessionStoreKey)
	valueBytes := store.Get(keyBytes)

	if valueBytes == nil {
		return 0, nil
	}

	var count uint64

	if err := k.cdc.UnmarshalBinaryLengthPrefixed(valueBytes, &count); err != nil {
		return 0, errorUnmarshal()
	}

	return count, nil
}

func (k Keeper) SetVPNStatus(ctx csdkTypes.Context, vpnID string, status string) csdkTypes.Error {
	vpnDetails, err := k.GetVPNDetails(ctx, vpnID)

	if err != nil {
		return err
	}

	vpnDetails.Status = status

	if err := k.SetVPNDetails(ctx, vpnID, vpnDetails); err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddActiveNodeID(ctx csdkTypes.Context, nodeID string) csdkTypes.Error {
	nodeIDs, err := k.GetActiveNodeIDs(ctx)

	if err != nil {
		return err
	}

	nodeIDs = append(nodeIDs, nodeID)

	if err := k.SetActiveNodeIDs(ctx, nodeIDs); err != nil {
		return err
	}

	return nil
}

func (k Keeper) RemoveActiveNodeID(ctx csdkTypes.Context, nodeID string) csdkTypes.Error {
	oldNodeIDs, err := k.GetActiveNodeIDs(ctx)

	if err != nil {
		return err
	}

	var nodeIDs []string

	for _, id := range oldNodeIDs {
		if id != nodeID {
			nodeIDs = append(nodeIDs, id)
		}
	}

	if err := k.SetActiveNodeIDs(ctx, nodeIDs); err != nil {
		return err
	}

	return nil
}

func (k Keeper) SetSessionStatus(ctx csdkTypes.Context, sessionID string, status string) csdkTypes.Error {
	sessionDetails, err := k.GetSessionDetails(ctx, sessionID)

	if err != nil {
		return err
	}
	sessionDetails.Status = status

	if err := k.SetSessionDetails(ctx, sessionID, sessionDetails); err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddActiveSessionID(ctx csdkTypes.Context, sessionID string) csdkTypes.Error {
	sessionIDs, err := k.GetActiveSessionIDs(ctx)

	if err != nil {
		return err
	}

	sessionIDs = append(sessionIDs, sessionID)

	if err := k.SetActiveSessionIDs(ctx, sessionIDs); err != nil {
		return err
	}

	return nil
}

func (k Keeper) RemoveActiveSessionID(ctx csdkTypes.Context, sessionID string) csdkTypes.Error {
	oldSessionIDs, err := k.GetActiveSessionIDs(ctx)

	if err != nil {
		return err
	}

	var sessionIDs []string

	for _, id := range oldSessionIDs {
		if id != sessionID {
			sessionIDs = append(sessionIDs, id)
		}
	}

	if err := k.SetActiveSessionIDs(ctx, sessionIDs); err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddVPN(ctx csdkTypes.Context, vpnID string, vpnDetails *sdkTypes.VPNDetails) csdkTypes.Error {
	if err := k.SetVPNDetails(ctx, vpnID, vpnDetails); err != nil {
		return err
	}

	vpnsCount, err := k.GetVPNsCount(ctx)

	if err != nil {
		return err
	}

	if err := k.SetVPNsCount(ctx, vpnsCount+1); err != nil {
		return err
	}

	return nil
}

func (k Keeper) AddSession(ctx csdkTypes.Context, sessionID string, sessionDetails *sdkTypes.SessionDetails) csdkTypes.Error {
	if err := k.SetSessionDetails(ctx, sessionID, sessionDetails); err != nil {
		return err
	}

	sessionsCount, err := k.GetSessionsCount(ctx)

	if err != nil {
		return err
	}

	if err := k.SetSessionsCount(ctx, sessionsCount+1); err != nil {
		return err
	}

	return nil
}
