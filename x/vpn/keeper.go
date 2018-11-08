package vpn

import (
	"encoding/json"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	vpnTypes "github.com/ironman0x7b2/sentinel-hub/types"
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

func (k Keeper) SetVpnDetails(ctx sdkTypes.Context, vpnDetails vpnTypes.Registervpn, vpnId string) (string, sdkTypes.Error) {
	vpnStore := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes, err := json.Marshal(vpnDetails)
	if err != nil {
		return "", sdkTypes.ErrInternal("Marshal vpn details failed")
	}
	vpnStore.Set(vpnIdBytes, vpnDetailsBytes)

	return vpnId, nil
}

func (k Keeper) GetVpnDetails(ctx sdkTypes.Context, vpnId string) (vpnDetails vpnTypes.Registervpn, error sdkTypes.Error) {
	store := ctx.KVStore(k.VpnStoreKey)
	vpnIdBytes := []byte(vpnId)
	vpnDetailsBytes := store.Get(vpnIdBytes)
	var Details vpnTypes.Registervpn
	err := json.Unmarshal(vpnDetailsBytes, &Details)
	if err != nil {
		return vpnTypes.Registervpn{}, sdkTypes.ErrInternal("Unmarshaling vpn detail failed")
	}
	return Details, nil
}
