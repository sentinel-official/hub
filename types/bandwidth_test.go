package types

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBandwidth_Add(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	type args struct {
		v Bandwidth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bandwidth
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
			},
			Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
			},
			Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.Add(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_CeilTo(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	type args struct {
		precision sdk.Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bandwidth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.CeilTo(tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CeilTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAllLTE(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	type args struct {
		v Bandwidth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			false,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			false,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAllLTE(tt.args.v); got != tt.want {
				t.Errorf("IsAllLTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAllPositive(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(1000),
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(1000),
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(1000),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAllPositive(); got != tt.want {
				t.Errorf("IsAllPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAllZero(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(1000),
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(1000),
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAllZero(); got != tt.want {
				t.Errorf("IsAllZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAnyGT(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	type args struct {
		v Bandwidth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			true,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			false,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			true,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAnyGT(tt.args.v); got != tt.want {
				t.Errorf("IsAnyGT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAnyNegative(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(0),
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(1000),
			},
			true,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(1000),
			},
			false,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(0),
			},
			false,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAnyNegative(); got != tt.want {
				t.Errorf("IsAnyNegative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_IsAnyZero(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(0),
			},
			true,
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(1000),
			},
			false,
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(-1000),
			},
			true,
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			true,
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(1000),
			},
			true,
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(-1000),
			},
			false,
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(0),
			},
			true,
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(1000),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.IsAnyZero(); got != tt.want {
				t.Errorf("IsAnyZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_Sub(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	type args struct {
		v Bandwidth
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Bandwidth
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
			},
			Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
			},
			Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
			},
			Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
			},
			Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
			},
			Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(0)},
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(1000)},
			},
			Bandwidth{Upload: sdk.NewInt(0), Download: sdk.NewInt(-1000)},
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(-1000)},
			},
			Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(1000)},
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(0)},
			},
			Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(0)},
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			args{
				Bandwidth{Upload: sdk.NewInt(1000), Download: sdk.NewInt(1000)},
			},
			Bandwidth{Upload: sdk.NewInt(-1000), Download: sdk.NewInt(-1000)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.Sub(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBandwidth_Sum(t *testing.T) {
	type fields struct {
		Upload   sdk.Int
		Download sdk.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   sdk.Int
	}{
		{
			"negative upload and negative download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(-1000),
			},
			sdk.NewInt(-2000),
		},
		{
			"negative upload and zero download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(0),
			},
			sdk.NewInt(-1000),
		},
		{
			"negative upload and positive download",
			fields{
				Upload:   sdk.NewInt(-1000),
				Download: sdk.NewInt(1000),
			},
			sdk.NewInt(1).Add(sdk.NewInt(-1)),
		},
		{
			"zero upload and negative download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(-1000),
			},
			sdk.NewInt(-1000),
		},
		{
			"zero upload and zero download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(0),
			},
			sdk.NewInt(0),
		},
		{
			"zero upload and positive download",
			fields{
				Upload:   sdk.NewInt(0),
				Download: sdk.NewInt(1000),
			},
			sdk.NewInt(1000),
		},
		{
			"positive upload and negative download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(-1000),
			},
			sdk.NewInt(1).Add(sdk.NewInt(-1)),
		},
		{
			"positive upload and zero download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(0),
			},
			sdk.NewInt(1000),
		},
		{
			"positive upload and positive download",
			fields{
				Upload:   sdk.NewInt(1000),
				Download: sdk.NewInt(1000),
			},
			sdk.NewInt(2000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bandwidth{
				Upload:   tt.fields.Upload,
				Download: tt.fields.Download,
			}
			if got := b.Sum(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBandwidth(t *testing.T) {
	type args struct {
		upload   sdk.Int
		download sdk.Int
	}
	tests := []struct {
		name string
		args args
		want Bandwidth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBandwidth(tt.args.upload, tt.args.download); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBandwidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBandwidthFromInt64(t *testing.T) {
	type args struct {
		upload   int64
		download int64
	}
	tests := []struct {
		name string
		args args
		want Bandwidth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBandwidthFromInt64(tt.args.upload, tt.args.download); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBandwidthFromInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
