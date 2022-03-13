package types

import (
	"reflect"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	hubtypes "github.com/sentinel-official/hub/types"
)

func TestSubscription_GetNode(t *testing.T) {
	type fields struct {
		Node string
	}
	tests := []struct {
		name   string
		fields fields
		want   hubtypes.NodeAddress
	}{
		{
			"empty",
			fields{
				Node: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Node: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			hubtypes.NodeAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Subscription{
				Node: tt.fields.Node,
			}
			if got := m.GetNode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscription_GetOwner(t *testing.T) {
	type fields struct {
		Owner string
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.AccAddress
	}{
		{
			"empty",
			fields{
				Owner: "",
			},
			nil,
		},
		{
			"20 bytes",
			fields{
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			sdk.AccAddress{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Subscription{
				Owner: tt.fields.Owner,
			}
			if got := m.GetOwner(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOwner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscription_Validate(t *testing.T) {
	type fields struct {
		Id       uint64
		Owner    string
		Node     string
		Price    sdk.Coin
		Deposit  sdk.Coin
		Plan     uint64
		Denom    string
		Expiry   time.Time
		Free     sdk.Int
		Status   hubtypes.Status
		StatusAt time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"zero id",
			fields{
				Id: 0,
			},
			true,
		},
		{
			"positive id",
			fields{
				Id: 1000,
			},
			true,
		},
		{
			"empty owner",
			fields{
				Id:    1000,
				Owner: "",
			},
			true,
		},
		{
			"invalid owner",
			fields{
				Id:    1000,
				Owner: "invalid",
			},
			true,
		},
		{
			"invalid prefix owner",
			fields{
				Id:    1000,
				Owner: "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
			},
			true,
		},
		{
			"10 bytes owner",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgslawd5s",
			},
			true,
		},
		{
			"20 bytes owner",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"30 bytes owner",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfqyy3zxfp9ycnjs2fszvfck8",
			},
			true,
		},
		{
			"empty node",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "",
			},
			true,
		},
		{
			"invalid node",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "invalid",
			},
			true,
		},
		{
			"invalid prefix node",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
			},
			true,
		},
		{
			"20 bytes node",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price: sdk.Coin{Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"empty price",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price: sdk.Coin{Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"empty denom price",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price: sdk.Coin{Denom: "", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"invalid denom price",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price: sdk.Coin{Denom: "o", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative amount price",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero amount price",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:  "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive amount price",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"empty deposit",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"empty denom deposit",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"invalid denom deposit",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "o", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative amount deposit",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero amount deposit",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive amount deposit",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Free:    sdk.NewInt(0),
			},
			true,
		},
		{
			"zero plan",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Plan:    0,
				Free:    sdk.NewInt(0),
			},
			true,
		},
		{
			"positive plan",
			fields{
				Id:      1000,
				Owner:   "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Node:    "sentnode1qypqxpq9qcrsszgszyfpx9q4zct3sxfqelr5ey",
				Price:   sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Deposit: sdk.Coin{Denom: "one", Amount: sdk.NewInt(1000)},
				Plan:    1000,
			},
			true,
		},
		{
			"empty denom",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:  1000,
				Denom: "",
			},
			true,
		},
		{
			"invalid denom",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:  1000,
				Denom: "o",
			},
			true,
		},
		{
			"one denom",
			fields{
				Id:    1000,
				Owner: "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:  1000,
				Denom: "one",
			},
			true,
		},
		{
			"zero expiry",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Time{},
			},
			true,
		},
		{
			"now expiry",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(0),
			},
			true,
		},
		{
			"negative free",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero free",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(0),
			},
			true,
		},
		{
			"positive free",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(1000),
			},
			true,
		},
		{
			"unknown status",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(1000),
				Status: hubtypes.StatusUnknown,
			},
			true,
		},
		{
			"inactive status",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(1000),
				Status: hubtypes.StatusInactive,
			},
			true,
		},
		{
			"inactive pending status",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(1000),
				Status: hubtypes.StatusInactivePending,
			},
			true,
		},
		{
			"active status",
			fields{
				Id:     1000,
				Owner:  "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:   1000,
				Denom:  "one",
				Expiry: time.Now(),
				Free:   sdk.NewInt(1000),
				Status: hubtypes.StatusActive,
			},
			true,
		},
		{
			"zero status_at",
			fields{
				Id:       1000,
				Owner:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:     1000,
				Denom:    "one",
				Expiry:   time.Now(),
				Free:     sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Time{},
			},
			true,
		},
		{
			"now status_at",
			fields{
				Id:       1000,
				Owner:    "sent1qypqxpq9qcrsszgszyfpx9q4zct3sxfq0fzduj",
				Plan:     1000,
				Denom:    "one",
				Expiry:   time.Now(),
				Free:     sdk.NewInt(1000),
				Status:   hubtypes.StatusActive,
				StatusAt: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Subscription{
				Id:       tt.fields.Id,
				Owner:    tt.fields.Owner,
				Node:     tt.fields.Node,
				Price:    tt.fields.Price,
				Deposit:  tt.fields.Deposit,
				Plan:     tt.fields.Plan,
				Denom:    tt.fields.Denom,
				Expiry:   tt.fields.Expiry,
				Free:     tt.fields.Free,
				Status:   tt.fields.Status,
				StatusAt: tt.fields.StatusAt,
			}
			if err := m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
