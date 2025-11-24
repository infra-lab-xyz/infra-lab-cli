package utils

import (
	"os"
	"os/user"
	"testing"
)

func TestIsDirExist(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile, err := os.CreateTemp(tmpDir, "file")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer func(tmpFile *os.File) {
		_ = tmpFile.Close()
	}(tmpFile)

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

func TestMapToString(t *testing.T) {
	type args struct {
		data      map[string]any
		separator string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "Single key val pair with space",
			args: args{
				data: map[string]any{
					"foo": "bar",
				},
				separator: " ",
			},
			wantResult: "foo bar",
		},
		{
			name: "Single key (str) val (int) pair with space",
			args: args{
				data: map[string]any{
					"--cpu": 2,
				},
				separator: " ",
			},
			wantResult: "--cpu 2",
		},
		{
			name: "Single key val pair with equal sign",
			args: args{
				data: map[string]any{
					"foo": "bar",
				},
				separator: "=",
			},
			wantResult: "foo=bar",
		},
		{
			name: "Two key val pair with equal sign",
			args: args{
				data: map[string]any{
					"foo": "bar",
					"key": "val",
				},
				separator: "=",
			},
			wantResult: "foo=bar key=val",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := MapToString(tt.args.data, tt.args.separator); gotResult != tt.wantResult {
				t.Errorf("MapToString() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
