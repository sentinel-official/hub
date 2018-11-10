package vpn

import (
	"encoding/json"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type Keeper struct {
	VpnStoreKey sdkTypes.StoreKey
	IbcStoreKey sdkTypes.StoreKey
}

func NewKeeper(vpnKey sdkTypes.StoreKey, ibcKey sdkTypes.StoreKey) Keeper {

	return Keeper{
		VpnStoreKey: vpnKey,
		IbcStoreKey: ibcKey,
	}
}

func (k Keeper) SetVpnDetails(ctx sdkTypes.Context, vpnId sdkTypes.AccAddress, vpnDetails hubTypes.VpnDetails) error {

	vpnStore := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)

	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		return err
	}

	vpnStore.Set(vpnIdBytes, vpnDetailsBytes)

	return nil
}

func (k Keeper) GetVpnDetails(ctx sdkTypes.Context, vpnId sdkTypes.AccAddress) ([]byte, error) {

	store := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes := store.Get(vpnIdBytes)

	return vpnDetailsBytes, nil
}

func (k Keeper) SetAliveNode(ctx sdkTypes.Context, vpnId sdkTypes.AccAddress, Details hubTypes.VpnDetails) error {
	store := ctx.KVStore(k.VpnStoreKey)

	Details.Info.Status = true
	Details.Info.BlockHeight = ctx.BlockHeight()

	DetailsBytes, err := json.Marshal(Details)

	if err != nil {
		panic(err)
	}

	store.Set(vpnId, DetailsBytes)

	return nil
}
