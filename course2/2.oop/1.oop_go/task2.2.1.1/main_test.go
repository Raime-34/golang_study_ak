package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TProcessAccount(t *testing.T, account Account, expected string) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	ProcessAccount(account)
	w.Close()

	var stdout bytes.Buffer
	stdout.ReadFrom(r)
	os.Stdout = old

	if stdout.String() != expected {
		t.Errorf("%s error", reflect.TypeOf(account))
	}
}

func TestCurrentAccount(t *testing.T) {
	account := CurrentAccount{}
	TProcessAccount(t, &account, "Balance: 300.00\n")
}

func TestCurrentAccount2(t *testing.T) {
	account := CurrentAccount{}
	if err := account.Withdraw(1000); err == nil {
		t.Errorf("Current account 0 money withdraw test")
	}
}

func TestSavingAccount(t *testing.T) {
	account := SavingAccount{}
	TProcessAccount(t, &account, "Balance: 300.00\n")
}

func TestSavingAccount2(t *testing.T) {
	account := SavingAccount{}
	if err := account.Withdraw(1000); err == nil {
		t.Errorf("Saving account 0 money withdraw test")
	}
}
