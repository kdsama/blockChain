package api

import (
	"blockchain/blockchain"
	"encoding/json"
	"fmt"
	"net/http"
)

type BlockChainHandler struct {
	service *blockchain.Blockchain
	p2p     *P2pServer
}
type postBlockRequest struct {
	Data string `json:"data"`
}

//NewHandler returns handler for blockchain, all the requests will be received here
func NewHandler(blockchain *blockchain.Blockchain, p2p *P2pServer) *BlockChainHandler {
	return &BlockChainHandler{blockchain, p2p}
}

func (bch *BlockChainHandler) Blocks(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		bch.getBlocks(w, req)
	case http.MethodPost:
		bch.postBlocks(w, req)
	}

}
func (bch *BlockChainHandler) getBlocks(w http.ResponseWriter, req *http.Request) {
	response := bch.service.Get()
	body, err := json.Marshal(response)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(body)

}

func (bch *BlockChainHandler) postBlocks(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var t postBlockRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))

	}
	bch.service.AddBlock(t.Data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Block was added successfully")))
	fmt.Printf("%d", len(bch.service.Get()))
	fmt.Printf("We need to sync All Chains")
	bch.p2p.syncChain()

}
