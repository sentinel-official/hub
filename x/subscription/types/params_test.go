package types

import (
	"testing"
	"time"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		InactiveDuration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"negative inactive duration",
			fields{
				InactiveDuration: -1000,
			},
			true,
		},
		{
			"zero inactive duration",
			fields{
				InactiveDuration: 0,
			},
			true,
		},
		{
			"positive inactive duration",
			fields{
				InactiveDuration: 1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				InactiveDuration: tt.fields.InactiveDuration,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
