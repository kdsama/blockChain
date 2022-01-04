package crypto

import (
	"blockchain/utils"
	"bytes"
	"encoding/gob"
	"errors"
	"log"
	"math/big"
)

type Output struct {
	Amount  int64  `json:"amount"`
	Address string `json:"address"`
}

type Input struct {
	Timestamp int64  `json:"timestamp"`
	Balance   int64  `json:"balance"`
	PublicKey string `json:"publicKey"`
	Signature string `json:"signature"`
	r         *big.Int
	s         *big.Int
}

//Transaction entity with input output and ids
type Transaction struct {
	Id      string
	Input   Input    `json:"input"`
	Outputs []Output `json:"outputs"`
}

var (
	errAmountExceedsBalance = errors.New("Amount exceeds the balance of sender")
)

//NewTransaction returns a new transaction which has info of sender receiver amount etc
func NewTransaction(sendersWallet *Wallet, recipient string, amount int64) (*Transaction, error) {
	id := utils.GenerateUUID()
	if amount > sendersWallet.balance {
		return &Transaction{}, errAmountExceedsBalance
	}

	var outputs []Output
	outputs = append(outputs, Output{sendersWallet.balance - amount, sendersWallet.publicKey})
	outputs = append(outputs, Output{amount, recipient})
	return TransactionWithOutputs(sendersWallet, outputs, id)

}

//SignTransaction signs transaction and insert value for the input field of transaction
func (t *Transaction) SignTransaction(w *Wallet) {
	data := utils.NewSHA256ForByteData(t.StructToByteOutput())

	sig, r, s := utils.SignOutput(w.KeyPair, data)

	toAddInput := Input{utils.MakeTimestamp(), w.balance, w.publicKey, sig, r, s}

	t.Input = toAddInput

}

// StructToByteOutput just converting output struct to byte
func (t *Transaction) StructToByteOutput() []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(t.Outputs)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	return network.Bytes()
}

//VerifyTransaction ::: Verify the Transaction
func (t *Transaction) VerifyTransaction() bool {
	return utils.VerifySignature(t.Input.PublicKey, t.Input.Signature, utils.NewSHA256ForByteData(t.StructToByteOutput()), t.Input.r, t.Input.s)
}

func (t *Transaction) updateTransaction(sendersWallet *Wallet, recipient string, amount int64) error {
	var senderOutput int

	for i := range t.Outputs {
		if t.Outputs[i].Address == sendersWallet.publicKey {
			senderOutput = i

			break
		}
	}
	if amount > t.Outputs[senderOutput].Amount {
		return errAmountExceedsBalance
	}
	t.Input.Balance = t.Input.Balance - amount
	t.Outputs[senderOutput].Amount = t.Outputs[senderOutput].Amount - amount

	t.Outputs = append(t.Outputs, Output{amount, recipient})

	// Update the value of the wallet
	// sendersWallet.balance = t.Outputs[senderOutput].Amount
	// fmt.Println(sendersWallet)
	t.SignTransaction(sendersWallet)
	// As the amount is different now , the signature should not be valid anymore. Thats why we need to generate a new input object
	return nil
}

func TransactionWithOutputs(sendersWallet *Wallet, outputs []Output, id string) (*Transaction, error) {
	toReturn := &Transaction{id, Input{}, outputs}
	// fmt.Println("WHAT ABOUD DOS ???", sendersWallet)
	toReturn.SignTransaction(sendersWallet)
	return toReturn, nil
}

//blockchain wallet is a special wallet . why ? it generate signatures to confirm and authenticate transactions.
// miners shouldnt be signing the rewards for miner themselves
//RewardTransaction
func RewardTransaction(m *Wallet, bW *Wallet) (*Transaction, error) {
	id := utils.GenerateUUID()
	// fmt.Println("REWARD TRANSACTION MAYBE ???", m)
	return TransactionWithOutputs(bW, []Output{{Amount: int64(20), Address: m.publicKey}}, id)
}
