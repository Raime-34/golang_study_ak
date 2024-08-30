package main

import "testing"

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Zero check",
			args{
				n: 0,
			},
			1,
		},

		{
			"Check 1",
			args{
				n: 0,
			},
			1,
		},

		{
			"Negative check",
			args{
				n: -200,
			},
			1,
		},

		{
			"Normal check",
			args{
				n: 5,
			},
			120,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Factorial(tt.args.n); got != tt.want {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
