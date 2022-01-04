package crypto

import (
	"blockchain/blockchain"
	"encoding/json"
	"fmt"
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

// func TestWalletTransactionPool(t *testing.T) {
// 	w := NewWallet()
// 	tp := NewTransactionPool([]Transaction{})

// 	t.Run("Creating a transaction", func(t *testing.T) {
// 		amount := int64(50)
// 		recipient := "r4nd-add4rss-som21234"
// 		_, err := w.CreateTransaction(recipient, amount, tp)
// 		if err != errAmountExceedsBalance && err != errAlreadyExists && err != nil {
// 			t.Error("Did not expect an error here ")
// 		}
// 		want := amount
// 		got := tp.Transactions[0].Input.Balance

// 		if want != got {
// 			t.Errorf("Wanted %d , but got %d", want, got)
// 		}

// 	})
// }

func TestCalculateBalance(t *testing.T) {

	t.Run("Verify Receiver's wallet", func(t *testing.T) {
		rw := NewWallet()

		sw := NewWallet()
		bl := blockchain.NewBlockChain()
		// fmt.Println(sw.publicKey, rw.publicKey)
		tp := NewTransactionPool([]Transaction{})
		addBalance := int64(20)
		repeat := 3

		// fmt.Println("New blockchain length", len(bl.Chain))
		_, err := sw.CreateTransaction(sw.publicKey, int64(0), tp, bl)
		for i := 0; i < repeat; i++ {

			_, err := sw.CreateTransaction(rw.publicKey, addBalance, tp, bl)
			fmt.Println("This is now new", tp)
			// fmt.Println(len(tp.Transactions[len(tp.Transactions)-1].Outputs))
			if err != nil && err != errAlreadyExists {
				// fmt.Println("???")
				// fmt.Println(err)
				t.Error("Did not expect an error here", err)
			}
		}
		body, err := json.Marshal(tp.Transactions)
		if err != nil {

			// fmt.Println("Marshalling error", err)

		}
		bl.AddBlock(string(body))
		want := int64(INITIAL_BALANCE) + int64(addBalance*int64(repeat))

		got := rw.CalculateBalance(bl)
		fmt.Println(got, want)
		// fmt.Println(got)
		if want != got {
			t.Errorf("Expected %d but got %d", want, got)
		}
	})
	// t.Run("Verify Sender's wallet", func(t *testing.T) {
	// 	rw := NewWallet()

	// 	sw := NewWallet()
	// 	bl := blockchain.NewBlockChain()
	// 	// fmt.Println(sw.publicKey, rw.publicKey)
	// 	tp := NewTransactionPool([]Transaction{})
	// 	addBalance := int64(20)
	// 	repeat := 3

	// 	// fmt.Println("New blockchain length", len(bl.Chain))
	// 	_, err := sw.CreateTransaction(sw.publicKey, int64(0), tp, bl)
	// 	for i := 0; i < repeat; i++ {

	// 		_, err := sw.CreateTransaction(rw.publicKey, addBalance, tp, bl)
	// 		// fmt.Println("This is now new", tp)
	// 		// fmt.Println(len(tp.Transactions[len(tp.Transactions)-1].Outputs))
	// 		if err != nil && err != errAlreadyExists {
	// 			// fmt.Println("???")
	// 			// fmt.Println(err)
	// 			t.Error("Did not expect an error here", err)
	// 		}
	// 	}
	// 	fmt.Println(tp.Transactions[1])
	// 	body, err := json.Marshal(tp.Transactions)
	// 	if err != nil {

	// 		// fmt.Println("Marshalling error", err)

	// 	}
	// 	bl.AddBlock(string(body))
	// 	want := int64(INITIAL_BALANCE) - int64(addBalance*int64(repeat))
	// 	got := sw.CalculateBalance(bl)

	// 	if want != got {
	// 		t.Errorf("Expected %d but got %d", want, got)
	// 	}
	// })

}
