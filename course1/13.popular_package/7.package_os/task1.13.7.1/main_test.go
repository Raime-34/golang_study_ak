package main

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	path := "../f/text.txt"

	type args struct {
		filePath string
		data     []byte
		perm     os.FileMode
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Not existing file test",
			args: args{
				filePath: path,
				data:     []byte("Message"),
				perm:     0644,
			},
			wantErr: false,
		},
		{
			name: "Existing file test",
			args: args{
				filePath: path,
				data:     []byte("Message"),
				perm:     0644,
			},
			wantErr: false,
		},
		{
			name: "Wrong path test",
			args: args{
				filePath: "fewfwe",
				data:     []byte("Message"),
				perm:     0644,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.filePath, tt.args.data, tt.args.perm); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
