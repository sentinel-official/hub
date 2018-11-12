package vpn

import (
	"encoding/json"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Keeper struct {
	VpnStoreKey csdkTypes.StoreKey
}

func NewKeeper(vpnKey csdkTypes.StoreKey) Keeper {

	return Keeper{
		VpnStoreKey: vpnKey,
	}
}

func (k Keeper) SetVpnDetails(ctx csdkTypes.Context, vpnId string, vpnDetails sdkTypes.VpnDetails) error {

	vpnStore := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)

	vpnDetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		return err
	}

	vpnStore.Set(vpnIdBytes, vpnDetailsBytes)

	return nil
}

func (k Keeper) GetVpnDetails(ctx csdkTypes.Context, vpnId string) ([]byte, error) {

	store := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes := store.Get(vpnIdBytes)

	return vpnDetailsBytes, nil
}

func (k Keeper) SetVpnStatus(ctx csdkTypes.Context, vpnId string, vpnDetails sdkTypes.VpnDetails, status bool) error {
	store := ctx.KVStore(k.VpnStoreKey)

	vpnDetails.Info.Status = status
	vpnDetails.Info.BlockHeight = ctx.BlockHeight()
	vpnIdBytes := []byte(vpnId)
	DetailsBytes, err := json.Marshal(vpnDetails)

	if err != nil {
		panic(err)
	}

	store.Set(vpnIdBytes, DetailsBytes)

	return nil
}
