package helm

import (
	"testing"
)

func Test_parseSchema(t *testing.T) {
	type args struct {
		repoUrl string
	}
	tests := []struct {
		name       string
		args       args
		wantSchema string
		wantErr    bool
	}{
		{
			name: "http schema",
			args: args{
				repoUrl: "http://example.com/path/to/repo",
			},
			wantSchema: "http",
			wantErr:    false,
		},
		{
			name: "https schema",
			args: args{
				repoUrl: "https://example.com/path/to/repo",
			},
			wantSchema: "https",
			wantErr:    false,
		},
		{
			name: "oci schema",
			args: args{
				repoUrl: "oci://example.com/path/to/repo",
			},
			wantSchema: "oci",
			wantErr:    false,
		},
		{
			name: "No schema",
			args: args{
				repoUrl: "example.com/path/to/repo",
			},
			wantSchema: "",
			wantErr:    true,
		},
		{
			name: "Wrong url",
			args: args{
				repoUrl: "example.com/path/to://repo",
			},
			wantSchema: "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSchema, err := parseSchema(tt.args.repoUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSchema() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSchema != tt.wantSchema {
				t.Errorf("parseSchema() gotSchema = %v, want %v", gotSchema, tt.wantSchema)
			}
		})
	}
}
