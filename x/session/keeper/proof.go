package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/session/types"
)

func (k *Keeper) VerifyProof(ctx sdk.Context, address sdk.AccAddress, proof types.Proof, signature []byte) error {
	account := k.GetAccount(ctx, address)
	if account == nil {
		return fmt.Errorf("account for address %s does not exist", address)
	}
	if account.GetPubKey() == nil {
		return fmt.Errorf("public key for address %s does not exist", address)
	}

	bytes, err := proof.Marshal()
	if err != nil {
		return err
	}

	if !account.GetPubKey().VerifySignature(bytes, signature) {
		return fmt.Errorf("either message or signature is invalid")
	}

	return nil
}
