package crypto

import (
	"blockchain/utils"
	"errors"
	"fmt"
)

type Wallet struct {
	balance   int64
	KeyPair   string
	publicKey string
}

var INITIAL_BALANCE = 500
var (
	errAlreadyExists = errors.New("Transaction already exists")
)

//NewWallet generates a new wallet
func NewWallet() *Wallet {
	pKey := utils.GenerateEllepticKeyPair()
	return &Wallet{int64(INITIAL_BALANCE), utils.EncodeECDSAPrivateKey(pKey), utils.EncodeECDSAPublicKey(&pKey.PublicKey)}
}

//ToString prints a string representation
func (w *Wallet) ToString() string {
	return fmt.Sprintf("Balance : %d \n Public Key%x", w.balance, w.publicKey)
}

func (w *Wallet) GetPublicKey() string {
	return w.publicKey
}

func (w *Wallet) CreateTransaction(recp string, amount int64, transactionPool *TransactionPool) (*Transaction, error) {
	if amount > w.balance {
		return &Transaction{}, errAmountExceedsBalance
	}
	// fmt.Println("AMOUNT IS", amount)
	transaction, err := transactionPool.ExistingTransaction(w.publicKey)
	if err != nil {
		if err == errAlreadyExists {
			errNew := transaction.updateTransaction(w, recp, amount)
			if errNew != nil {
				return &Transaction{}, errNew
			}
			return transaction, err

		} else {
			panic(0)
		}

	} else {
		transaction, err = NewTransaction(w, recp, amount)
		if err != nil {
			return &Transaction{}, err
		}
		transactionPool.UpdateOrAddTransaction(transaction)
	}
	return transaction, nil

}
