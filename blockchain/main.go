package main

import (
	"blockchain/api"
	"blockchain/blockchain"
	"blockchain/crypto"
	"blockchain/miner"
	"blockchain/ws"
	"log"
	"net/http"
	"os"
)

func main() {
	// fmt.Println("This project is for Block Chain development. The course is in Nodejs , I will be writing it in Go language though :D  ")
	// // nowTime := utils.MakeTimestamp()
	// newBlock := blocks.NewBlock(utils.MakeTimestamp(), "aasdasdas", "aazxzxzx", "dddqwewqew")
	// nextBlock := blocks.Mineblock(newBlock, "YOLO DATA")
	// fmt.Println(nextBlock)
	// l := utils.NewSHA256(123123, "asdasdasdasd", "asdasdqweqwe13123wesa")
	// // y := hex.EncodeToString([]byte(l))
	// fmt.Println(l)
	bl := blockchain.NewBlockChain()
	// bl.AddBlock("Yamla")
	// bl.AddBlock("Pagla")
	tp := crypto.NewTransactionPool([]crypto.Transaction{})
	wallet := crypto.NewWallet()
	wsHandler := ws.NewP2pServer(bl, tp)
	handler := api.NewBlockChainHandler(bl, wsHandler)
	cHandler := api.NewCryptoHandler(wallet, tp, wsHandler, bl)
	mine := miner.NewMiner(bl, tp, wallet, wsHandler)
	mHandler := api.NewMinerHandler(mine)
	http.HandleFunc("/blocks", handler.Blocks)
	http.HandleFunc("/ws", wsHandler.Listen)
	http.HandleFunc("/transactions", cHandler.Transactions)
	http.HandleFunc("/transact", cHandler.Transactions)
	http.HandleFunc("/public-key", cHandler.PublicKey)
	http.HandleFunc("/mine-transactions", mHandler.Mine)
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}
