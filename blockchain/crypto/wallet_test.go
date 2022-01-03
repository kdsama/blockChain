package crypto

import (
	"testing"
)

func TestWallet(t *testing.T) {
	w := NewWallet()
	want := int64(INITIAL_BALANCE)
	got := w.balance
	if want != got {
		t.Errorf("wanted %d but got %d", want, got)
	}
}

func TestWalletTransactionPool(t *testing.T) {
	w := NewWallet()
	tp := NewTransactionPool([]Transaction{})

	t.Run("Creating a transaction", func(t *testing.T) {
		amount := int64(50)
		recipient := "r4nd-add4rss-som21234"
		_, err := w.CreateTransaction(recipient, amount, tp)
		if err != errAmountExceedsBalance && err != errAlreadyExists && err != nil {
			t.Error("Did not expect an error here ")
		}
		want := amount
		got := tp.Transactions[0].Input.Balance

		if want != got {
			t.Errorf("Wanted %d , but got %d", want, got)
		}

	})
}
