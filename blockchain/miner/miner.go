package miner

import (
	"blockchain/blockchain"
	"blockchain/blockchain/blocks"
	"blockchain/crypto"
	"blockchain/ws"
	"encoding/json"
	"fmt"
)

type Miner struct {
	Blockchain *blockchain.Blockchain
	Tp         *crypto.TransactionPool
	W          *crypto.Wallet
	P2p        *ws.P2pServer
}

// var MINER_REWARD = int64(20)

func NewMiner(Blockchain *blockchain.Blockchain, Tp *crypto.TransactionPool, W *crypto.Wallet, p2p *ws.P2pServer) *Miner {
	return &Miner{Blockchain, Tp, W, p2p}
}

func (m *Miner) Mine() *blocks.Block {
	validTransactions := m.Tp.ValidTransactions()
	cReward, err := crypto.RewardTransaction(m.W, crypto.NewBlockChainWallet())
	if err != nil {
		fmt.Println("Did not expect an error here ", err)
	}
	validTransactions = append(validTransactions, *cReward)
	body, err := json.Marshal(validTransactions)
	if err != nil {

		fmt.Println("Marshalling error", err)

	}
	block := m.Blockchain.AddBlock(string(body))
	m.P2p.SyncChain()
	m.P2p.BroadcastClearTransaction()
	return block
}
