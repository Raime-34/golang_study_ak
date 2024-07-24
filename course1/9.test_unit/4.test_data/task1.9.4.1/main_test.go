package main

import (
	"bytes"
	"math"
	"os"
	"reflect"
	"testing"
)

func Test_average(t *testing.T) {
	type args struct {
		xs []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		isNan bool
	}{
		{
			name: "Zero check",
			args: args{
				xs: generateSlice(0),
			},
			want:  math.NaN(),
			isNan: true,
		},

		{
			name: "Regular check",
			args: args{
				xs: generateSlice(10),
			},
			want:  5.5,
			isNan: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := average(tt.args.xs); (got != tt.want && tt.isNan == false) || math.IsNaN(got) != tt.isNan {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateSlice(t *testing.T) {
	first := generateSlice(10)
	second := generateSlice(10)

	if !reflect.DeepEqual(first, second) {
		t.Errorf("second generated slice is %v, expected %v", second, first)
	}
}

func TestMainFunc(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	w.Close()
	os.Stdout = old

	var stdout bytes.Buffer
	stdout.ReadFrom(r)
	expected := "5.5\n"

	if stdout.String() != expected {
		t.Errorf("expected %v, got %v", expected, stdout.String())
	}
}
