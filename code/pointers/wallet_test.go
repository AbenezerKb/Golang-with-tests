package code

import "testing"

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(15)
	got := wallet.Balance()
	expected := 15.0

	if got != expected {

		t.Errorf("got %v expected %v", got, expected)

	}

}
