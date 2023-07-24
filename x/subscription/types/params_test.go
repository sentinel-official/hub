package types

import (
	"testing"
	"time"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		InactivePendingDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"negative inactive_pending_duration",
			fields{
				InactivePendingDuration: -1000,
			},
			true,
		},
		{
			"zero inactive_pending_duration",
			fields{
				InactivePendingDuration: 0,
			},
			true,
		},
		{
			"positive inactive_pending_duration",
			fields{
				InactivePendingDuration: 1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				InactivePendingDuration: tt.fields.InactivePendingDuration,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
