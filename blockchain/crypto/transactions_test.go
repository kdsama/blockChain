package crypto

import (
	"blockchain/utils"
	"fmt"
	"testing"
)

func TestTransactions(t *testing.T) {
	t.Run("Try sending exceeding amount than the current balance", func(t *testing.T) {
		wallet := NewWallet()
		amount := 550
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := errAmountExceedsBalance
		_, got := NewTransaction(wallet, recipientAddress, int64(amount))
		if want != got {
			t.Errorf("This is shambles, wanted %s , but got %s", want, got)
		}
	})
	t.Run("For the new transaction Address confirm the subtracted amount", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := int64(470)
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		if err != nil {
			t.Errorf("Did not expect an error , but got %s", err)
		}
		got := obj.Outputs[0].Amount
		if want != got {
			t.Errorf("wanted subtracted amount to be %d but got %d", want, got)
		}
	})

	t.Run("For the new transaction , confirm the recipients address", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := recipientAddress
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		if err != nil {
			t.Errorf("Did not expect an error , but got %s", err)
		}
		got := obj.Outputs[1].Address
		if want != got {
			t.Errorf("wanted subtracted amount to be %s but got %s", want, got)
		}
	})

	t.Run("Wallet public Key should equal transactions outputs first address", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := wallet.publicKey
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		if err != nil {
			t.Errorf("Did not expect an error , but got %s", err)
		}
		got := obj.Outputs[0].Address
		if want != got {
			t.Errorf("wanted subtracted amount to be %s but got %s", want, got)
		}
	})

	t.Run("Create a transaction and confirm input amount", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := wallet.balance
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		if err != nil {
			t.Errorf("Did not expect an error , but got %s", err)
		}
		got := obj.Input.Balance
		if want != got {
			t.Errorf("wanted subtracted amount to be %d but got %d", want, got)
		}
	})

	t.Run("Validate a Valid Transaction", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := true
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		if err != nil {
			t.Error("Didnot expect an error for valid transaction")
		}
		got := obj.VerifyTransaction()
		fmt.Println(got)
		fmt.Println(want)
		if want != got {
			t.Errorf("Expected %t but got %t", want, got)
		}

	})

	t.Run("Invalidate  an InValid Transaction : Tampering with Amount", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := false
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		obj.Outputs[0].Amount = 1200
		if err != nil {
			t.Error("Didnot expect an error for valid transaction")
		}
		got := obj.VerifyTransaction()
		fmt.Println(got)
		fmt.Println(want)
		if want != got {
			t.Errorf("Expected %t but got %t", want, got)
		}

	})

	t.Run("Invalidate  an InValid Transaction : Empty Outputs", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := false
		obj, err := NewTransaction(wallet, recipientAddress, int64(amount))
		obj.Outputs = []Output{}
		if err != nil {
			t.Error("Didnot expect an error for valid transaction")
		}
		got := obj.VerifyTransaction()
		fmt.Println(got)
		fmt.Println(want)
		if want != got {
			t.Errorf("Expected %t but got %t", want, got)
		}

	})
}

func TestUpdateTransactions(t *testing.T) {
	t.Run("Updating a Transaction", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := int64(430)
		obj, _ := NewTransaction(wallet, recipientAddress, int64(amount))
		newAmount := int64(40)
		err := obj.updateTransaction(wallet, recipientAddress, newAmount)
		if err != nil {
			t.Error("Didnot expect an error here")
		}

		got := obj.Input.Balance
		if want != got {
			t.Errorf("Expected %d but got %d", want, got)
		}

	})

	t.Run("Updating a Transaction :: INVALID TRANSACTION", func(t *testing.T) {
		wallet := NewWallet()
		amount := 30
		recipientAddress := utils.EncodeECDSAPublicKey(&utils.GenerateEllepticKeyPair().PublicKey)
		want := errAmountExceedsBalance
		obj, _ := NewTransaction(wallet, recipientAddress, int64(amount))
		newAmount := int64(490)
		got := obj.updateTransaction(wallet, recipientAddress, newAmount)
		if want != got {
			t.Errorf("Expected %s but got %s", want, got)
		}

	})
}

// func gotAddress(outputs []output)(bool) {
// 	for i := range outputs {
// 		if outputs[i].address ==
// 	}
// }
