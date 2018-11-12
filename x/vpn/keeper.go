package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper struct {
	VPNStoreKey csdkTypes.StoreKey
}

func NewKeeper(vpnKey csdkTypes.StoreKey) Keeper {
	return Keeper{
		VPNStoreKey: vpnKey,
	}
}

func (k Keeper) SetVPNDetails(ctx csdkTypes.Context, vpnId string, vpnDetails sdkTypes.VPNDetails) {
	vpnStore := ctx.KVStore(k.VPNStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		panic(err)
	}

	vpnStore.Set(vpnIdBytes, vpnDetailsBytes)
}

func (k Keeper) GetVPNDetails(ctx csdkTypes.Context, vpnId string) *sdkTypes.VPNDetails {
	store := ctx.KVStore(k.VPNStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes := store.Get(vpnIdBytes)

	if vpnDetailsBytes == nil {
		return nil
	}

	var vpnDetails sdkTypes.VPNDetails

	if err := json.Unmarshal(vpnDetailsBytes, &vpnDetails); err != nil {
		panic(err)
	}

	return &vpnDetails
}

func (k Keeper) SetVPNStatus(ctx csdkTypes.Context, vpnId string, status bool) {
	vpnDetails := k.GetVPNDetails(ctx, vpnId)
	vpnDetails.Info.Status = status
	vpnDetails.Info.BlockHeight = ctx.BlockHeight()

	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		panic(err)
	}

	store := ctx.KVStore(k.VPNStoreKey)
	store.Set(vpnIdBytes, vpnDetailsBytes)
}
