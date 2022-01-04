package api

import (
	"blockchain/blockchain"
	"blockchain/crypto"
	"blockchain/ws"
	"encoding/json"
	"fmt"
	"net/http"
)

type CryptoHandler struct {
	w   *crypto.Wallet
	tp  *crypto.TransactionPool
	p2p *ws.P2pServer
	bl  *blockchain.Blockchain
}

type transactRequest struct {
	Recipient string
	Amount    int
}

func NewCryptoHandler(w *crypto.Wallet, tp *crypto.TransactionPool, p2p *ws.P2pServer, bl *blockchain.Blockchain) *CryptoHandler {
	return &CryptoHandler{w, tp, p2p, bl}
}
func (ch *CryptoHandler) Transactions(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		ch.getTransactions(w, req)
	case http.MethodPost:
		ch.postTransactions(w, req)
	}

}

func (ch *CryptoHandler) PublicKey(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		ch.getPublicKey(w, req)

	}

}

func (ch *CryptoHandler) getPublicKey(w http.ResponseWriter, req *http.Request) {
	var response struct {
		Key string `json:"key"`
	}
	response.Key = ch.w.GetPublicKey()
	fmt.Println(response.Key)
	body, err := json.Marshal(response)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)

}

func (ch *CryptoHandler) getTransactions(w http.ResponseWriter, req *http.Request) {
	response := ch.tp.GetTransactions()
	body, err := json.Marshal(response)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)

}

func (ch *CryptoHandler) postTransactions(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	// fmt.Println(req.Body)
	var t transactRequest

	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))

	}
	// fmt.Println(t)
	transaction, err := ch.w.CreateTransaction(t.Recipient, int64(t.Amount), ch.tp, ch.bl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	ch.p2p.BroadcastTransaction(transaction)
	// bch.service.AddBlock(t.Data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Transaction was done successfully")))

}
