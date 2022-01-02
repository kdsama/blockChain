package blockchain

import (
	"blockchain/blockchain/blocks"
	"blockchain/utils"
	"errors"
	"fmt"
	"reflect"
)

var (
	errNotLongEnough = errors.New("The Input Chain is not Long enough")
	errInvalidChain  = errors.New("Invalid Chain")
)

// Blockchain struct comprising of blocks
type Blockchain struct {
	Chain []*blocks.Block `json:"chain"`
}

// NewBlockChain Create a New Block
func NewBlockChain() *Blockchain {
	genesisBlock := blocks.GenesisBlock()
	chain := []*blocks.Block{genesisBlock}
	return &Blockchain{chain}
}

//AddBlock adds a newblock to the blockchain
func (bl *Blockchain) AddBlock(data string) *blocks.Block {
	lastBlock := bl.Chain[len(bl.Chain)-1]
	newBlock := blocks.Mineblock(lastBlock, data)
	bl.Chain = append(bl.Chain, newBlock)
	return newBlock
}

func (bl *Blockchain) isValidChain(ch []*blocks.Block) bool {
	ok := reflect.DeepEqual(ch[0], blocks.GenesisBlock())
	fmt.Println(ch[0], blocks.GenesisBlock())
	if !ok {
		fmt.Println("RETURNING FROM HERE ?????")
		return false
	}
	for i := 1; i < len(bl.Chain); i++ {

		currentBlock := ch[i]
		lastBlock := ch[i-1]
		ObjectBlock := bl.Chain[i]
		fmt.Println(currentBlock, ObjectBlock)
		if currentBlock.LastHash != lastBlock.Hash || currentBlock.Hash != utils.NewSHA256(ObjectBlock.Timestamp, ObjectBlock.LastHash, ObjectBlock.Data, ObjectBlock.Nonce, ObjectBlock.Difficulty) {
			fmt.Println(currentBlock, ObjectBlock)
			fmt.Println(currentBlock.Hash, utils.NewSHA256(ObjectBlock.Timestamp, ObjectBlock.LastHash, ObjectBlock.Data, ObjectBlock.Nonce, ObjectBlock.Difficulty))
			return false
		}

	}
	return true
}

// Get returns the whole chain
func (bl *Blockchain) Get() []*blocks.Block {
	return bl.Chain
}

// ReplaceChain :: Replace chain with a new chain if constraints meet (isValid, bigger chain)
func (bl *Blockchain) ReplaceChain(newChain []*blocks.Block) (*Blockchain, error) {

	if len(newChain) <= len(bl.Chain) {

		return &Blockchain{}, errNotLongEnough
	} else {
		if !bl.isValidChain(newChain) {

			return &Blockchain{}, errInvalidChain
		} else {

			bl.Chain = newChain
			return bl, nil
		}
	}

}
