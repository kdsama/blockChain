package crypto

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

func (tp *TransactionPool) GetTransactions() *[]Transaction {
	return &tp.Transactions
}
