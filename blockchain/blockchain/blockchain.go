package blockchain

import (
	"blockchain/blockchain/blocks"
	"blockchain/utils"
	"errors"
	"reflect"
)

var (
	errNotLongEnough = errors.New("The Input Chain is not Long enough")
	errInvalidChain  = errors.New("Invalid Chain")
)

// Blockchain struct comprising of blocks
type Blockchain struct {
	chain []*blocks.Block
}

// NewBlockChain Create a New Block
func NewBlockChain() *Blockchain {
	genesisBlock := blocks.GenesisBlock()
	chain := []*blocks.Block{genesisBlock}
	return &Blockchain{chain}
}
func (bl *Blockchain) addBlock(data string) *blocks.Block {
	lastBlock := bl.chain[len(bl.chain)-1]
	newBlock := blocks.Mineblock(lastBlock, data)
	bl.chain = append(bl.chain, newBlock)
	return newBlock
}

func (bl *Blockchain) isValidChain(ch []*blocks.Block) bool {
	ok := reflect.DeepEqual(ch[0], blocks.GenesisBlock())
	if !ok {

		return false
	}
	for i := 1; i < len(bl.chain); i++ {

		currentBlock := ch[i]
		lastBlock := ch[i-1]
		ObjectBlock := bl.chain[i]

		if currentBlock.LastHash != lastBlock.Hash || currentBlock.Hash != utils.NewSHA256(ObjectBlock.Timestamp, ObjectBlock.LastHash, ObjectBlock.Data) {

			return false
		}

	}
	return true
}

func (bl *Blockchain) replaceChain(newChain []*blocks.Block) (*Blockchain, error) {

	if len(newChain) <= len(bl.chain) {

		return &Blockchain{}, errNotLongEnough
	} else {
		if !bl.isValidChain(newChain) {

			return &Blockchain{}, errInvalidChain
		} else {

			bl.chain = newChain
			return bl, nil
		}
	}

}
