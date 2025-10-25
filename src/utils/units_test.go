package utils

import (
	"testing"
)

func TestByteCountIEC(t *testing.T) {
	type args struct {
		b int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Bytes less than 1 KiB",
			args: args{b: 512},
			want: "512 B",
		},
		{
			name: "Exactly 1 KiB",
			args: args{b: 1024},
			want: "1.0 KiB",
		},
		{
			name: "Exactly 1.5 KiB",
			args: args{b: 1536},
			want: "1.5 KiB",
		},
		{
			name: "Multiple KiB",
			args: args{b: 2048},
			want: "2.0 KiB",
		},
		{
			name: "Exactly 1 MiB",
			args: args{b: 1024 * 1024},
			want: "1.0 MiB",
		},
		{
			name: "Exactly 1 GiB",
			args: args{b: 1024 * 1024 * 1024},
			want: "1.0 GiB",
		},
		{
			name: "Exactly 1 TiB",
			args: args{b: 1024 * 1024 * 1024 * 1024},
			want: "1.0 TiB",
		},
		{
			name: "Exactly 1 PiB",
			args: args{b: 1024 * 1024 * 1024 * 1024 * 1024},
			want: "1.0 PiB",
		},
		{
			name: "Exactly 1 EiB",
			args: args{b: 1024 * 1024 * 1024 * 1024 * 1024 * 1024},
			want: "1.0 EiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteCountIEC(tt.args.b); got != tt.want {
				t.Errorf("ByteCountIEC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConvertToDesiredUnit(t *testing.T) {
	type args struct {
		value       string
		desiredUnit string
	}
	tests := []struct {
		name    string
		args    args
		want    Unit
		wantErr bool
	}{
		{
			name:    "1G to 1024M",
			args:    args{value: "1G", desiredUnit: "M"},
			want:    Unit{FloatStr: "1024.0M", FloatValue: 1024.0, IntStr: "1024M", IntValue: 1024},
			wantErr: false,
		},
		{
			name:    "1G to 1G",
			args:    args{value: "1G", desiredUnit: "G"},
			want:    Unit{FloatStr: "1.0G", FloatValue: 1.0, IntStr: "1G", IntValue: 1},
			wantErr: false,
		},
		{
			name:    "1024M to 1G",
			args:    args{value: "1024M", desiredUnit: "G"},
			want:    Unit{FloatStr: "1.0G", FloatValue: 1.0, IntStr: "1G", IntValue: 1},
			wantErr: false,
		},
		{
			name:    "1048576k to 1G",
			args:    args{value: "1048576k", desiredUnit: "G"},
			want:    Unit{FloatStr: "1.0G", FloatValue: 1.0, IntStr: "1G", IntValue: 1},
			wantErr: false,
		},
		{
			name:    "2.5G to 2560M",
			args:    args{value: "2.5G", desiredUnit: "M"},
			want:    Unit{FloatStr: "2560.0M", FloatValue: 2560.0, IntStr: "2560M", IntValue: 2560},
			wantErr: false,
		},
		{
			name:    "2560M to 2.5G",
			args:    args{value: "2560M", desiredUnit: "G"},
			want:    Unit{FloatStr: "2.5G", FloatValue: 2.5, IntStr: "2G", IntValue: 2},
			wantErr: false,
		},
		{
			name:    "3000M to 2.9G",
			args:    args{value: "3000M", desiredUnit: "G"},
			want:    Unit{FloatStr: "2.9G", FloatValue: 2.9, IntStr: "2G", IntValue: 2},
			wantErr: false,
		},
		{
			name:    "10.5G to 10.5G",
			args:    args{value: "10.5G", desiredUnit: "G"},
			want:    Unit{FloatStr: "10.5G", FloatValue: 10.5, IntStr: "10G", IntValue: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToDesiredUnit(tt.args.value, tt.args.desiredUnit)
			if got != tt.want {
				t.Errorf("ConvertToDesiredUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
