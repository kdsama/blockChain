package api

import (
	"blockchain/miner"
	"fmt"
	"net/http"
)

type MinerHandler struct {
	miner *miner.Miner
}

func NewMinerHandler(miner *miner.Miner) *MinerHandler {
	return &MinerHandler{miner}
}
func (mh *MinerHandler) Mine(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		mh.mineBlock(w, req)
		// case http.MethodPost:
		// 	mh.postBlocks(w, req)
		// }
	}
}

func (mh *MinerHandler) mineBlock(w http.ResponseWriter, req *http.Request) {

	fmt.Println("LEts check the transactionPool")
	fmt.Println(mh.miner.Tp.Transactions)

	mh.miner.Mine()
	fmt.Println("LEts check the transactionPool")
	fmt.Println(mh.miner.Tp.Transactions)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintln("Transactions were mined successfully")))
}

// func (bch *BlockChainHandler) postBlocks(w http.ResponseWriter, req *http.Request) {

// 	decoder := json.NewDecoder(req.Body)
// 	var t postBlockRequest
// 	err := decoder.Decode(&t)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(fmt.Sprintln(http.StatusBadRequest)))

// 	}
// 	bch.service.AddBlock(t.Data)
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(fmt.Sprintln("Block was added successfully")))
// 	fmt.Printf("%d", len(bch.service.Get()))
// 	fmt.Printf("We need to sync All Chains")
// 	bch.p2p.SyncChain()

// }
