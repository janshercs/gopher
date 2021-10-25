package pointer

import (
	"errors"
	"fmt"
)

var errInsufficientFunds = errors.New("Oh geez!")
var errNegativeDeposit = errors.New("ERROR! NEGATIVE DEPOSIT!")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount < 0 {
		return errNegativeDeposit
	}

	w.balance += amount
	return nil
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // pointers here don't have to be dereferenced!
}
