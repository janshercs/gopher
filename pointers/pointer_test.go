package pointer

import "testing"

func TestWallet(t *testing.T) {
	assertWalletBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, err error, want error) {
		t.Helper()
		if err == nil {
			t.Fatal("wanted error but did not get one.")
		}
		if err != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("Didn't want no error but got one.")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertNoError(t, err)
		assertWalletBalance(t, wallet, want)
	})

	t.Run("Deposit negative amount", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Deposit(Bitcoin(-10))
		want := Bitcoin(10)

		assertWalletBalance(t, wallet, want)
		assertError(t, err, errNegativeDeposit)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertNoError(t, err)
		assertWalletBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertWalletBalance(t, wallet, startingBalance)
		assertError(t, err, errInsufficientFunds)
	})
}
