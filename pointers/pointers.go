package main

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// the * lets us point to the wallet we're using with these funcs
// struct pointers is the official name
// they're automatically dereferenced in Go

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
