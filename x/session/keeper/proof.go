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

	pubKey := account.GetPubKey()
	if pubKey == nil {
		return fmt.Errorf("public key for address %s does not exist", address)
	}

	message, err := proof.Marshal()
	if err != nil {
		return err
	}

	if !pubKey.VerifySignature(message, signature) {
		return fmt.Errorf("invalid signature")
	}

	return nil
}
