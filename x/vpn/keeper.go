package vpn

import (
	"encoding/json"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	hubTypes "github.com/ironman0x7b2/sentinel-hub/types"
)

type Keeper struct {
	VpnStoreKey sdkTypes.StoreKey
}

func NewKeeper(vpnKey sdkTypes.StoreKey) Keeper {

	return Keeper{
		VpnStoreKey: vpnKey,
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

func (k Keeper) SetVpnStatus(ctx sdkTypes.Context, vpnId sdkTypes.AccAddress, vpnDetails hubTypes.VpnDetails) error {
	store := ctx.KVStore(k.VpnStoreKey)

	vpnDetails.Info.Status = true
	vpnDetails.Info.BlockHeight = ctx.BlockHeight()

	DetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		panic(err)
	}

	store.Set(vpnId, DetailsBytes)

	return nil
}
