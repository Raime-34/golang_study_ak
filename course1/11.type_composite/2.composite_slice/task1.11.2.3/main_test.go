package main

import "testing"

func Test_bitwiseXOR(t *testing.T) {
	type args struct {
		n   int
		res int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Equal numbers",
			args: args{
				n:   10,
				res: 10,
			},
			want: 0,
		},
		{
			name: "Different numbers",
			args: args{
				n:   10,
				res: 5,
			},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseXOR(tt.args.n, tt.args.res); got != tt.want {
				t.Errorf("bitwiseXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findSingleNumber(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Base test",
			args: args{
				numbers: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSingleNumber(tt.args.numbers); got != tt.want {
				t.Errorf("findSingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
