package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	checkBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("Got: %s, Want: %s", got, want)
		}
	}

	checkError := func(t *testing.T, value error, want error) {
		t.Helper()
		if value == nil {
			t.Fatal("Wanted an error, but didn't get one")
		}

		if value != want {
			t.Errorf("Got: %q, Expected: %q", value, want)
		}
	}

	checkInexistentError := func(t *testing.T, result error) {
		t.Helper()
		if result == nil {
			t.Fatal("Non expected error")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// fmt.Printf("The address of balance in the Test is %v \n", &wallet.balance)

		checkBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		// fmt.Printf("The address of balance in the Test is %v \n", &wallet.balance)

		checkBalance(t, wallet, Bitcoin(10))
		checkInexistentError(t, err)

	})

	t.Run("Withdraw with insufficient balance", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		checkBalance(t, wallet, initialBalance)
		checkError(t, err, ErrInsufficientBalance)

	})
}
