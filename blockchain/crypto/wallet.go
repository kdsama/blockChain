package crypto

import (
	"blockchain/blockchain"
	"blockchain/utils"
	"encoding/json"
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
func NewBlockChainWallet() *Wallet {
	pKey := utils.GenerateEllepticKeyPair()
	return &Wallet{int64(-1), utils.EncodeECDSAPrivateKey(pKey), utils.EncodeECDSAPublicKey(&pKey.PublicKey)}
}

//ToString prints a string representation
func (w *Wallet) ToString() string {
	return fmt.Sprintf("Balance : %d \n Public Key%x", w.balance, w.publicKey)
}

func (w *Wallet) GetPublicKey() string {
	return w.publicKey
}

func (w *Wallet) CreateTransaction(recp string, amount int64, transactionPool *TransactionPool, bl *blockchain.Blockchain) (*Transaction, error) {
	w.balance = w.CalculateBalance(bl)
	// fmt.Println("New balance", w.balance)
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

func (w *Wallet) CalculateBalance(bl *blockchain.Blockchain) int64 {
	// balance := w.balance
	balanceNow := w.balance
	// fmt.Println("BALANCE NOW IS", w.balance)
	var wantBlock []Transaction
	var latestBlock Transaction
	// fmt.Println("New Block of the chain", bl.Chain[len(bl.Chain)-1].Data)
	if len(bl.Chain) > 1 {

		for i := 1; i < len(bl.Chain); i++ {
			var tr []Transaction

			if err := json.Unmarshal([]byte(bl.Chain[i].Data), &tr); err != nil {
				// fmt.Println("Continue,FirstBlock???")
				// Will continue in case of first block
				continue
			}

			// fmt.Println("New Block")
			prev := int64(0)
			for j := range tr {
				wantBlock = append(wantBlock, tr[j])
				if tr[j].Input.PublicKey == w.publicKey {

					if tr[j].Input.Timestamp > prev {
						prev = tr[j].Input.Timestamp
						latestBlock = tr[j]
					}
				}
			}
			// fmt.Println("<<<<<<<<<<<<<<>>>>>>>>>>>>>>>>>>>>>>>>>>", latestBlock.Outputs)
			// x, v :=
		}
	}

	if len(wantBlock) == 0 {
		// fmt.Println("No block exists where this wallet's address is present")
		return balanceNow
	}
	// fmt.Println(len(latestBlock.Outputs))
	for j := range wantBlock {

		if wantBlock[j].Input.Timestamp > latestBlock.Input.Timestamp {
			if wantBlock[j].Input.PublicKey != w.publicKey {

				for k := range wantBlock[j].Outputs {
					if wantBlock[j].Outputs[k].Address == w.publicKey {

						balanceNow += wantBlock[j].Outputs[k].Amount
					}

				}
			} else {
				for k := 1; k < len(wantBlock[j].Outputs); k++ {
					if wantBlock[j].Outputs[k].Address == w.publicKey {

						balanceNow -= wantBlock[j].Outputs[k].Amount
					}

				}
			}

		}
	}
	return balanceNow
}
