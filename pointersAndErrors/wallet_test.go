package wallet

import (
	"testing"
)


func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	}) 
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		_, err := wallet.Withdraw(BitCoin(10))
		
		assertNoError(t, err)
		assertBalance(t, wallet, BitCoin(10))
	})
	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(10)}
		amountWithdrawn, err := wallet.Withdraw(BitCoin(20))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, wallet.Balance() - amountWithdrawn)
	})
}

func assertBalance (t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got ,want)
	}
}

func assertNoError(t testing.TB, got error) {
	if got != nil {
		t.Errorf("expected nil, got error")
	}
}

func assertError (t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected Error, got nil")
	}
	if got != want {
		t.Errorf("got %v, want %s", got, want)
	}
}