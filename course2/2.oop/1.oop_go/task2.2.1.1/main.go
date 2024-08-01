package main

import "fmt"

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type BaseAccount struct {
	balance float64
}

func (b *BaseAccount) Balance() float64 {
	return b.balance
}

func (b *BaseAccount) Deposit(amount float64) error {
	b.balance += amount
	return nil
}

type CurrentAccount struct {
	BaseAccount
}

func (c *CurrentAccount) Withdraw(amount float64) error {
	if c.balance < amount {
		return fmt.Errorf("недостаточно средств")
	}

	c.balance -= amount
	return nil
}

type SavingAccount struct {
	BaseAccount
}

func (s *SavingAccount) Withdraw(amount float64) error {
	if s.balance < 500 {
		return fmt.Errorf("недостаточно средств")
	}

	s.balance -= amount
	return nil
}

func ProcessAccount(account Account) {
	account.Deposit(500)
	account.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", account.Balance())
}
