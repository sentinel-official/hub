package v0_6

import (
	"github.com/sentinel-official/hub/x/swap/types"
	legacy "github.com/sentinel-official/hub/x/swap/types/legacy/v0.5"
)

func MigrateSwap(item legacy.Swap) types.Swap {
	return types.Swap{
		TxHash:   item.TxHash.Bytes(),
		Receiver: item.Receiver.String(),
		Amount:   item.Amount,
	}
}

func MigrateSwaps(items legacy.Swaps) types.Swaps {
	var swaps types.Swaps
	for _, item := range items {
		swaps = append(swaps, MigrateSwap(item))
	}

	return swaps
}
