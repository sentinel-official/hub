package types

import (
	"strings"
	"testing"

	hub "github.com/sentinel-official/hub/types"
)

func TestProvider_Validate(t *testing.T) {
	tests := []struct {
		name    string
		p       Provider
		wantErr bool
	}{
		{"address nil", Provider{nil, "", "", "", ""}, true},
		{"address zero", Provider{hub.ProvAddress{}, "", "", "", ""}, true},
		{"address empty", Provider{hub.ProvAddress(""), "", "", "", ""}, true},
		{"name length 0", Provider{hub.ProvAddress("address-1"), "", "", "", ""}, true},
		{"name length greater than 64", Provider{hub.ProvAddress("address-1"), strings.Repeat("-", 64+1), "", "", ""}, true},
		{"identity length greater than 64", Provider{hub.ProvAddress("address-1"), "name", strings.Repeat("-", 64+1), "", ""}, true},
		{"website length greater than 64", Provider{hub.ProvAddress("address-1"), "name", "", strings.Repeat("-", 64+1), ""}, true},
		{"description length greater than 256", Provider{hub.ProvAddress("address-1"), "name", "", "", strings.Repeat("-", 256+1)}, true},
		{"valid", Provider{hub.ProvAddress("address-1"), "name", "", "", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.p.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
