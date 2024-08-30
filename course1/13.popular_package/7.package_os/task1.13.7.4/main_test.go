package main

import "testing"

func TestExecBin(t *testing.T) {
	type args struct {
		binPath string
		args    []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Execute with args",
			args: args{
				binPath: "./someBinary.exe",
				args:    []string{"Walter"},
			},
			want: "Hi, Walter!",
		},
		{
			name: "Execute without args",
			args: args{
				binPath: "./someBinary.exe",
			},
			want: "Hi!",
		},
		{
			name: "Not existing binary",
			args: args{
				binPath: "./someBi",
				args:    []string{"Walter"},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecBin(tt.args.binPath, tt.args.args...); got != tt.want {
				t.Errorf("ExecBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
