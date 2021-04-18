package v0_6

import (
	hubtypes "github.com/sentinel-official/hub/types/legacy/v0.6"
	"github.com/sentinel-official/hub/x/node/types"
	legacy "github.com/sentinel-official/hub/x/node/types/legacy/v0.5"
)

func MigrateNode(item legacy.Node) types.Node {
	return types.Node{
		Address:   item.Address.String(),
		Provider:  item.Provider.String(),
		Price:     item.Price,
		RemoteURL: item.RemoteURL,
		Status:    hubtypes.MigrateStatus(item.Status),
		StatusAt:  item.StatusAt,
	}
}

func MigrateNodes(items legacy.Nodes) types.Nodes {
	var nodes types.Nodes
	for _, item := range items {
		nodes = append(nodes, MigrateNode(item))
	}

	return nodes
}
