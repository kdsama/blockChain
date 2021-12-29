package blocks

import (
	"blockchain/utils"
	"time"
)

// THis is the function to be called when we have to append a new block to the blockchain
// For that We would be needing lastBlock's Hash

//Mineblock is to create a miners block , which then will be appended to a Blockchain
func Mineblock(b *Block, data string) *Block {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	lastHash := b.Hash
	hash := utils.NewSHA256(timestamp, lastHash, data)
	return NewBlock(timestamp, lastHash, hash, data)
}
