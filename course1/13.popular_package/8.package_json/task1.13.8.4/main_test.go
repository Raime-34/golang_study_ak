package main

import "testing"

func Test_writeJSON(t *testing.T) {
	type args struct {
		filePath string
		data     interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Regular test",
			args: args{
				filePath: "./f/data",
				data: struct {
					Name string
					Age  int
				}{
					"Walter",
					54,
				},
			},
			wantErr: false,
		},
		{
			name: "Incorrect type test",
			args: args{
				filePath: "/f/data",
				data:     make(chan int),
			},
			wantErr: true,
		},
		{
			name: "Incorrect directory test",
			args: args{
				filePath: "\\/f?/data#",
				data:     3,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeJSON(tt.args.filePath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
