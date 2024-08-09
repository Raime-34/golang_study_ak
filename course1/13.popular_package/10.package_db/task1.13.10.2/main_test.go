package main

import (
	"reflect"
	"testing"
)

func TestPrepareQuery(t *testing.T) {
	type args struct {
		operation string
		table     string
		user      User
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   []interface{}
		wantErr bool
	}{
		{
			name: "select generating test",
			args: args{
				operation: SELECT,
				table:     "users",
				user: User{
					ID: 69,
				},
			},
			want:    "SELECT * FROM users WHERE id = ?",
			want1:   []interface{}{69},
			wantErr: false,
		},
		{
			name: "insert generating test",
			args: args{
				operation: INSERT,
				table:     "users",
				user: User{
					ID:   69,
					Name: "Walter",
					Age:  43,
				},
			},
			want:    "INSERT INTO users (id,name,age) VALUES (?,?,?)",
			want1:   []interface{}{69, "Walter", 43},
			wantErr: false,
		},
		{
			name: "update generating test",
			args: args{
				operation: UPDATE,
				table:     "users",
				user: User{
					ID:   69,
					Name: "Walter",
					Age:  43,
				},
			},
			want:    "UPDATE users SET age = ?, name = ? WHERE id = ?",
			want1:   []interface{}{43, "Walter", 69},
			wantErr: false,
		},
		{
			name: "update generating test",
			args: args{
				operation: DELETE,
				table:     "users",
				user: User{
					ID:   69,
					Name: "Walter",
					Age:  43,
				},
			},
			want:    "DELETE FROM users WHERE id = ?",
			want1:   []interface{}{69},
			wantErr: false,
		},
		{
			name: "unhandled operation test",
			args: args{
				operation: "fewfewfwefwe",
				table:     "users",
				user: User{
					ID:   69,
					Name: "Walter",
					Age:  43,
				},
			},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := PrepareQuery(tt.args.operation, tt.args.table, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrepareQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrepareQuery() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("PrepareQuery() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
