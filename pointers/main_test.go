package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("error is expected")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatalf("got err '%s' but didn't want to", got.Error())
		}
	}

	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.balance

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}
		err := wallet.Withdraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw with isufficient funds", func(t *testing.T) {
		currentBalance := Bitcoin(10)
		wallet := Wallet{
			balance: currentBalance,
		}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, currentBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
