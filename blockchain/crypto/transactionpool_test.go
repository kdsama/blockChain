package crypto

import (
	"blockchain/blockchain"
	"fmt"
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

func TestValidTransactions(t *testing.T) {
	tran := []Transaction{}
	tp := NewTransactionPool(tran)

	for i := 0; i <= 6; i++ {
		wa := NewWallet()
		bl := blockchain.NewBlockChain()
		fmt.Println("I >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>", i)
		transaction, err := wa.CreateTransaction("rand-4dr-355", 30, tp, bl)
		if err != nil {
			t.Error("Didnot expect an error here", err)
		}
		fmt.Println(tp.Transactions)

		if err != errAmountExceedsBalance && err != errAlreadyExists && err != nil {
			t.Error("Did not expect an error here ")
		}

		if i%2 == 0 || i == 0 {
			fmt.Println(len(tp.Transactions) - 1)
			tp.Transactions[len(tp.Transactions)-1].Input.Balance = int64(99999)
		} else {

			tran = append(tran, *transaction)
		}
	}

	t.Run("Mixing Valid and Corrupt Transaction", func(t *testing.T) {

		newTransactions := tp.ValidTransactions()
		ok := reflect.DeepEqual(newTransactions, tran)

		if !ok {
			t.Error("Expected Equal Same Transactions but did not get them ")
		}
	})

	t.Run("clear Transactions", func(t *testing.T) {
		want := 0
		tp.Clear()
		got := len(tp.Transactions)
		if want != got {
			t.Errorf("Expected final length of transaction list be %d, but got %d", want, got)
		}
	})
}
