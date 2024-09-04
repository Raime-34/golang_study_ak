package main

import "testing"

func Test_isValidEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Correct email test",
			args: args{
				email: "user@example.com",
			},
			want: true,
		},
		{
			name: "Incorrect email test",
			args: args{
				email: "user",
			},
			want: false,
		},
		{
			name: "email without username test",
			args: args{
				email: "@example.com",
			},
			want: false,
		},
		{
			name: "email in the middle of str test",
			args: args{
				email: "us@example.comer",
			},
			want: false,
		},
		{
			name: "email in the start of str test",
			args: args{
				email: "@example.comer",
			},
			want: false,
		},
		{
			name: "incorrect suffix test",
			args: args{
				email: "user@example",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidEmail(tt.args.email); got != tt.want {
				t.Errorf("isValidEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
