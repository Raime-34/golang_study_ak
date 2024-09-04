package main

import "testing"

func Test_concatStrings(t *testing.T) {
	type args struct {
		xs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Regular check",
			args: args{
				xs: []string{"Hello", " ", "world!"},
			},
			want: "Hello world!",
		},

		{
			name: "Single word check",
			args: args{
				xs: []string{"Hi!"},
			},
			want: "Hi!",
		},

		{
			name: "Empty slice check",
			args: args{
				xs: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatStrings(tt.args.xs...); got != tt.want {
				t.Errorf("concatStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
