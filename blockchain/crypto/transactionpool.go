package crypto

import (
	"errors"
)

var (
	errMalformedTransaction        = errors.New("Found Malformed Transaction")
	errInvalidTransactionSignature = errors.New("Invalid Transaction Signature")
)

type TransactionPool struct {
	Transactions []Transaction `json:"transactions"`
}

func NewTransactionPool(transactions []Transaction) *TransactionPool {
	return &TransactionPool{transactions}
}

func (tp *TransactionPool) UpdateOrAddTransaction(transaction *Transaction) {
	index := -1
	for i := range tp.Transactions {
		if tp.Transactions[i].Id == transaction.Id {
			// Mean
			index = i
			break
		}
	}
	if index != -1 {
		// Update existing transaction
		tp.Transactions[index] = *transaction
	} else {
		tp.Transactions = append(tp.Transactions, *transaction)
	}
}

func (tp *TransactionPool) ExistingTransaction(publicKey string) (*Transaction, error) {
	for i := range tp.Transactions {
		if tp.Transactions[i].Input.PublicKey == publicKey {
			return &tp.Transactions[i], errAlreadyExists
		}
	}
	return &Transaction{}, nil
}

func (tp *TransactionPool) GetTransactions() []Transaction {
	return tp.Transactions
}

func (tp *TransactionPool) ValidTransactions() []Transaction {
	// Total output amount matches balance in input
	// Verify signature of every transaction to confirm data is not corrupted
	validTr := []Transaction{}
	for index := range tp.Transactions {
		inputTotal := tp.Transactions[index].Input.Balance
		OutputTotal := int64(0)
		for indexOutput := range tp.Transactions[index].Outputs {
			OutputTotal += tp.Transactions[index].Outputs[indexOutput].Amount
		}

		if inputTotal != OutputTotal {
			continue
		}
		if !tp.Transactions[index].VerifyTransaction() {
			continue
		}

		validTr = append(validTr, tp.Transactions[index])

	}
	return validTr
}
func (tp *TransactionPool) Clear() {
	tp.Transactions = []Transaction{}
}
