package main

import (
	"blockchain/api"
	"blockchain/blockchain"
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
	wsHandler := api.NewP2pServer(bl)
	handler := api.NewHandler(bl, wsHandler)

	http.HandleFunc("/blocks", handler.Blocks)
	http.HandleFunc("/ws", wsHandler.Listen)
	log.Fatal(http.ListenAndServe(os.Args[1], nil))
}
