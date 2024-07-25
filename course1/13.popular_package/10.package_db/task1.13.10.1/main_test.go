package main

import (
	"reflect"
	"testing"
)

func TestCreateUserTable(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Connect test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserTable(); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInsertUser(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Insert test",
			args: args{
				user: User{
					ID:   69,
					Name: "Walter",
					Age:  54,
				},
			},
			wantErr: false,
		},
		{
			name: "Insert same value",
			args: args{
				user: User{
					ID:   69,
					Name: "Walter",
					Age:  54,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSelectUser(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{
			name: "Existing user select test",
			args: args{
				id: 69,
			},
			want: User{
				ID:   69,
				Name: "Walter",
				Age:  54,
			},
			wantErr: false,
		},
		{
			name: "Not existing user select test",
			args: args{
				id: 69453252,
			},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	type args struct {
		user User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Existing user update test",
			args: args{
				user: User{
					ID:   69,
					Name: "John",
					Age:  5000,
				},
			},
			wantErr: false,
		},
		{
			name: "Existing user update test",
			args: args{
				user: User{
					ID:   69534534,
					Name: "John",
					Age:  5000,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Existing user delete test",
			args: args{
				id: 69,
			},
			wantErr: false,
		},
		{
			name: "Not existing user delete test",
			args: args{
				id: 6921341,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteUser(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
