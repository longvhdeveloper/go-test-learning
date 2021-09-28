package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, expected Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != expected {
			t.Errorf("got %s want %s", got, expected)
		}
	}

	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err != want {
			t.Errorf("got %s, want %s", err.Error(), want)
		}
	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Error("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(50)
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(50))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(20))
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		assertBalance(t, wallet, Bitcoin(30))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
