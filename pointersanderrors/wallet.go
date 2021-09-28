package pointersanderrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}

func (wallet *Wallet) Balance() Bitcoin {
	//fmt.Printf("address of balance in Deposit is %v \n", &wallet.balance)
	return wallet.balance
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsufficientFunds
	}
	wallet.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
