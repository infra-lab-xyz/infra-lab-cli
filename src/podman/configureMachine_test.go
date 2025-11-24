package podman

import (
	"testing"
)

func Test_checkIfParamChanged(t *testing.T) {
	type args struct {
		param        ConfigParam
		currentValue int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Param not changed",
			args: args{
				param:        ConfigParam{ValueFlag: 2},
				currentValue: 2,
			},
			want: false,
		},
		{
			name: "Param changed",
			args: args{
				param:        ConfigParam{ValueFlag: 4},
				currentValue: 2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			param := tt.args.param
			_ = checkIfParamChanged(&param, tt.args.currentValue)
			if param.IsChanged != tt.want {
				t.Errorf("IsChanged = %v, want %v", param.IsChanged, tt.want)
			}
		})
	}
}
