package main

import (
	"strings"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Random string test",
			args: args{
				length: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateRandomString(tt.args.length); len(got) != tt.want {
				t.Errorf("GenerateRandomString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateActivationKey(t *testing.T) {
	activationKey := generateActivationKey()

	if len(activationKey) != 19 {
		t.Errorf("generateActivationKey() length of key is incorrect")
	}

	if strings.Count(activationKey, "-") != 3 {
		t.Errorf("generateActivationKey() incorrect key format")
	}
}
