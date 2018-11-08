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

func (k Keeper) SetVpnDetails(ctx sdkTypes.Context, vpnDetails hubTypes.VpnDetails, vpnId string) error {

	vpnStore := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		return err
	}

	vpnStore.Set(vpnIdBytes, vpnDetailsBytes)

	return nil
}

func (k Keeper) GetVpnDetails(ctx sdkTypes.Context, vpnId string) (hubTypes.VpnDetails, error) {
	var Details hubTypes.VpnDetails

	store := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes := store.Get(vpnIdBytes)
	err := json.Unmarshal(vpnDetailsBytes, &Details)

	if err != nil {
		return hubTypes.VpnDetails{}, err
	}

	return Details, nil
}
