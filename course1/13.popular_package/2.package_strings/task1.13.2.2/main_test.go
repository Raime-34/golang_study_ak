package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestFilterComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{
		{
			name: "Filter test",
			args: args{
				users: []User{
					{
						Name: "Betty",
						Comments: []Comment{
							{Message: "good Comment 1"},
							{Message: "BaD CoMmEnT 2"},
							{Message: "Bad Comment 3"},
							{Message: "Use camelCase please"},
						},
					},
					{
						Name: "Jhon",
						Comments: []Comment{
							{Message: "Good Comment 1"},
							{Message: "Good Comment 2"},
							{Message: "Good Comment 3"},
							{Message: "Bad Comment 4"},
						},
					},
				},
			},
			want: []User{
				{
					Name: "Betty",
					Comments: []Comment{
						{Message: "good Comment 1"},
						{Message: "Use camelCase please"},
					},
				},
				{
					Name: "Jhon",
					Comments: []Comment{
						{Message: "Good Comment 1"},
						{Message: "Good Comment 2"},
						{Message: "Good Comment 3"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FilterComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBadComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{
			name: "Bad comments test",
			args: args{
				users: []User{
					{
						Name: "Betty",
						Comments: []Comment{
							{Message: "good Comment 1"},
							{Message: "BaD CoMmEnT 2"},
							{Message: "Bad Comment 3"},
							{Message: "Use camelCase please"},
						},
					},
					{
						Name: "Jhon",
						Comments: []Comment{
							{Message: "Good Comment 1"},
							{Message: "Good Comment 2"},
							{Message: "Good Comment 3"},
							{Message: "Bad Comment 4"},
						},
					},
				},
			},
			want: []Comment{
				{Message: "BaD CoMmEnT 2"},
				{Message: "Bad Comment 3"},
				{Message: "Bad Comment 4"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBadComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBadComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetGoodComments(t *testing.T) {
	type args struct {
		users []User
	}
	tests := []struct {
		name string
		args args
		want []Comment
	}{
		{
			name: "Bad comments test",
			args: args{
				users: []User{
					{
						Name: "Betty",
						Comments: []Comment{
							{Message: "good Comment 1"},
							{Message: "BaD CoMmEnT 2"},
							{Message: "Bad Comment 3"},
							{Message: "Use camelCase please"},
						},
					},
					{
						Name: "Jhon",
						Comments: []Comment{
							{Message: "Good Comment 1"},
							{Message: "Good Comment 2"},
							{Message: "Good Comment 3"},
							{Message: "Bad Comment 4"},
						},
					},
				},
			},
			want: []Comment{
				{Message: "good Comment 1"},
				{Message: "Use camelCase please"},
				{Message: "Good Comment 1"},
				{Message: "Good Comment 2"},
				{Message: "Good Comment 3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGoodComments(tt.args.users); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetGoodComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBadComment(t *testing.T) {
	type args struct {
		comment string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "isBadComment test",
			args: args{
				"Bad",
			},
			want: true,
		},
		{
			name: "isBadComment test",
			args: args{
				"B4d",
			},
			want: false,
		},
		{
			name: "isBadComment test",
			args: args{
				"Badlands",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBadComment(tt.args.comment); got != tt.want {
				t.Errorf("IsBadComment() = %v, want %v", got, tt.want)
			}
		})
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

	if stdout.String() != "[{Betty [{good Comment 1} {Use camelCase please}]} {Jhon [{Good Comment 1} {Good Comment 2} {Good Comment 3}]}]\n" {
		t.Errorf("main function test fail")
	}
}
