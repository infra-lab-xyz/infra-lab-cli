package utils

import (
	"os"
	"os/user"
	"testing"
)

func Test_convertToMiB(t *testing.T) {
	type args struct {
		size string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "plain MiB value",
			args:    args{size: "2048"},
			want:    2048,
			wantErr: false,
		},
		{
			name:    "MiB with M suffix",
			args:    args{size: "2048M"},
			want:    2048,
			wantErr: false,
		},
		{
			name:    "MiB with m suffix",
			args:    args{size: "2048m"},
			want:    2048,
			wantErr: false,
		},
		{
			name:    "GiB with G suffix",
			args:    args{size: "2G"},
			want:    2048,
			wantErr: false,
		},
		{
			name:    "GiB with g suffix",
			args:    args{size: "2g"},
			want:    2048,
			wantErr: false,
		},
		{
			name:    "Fractional GiB",
			args:    args{size: "2.5G"},
			want:    2560,
			wantErr: false,
		},
		{
			name:    "Invalid input",
			args:    args{size: "abc"},
			want:    0,
			wantErr: true,
		},
		{
			name:    "Invalid input GB",
			args:    args{size: "1GB"},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertToMiB(tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("convertToMiB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("convertToMiB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertMiBToGiB(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "2GiB",
			args:    args{size: 2048},
			want:    2.0,
			wantErr: false,
		},
		{
			name:    "0.5GiB",
			args:    args{size: 512},
			want:    0.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertMiBToGiB(tt.args.size)
			if got != tt.want {
				t.Errorf("convertMiBToGiB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestIsDirExist(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "file")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			"existing directory",
			tmpDir,
			true,
		},
		{
			"existing file",
			tmpFile.Name(),
			false,
		},
		{
			"non-existent path",
			tmpDir + "/doesnotexist",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDirExist(tt.path); got != tt.want {
				t.Errorf("IsDirExist(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestExpandPath(t *testing.T) {
	usr, _ := user.Current()
	home := usr.HomeDir
	_ = os.Setenv("FOO", "bar")

	tests := []struct {
		name string
		path string
		want string
	}{
		{
			"home directory",
			"~/test",
			home + "/test",
		},
		{
			"expand env variable",
			"/tmp/$FOO",
			"/tmp/bar",
		},
		{
			"expand env variable with curvatures",
			"/tmp/${FOO}",
			"/tmp/bar",
		},
		{
			"home dir with env variable",
			"~/test/$FOO",
			home + "/test/bar",
		},
		{
			"nothing to expand",
			"/no/vars",
			"/no/vars",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpandPath(tt.path); got != tt.want {
				t.Errorf("IsDirExist(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}
