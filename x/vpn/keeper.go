package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

type Keeper struct {
	VPNStoreKey csdkTypes.StoreKey
	Account     auth.AccountKeeper
}

func NewKeeper(vpnKey csdkTypes.StoreKey, ak auth.AccountKeeper) Keeper {
	return Keeper{
		VPNStoreKey: vpnKey,
		Account:     ak,
	}
}

func (k Keeper) SetVPNDetails(ctx csdkTypes.Context, vpnID string, vpnDetails sdkTypes.VPNDetails) {
	vpnStore := ctx.KVStore(k.VPNStoreKey)
	vpnIDBytes := []byte(vpnID)
	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		panic(err)
	}

	vpnStore.Set(vpnIDBytes, vpnDetailsBytes)
}

func (k Keeper) GetVPNDetails(ctx csdkTypes.Context, vpnID string) *sdkTypes.VPNDetails {
	store := ctx.KVStore(k.VPNStoreKey)
	vpnIDBytes := []byte(vpnID)
	vpnDetailsBytes := store.Get(vpnIDBytes)

	if vpnDetailsBytes == nil {
		return nil
	}

	var vpnDetails sdkTypes.VPNDetails

	if err := json.Unmarshal(vpnDetailsBytes, &vpnDetails); err != nil {
		panic(err)
	}

	return &vpnDetails
}

func (k Keeper) SetVPNStatus(ctx csdkTypes.Context, vpnID string, status bool) {
	vpnDetails := k.GetVPNDetails(ctx, vpnID)
	vpnDetails.Info.Status = status
	vpnDetails.Info.BlockHeight = ctx.BlockHeight()

	vpnIDBytes := []byte(vpnID)
	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(vpnIDBytes, vpnDetailsBytes)
}

func (k Keeper) SetSessionDetails(ctx csdkTypes.Context, session sdkTypes.Session, sessionKey string) {
	store := ctx.KVStore(k.VPNStoreKey)

	sessionData, err := json.Marshal(session)

	if err != nil {
		panic(err)
	}

	store.Set([]byte(sessionKey), sessionData)
}

func (k Keeper) GetSessionDetails(ctx csdkTypes.Context, sessionID string) *sdkTypes.Session {
	store := ctx.KVStore(k.VPNStoreKey)

	var details sdkTypes.Session
	sessionData := store.Get([]byte(sessionID))

	err := json.Unmarshal(sessionData, &details)
	if err != nil {
		panic(err)
	}

	return &details
}

func (k Keeper) SetSessionStatus(ctx csdkTypes.Context, sessionID string, status bool) {

	sessionDetails := k.GetSessionDetails(ctx, sessionID)
	sessionDetails.Status = status
	sessionDetails.StartTime = ctx.BlockHeader().Time

	sessionIDBytes := []byte(sessionID)
	sessionDetailsBytes, err := json.Marshal(sessionDetails)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(sessionDetailsBytes, sessionIDBytes)

	activeSessions := k.GetActiveSessions(ctx)
	k.SetActiveSessions(ctx, activeSessions, sessionID, status)
}

func (k Keeper) SetActiveSessions(ctx csdkTypes.Context, activeSessions []string, sessionID string, status bool) {
	if status {
		activeSessions := append(activeSessions, sessionID)

		activeSessionBytes, err := json.Marshal(activeSessions)
		if err != nil {
			panic(err)
		}

		store := ctx.KVStore(k.VPNStoreKey)
		store.Set([]byte("activeSessions"), activeSessionBytes)
	}

	if !status {
		var list []string
		for _, session := range activeSessions {
			if session == sessionID {
				continue
			} else {
				list = append(list, session)
			}
		}

		activeSessionListBytes, err := json.Marshal(list)

		if err != nil {
			panic(err)
		}
		store := ctx.KVStore(k.VPNStoreKey)
		store.Set([]byte("activeSessions"), activeSessionListBytes)
	}
}

func (k Keeper) GetActiveSessions(ctx csdkTypes.Context) []string {

	var activeSessions sdkTypes.ActiveSessions
	store := ctx.KVStore(k.VPNStoreKey)
	activeSessionDetails := store.Get([]byte("activeSessions"))

	err := json.Unmarshal(activeSessionDetails, activeSessions)
	if err != nil {
		panic(err)
	}

	return activeSessions
}
