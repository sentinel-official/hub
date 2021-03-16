package types

import (
	"fmt"
)

const (
	AttributeKeyTxHash  = "tx_hash"
	AttributeKeyAddress = "address"
	AttributeKeyAmount  = "amount"
)

var (
	EventTypeSet = fmt.Sprintf("%s:set", ModuleName)
)
