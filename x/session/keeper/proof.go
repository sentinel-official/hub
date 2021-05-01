package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func (k *Keeper) VerifyProof(ctx sdk.Context, address sdk.AccAddress, proof types.Proof, signature []byte) error {
	account := k.GetAccount(ctx, address)
	if account == nil {
		return fmt.Errorf("account does not exist")
	}
	if account.GetPubKey() == nil {
		return fmt.Errorf("public key does not exist")
	}

	bytes, err := proof.Marshal()
	if err != nil {
		return err
	}

	if !account.GetPubKey().VerifySignature(bytes, signature) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
