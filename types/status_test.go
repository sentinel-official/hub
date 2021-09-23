package types

import (
	"testing"
)

func TestStatusFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Status
	}{
		{
			"empty",
			args{
				s: "",
			},
			StatusUnknown,
		},
		{
			"invalid",
			args{
				s: "invalid",
			},
			StatusUnknown,
		},
		{
			"unknown",
			args{
				s: "unknown",
			},
			StatusUnknown,
		},
		{
			"active",
			args{
				s: "Active",
			},
			StatusActive,
		},
		{
			"inactive pending",
			args{
				s: "InactivePending",
			},
			StatusInactivePending,
		},
		{
			"inactive",
			args{
				s: "Inactive",
			},
			StatusInactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StatusFromString(tt.args.s); got != tt.want {
				t.Errorf("StatusFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_Equal(t *testing.T) {
	type args struct {
		v Status
	}
	tests := []struct {
		name string
		s    Status
		args args
		want bool
	}{
		{
			"unknown and unknown",
			StatusUnknown,
			args{
				v: StatusUnknown,
			},
			true,
		},
		{
			"unknown and active",
			StatusUnknown,
			args{
				v: StatusActive,
			},
			false,
		},
		{
			"unknown and inactive pending",
			StatusUnknown,
			args{
				v: StatusInactivePending,
			},
			false,
		},
		{
			"unknown and inactive",
			StatusUnknown,
			args{
				v: StatusInactive,
			},
			false,
		},
		{
			"active and unknown",
			StatusActive,
			args{
				v: StatusUnknown,
			},
			false,
		},
		{
			"active and active",
			StatusActive,
			args{
				v: StatusActive,
			},
			true,
		},
		{
			"active and inactive pending",
			StatusActive,
			args{
				v: StatusInactivePending,
			},
			false,
		},
		{
			"active and inactive",
			StatusActive,
			args{
				v: StatusInactive,
			},
			false,
		},
		{
			"inactive pending and unknown",
			StatusInactivePending,
			args{
				v: StatusUnknown,
			},
			false,
		},
		{
			"inactive pending and active",
			StatusInactivePending,
			args{
				v: StatusActive,
			},
			false,
		},
		{
			"inactive pending and inactive pending",
			StatusInactivePending,
			args{
				v: StatusInactivePending,
			},
			true,
		},
		{
			"inactive pending and inactive",
			StatusInactivePending,
			args{
				v: StatusInactive,
			},
			false,
		},
		{
			"inactive and unknown",
			StatusInactive,
			args{
				v: StatusUnknown,
			},
			false,
		},
		{
			"inactive and active",
			StatusInactive,
			args{
				v: StatusActive,
			},
			false,
		},
		{
			"inactive and inactive pending",
			StatusInactive,
			args{
				v: StatusInactivePending,
			},
			false,
		},
		{
			"inactive and inactive",
			StatusInactive,
			args{
				v: StatusInactive,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Equal(tt.args.v); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatus_IsValid(t *testing.T) {
	tests := []struct {
		name string
		s    Status
		want bool
	}{
		{
			"unknown",
			StatusUnknown,
			false,
		},
		{
			"active",
			StatusActive,
			true,
		},
		{
			"inactive pending",
			StatusInactive,
			true,
		},
		{
			"inactive",
			StatusInactive,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
