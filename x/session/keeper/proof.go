package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func (k *Keeper) VerifyProof(ctx sdk.Context, address sdk.AccAddress, proof types.Proof, signature []byte) error {
	account := k.GetAccount(ctx, address)
	if account == nil {
		return types.ErrorAccountDoesNotExist
	}
	if account.GetPubKey() == nil {
		return types.ErrorPublicKeyDoesNotExist
	}

	bytes, err := proof.Marshal()
	if err != nil {
		return err
	}

	if !account.GetPubKey().VerifySignature(bytes, signature) {
		return types.ErrorInvalidSignature
	}

	return nil
}
