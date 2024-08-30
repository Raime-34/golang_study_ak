package main

import (
	"testing"
)

func Test_getStringHeader(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		{
			name: "Regular test",
			args: args{
				s: "Hi",
			},
			wantLength: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStringHeader(tt.args.s); got.Len != tt.wantLength {
				t.Errorf("getStringHeader() length %v, want %v", got.Len, tt.wantLength)
			}
		})
	}
}
