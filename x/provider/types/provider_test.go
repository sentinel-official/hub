package types

import (
	hubtypes "github.com/sentinel-official/hub/types"
	"testing"
)

func TestProvider_Validate(t *testing.T) {
	correctAddress := hubtypes.ProvAddress{}

	// we dont need to check for contents of the error message, checking whether the error is nil/not-nil is enough.
	tests := []struct {
		name string
		p    Provider
	}{
		{"address zero", Provider{"", "", "", "", ""}},
		{"address empty", Provider{correctAddress.String(), "", "", "", ""}},
		{"name length zero", Provider{correctAddress.String(), "", "", "", ""}},
		{"name length greater than 64", Provider{correctAddress.String(), GT64, "", "", ""}},
		{"identity length greater than 64", Provider{correctAddress.String(), "name", GT64, "", ""}},
		{"website length greater than 64", Provider{correctAddress.String(), "name", "", GT64, ""}},
		{"description length greater than 256", Provider{correctAddress.String(), "name", "", "", GT256}},
		{"valid", Provider{correctAddress.String(), "name", "", "", ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.p.Validate()
			if err == nil && tt.name != "valid" {
				t.Errorf("Validate() error = %v", err)
			}
		})
	}
}
