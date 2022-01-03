package crypto

import (
	"reflect"
	"testing"
)

func TestTransactionPool(t *testing.T) {
	tran := []Transaction{}
	tp := NewTransactionPool(tran)
	w := NewWallet()
	transaction, err := NewTransaction(w, "r4nd-4dress", 30)
	if err != nil {
		t.Error("Not expecting issue in creating a transaction")
	}
	tp.UpdateOrAddTransaction(transaction)
	t.Run("Check if A New transaction is added or not", func(t *testing.T) {
		// Currently there is only one transaction
		got := tp.Transactions[0].Id
		want := transaction.Id
		if got != want {
			t.Errorf("Expected %s but got %s", got, want)

		}
	})

	t.Run("Check if transaction is added", func(t *testing.T) {
		// Currently there is only one transaction
		oldTransaction := tp.Transactions[0]
		NewTransaction := oldTransaction
		err := NewTransaction.updateTransaction(w, "new-addr", 40)
		if err != nil {
			t.Error("Did not expect error in creating Transaction")

		}
		tp.UpdateOrAddTransaction(&NewTransaction)
		ok := reflect.DeepEqual(oldTransaction, tp.Transactions[0])
		if ok {
			t.Error("Did not expect the transactions to be equal")
		}
	})
}
