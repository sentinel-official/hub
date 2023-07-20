package types

import (
	"testing"
	"time"
)

func TestParams_Validate(t *testing.T) {
	type fields struct {
		ExpiryDuration           time.Duration
		ProofVerificationEnabled bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"negative expiry duration",
			fields{
				ExpiryDuration: -1000,
			},
			true,
		},
		{
			"zero expiry duration",
			fields{
				ExpiryDuration: 0,
			},
			true,
		},
		{
			"positive expiry duration",
			fields{
				ExpiryDuration: 1000,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Params{
				ExpiryDuration:           tt.fields.ExpiryDuration,
				ProofVerificationEnabled: tt.fields.ProofVerificationEnabled,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
