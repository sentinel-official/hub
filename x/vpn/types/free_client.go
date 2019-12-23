package types

import (
	"fmt"
	
	sdk "github.com/cosmos/cosmos-sdk/types"
	
	hub "github.com/sentinel-official/hub/types"
)

type FreeClient struct {
	NodeID hub.NodeID
	Client sdk.AccAddress
}

func (fc FreeClient) String() string {
	return fmt.Sprintf(`
	NodeID : %s
	Client : %s
`, fc.NodeID.String(), fc.Client.String())
}
