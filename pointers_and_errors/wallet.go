package main

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance Bitcoin
}

// NOTE: Go lets you create new types from existing ones. The syntax is: type MyName OriginalType

type Bitcoin int

type Stringer interface {
	String() string
}

// NOTE: These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in deposit is %p \n", &w.balance)
	w.balance += amount
}

/*
func (w *Wallet) Balance() int {
	return (*w).balance
}
*/

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insuficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
