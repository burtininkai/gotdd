package main

import (
	"errors"
	"fmt"
)

//Bitcoin type
type Bitcoin int

//Stringer interface for bitcoin
type Stringer interface {
	String() string //printed when used with the %s format string in prints.
}

//String func for bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Wallet struct
type Wallet struct {
	balance Bitcoin //symbol (so variables, types, functions et al) lower case marks as private
}

//Deposit func to add sum
func (w *Wallet) Deposit(amount Bitcoin) { //*Wallet meand a pointer to Wallet, instead of copy
	w.balance += amount
}

//ErrInsufficitentFunds meessage=. var allows to define global values for the package
var ErrInsufficitentFunds = errors.New("cannot withdraw, insufficient funds")

//Withdraw method for wallet
func (w *Wallet) Withdraw(ammount Bitcoin) error {
	if ammount > w.balance {
		return ErrInsufficitentFunds
	}
	w.balance -= ammount
	return nil
}

//Balance method will list balance amount
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
