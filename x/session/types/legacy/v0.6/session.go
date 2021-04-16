package v0_6

import (
	hubtypes "github.com/sentinel-official/hub/types/legacy/v0.6"
	"github.com/sentinel-official/hub/x/session/types"
	legacy "github.com/sentinel-official/hub/x/session/types/legacy/v0.5"
)

func MigrateSession(item legacy.Session) types.Session {
	return types.Session{
		Id:           item.ID,
		Subscription: item.Subscription,
		Node:         item.Node.String(),
		Address:      item.Address.String(),
		Duration:     item.Duration,
		Bandwidth:    hubtypes.MigrateBandwidth(item.Bandwidth),
		Status:       hubtypes.MigrateStatus(item.Status),
		StatusAt:     item.StatusAt,
	}
}

func MigrateSessions(items legacy.Sessions) types.Sessions {
	var sessions types.Sessions
	for _, item := range items {
		sessions = append(sessions, MigrateSession(item))
	}

	return sessions
}
