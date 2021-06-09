package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(value Bitcoin) {
	// fmt.Printf("The address of balance in the Test is %v \n", &w.balance)
	w.balance += value
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientBalance = errors.New("it is not possible to withdraw: insufficient balance")

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return ErrInsufficientBalance
	}
	w.balance -= value
	return nil
}
