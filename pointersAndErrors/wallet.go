package wallet

import (
	"errors"
	"fmt"
)

type BitCoin int

var ErrInsufficientFunds = errors.New("withdraw canceled, insufficient funds")

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) (BitCoin, error) {
	if amount > w.balance {
		return BitCoin(0), ErrInsufficientFunds
	}
	w.balance -= amount
	return amount, nil
}