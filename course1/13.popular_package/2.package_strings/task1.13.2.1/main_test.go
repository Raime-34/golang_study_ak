package main

import (
	"reflect"
	"testing"
)

func TestCountWordsInText(t *testing.T) {
	type args struct {
		txt   string
		words []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			name: "Regular test",
			args: args{
				txt: "Lorem ipsum dolor sit amet, " +
					"consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. " +
					"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. " +
					"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. " +
					"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				words: []string{"amet", "in", "ut"},
			},
			want: map[string]int{
				"amet": 1,
				"in":   3,
				"ut":   3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountWordsInText(tt.args.txt, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountWordsInText() = %v, want %v", got, tt.want)
			}
		})
	}
}
