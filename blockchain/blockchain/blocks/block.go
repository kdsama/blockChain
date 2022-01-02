package blocks

import (
	"blockchain/utils"
	"strings"
	"time"
)

type Block struct {
	Timestamp  int64  `json:"timestamp"`
	LastHash   string `json:"lasthash"`
	Hash       string `json:"hash"`
	Data       string `json:"data"`
	Nonce      int64  `json:"nonce"`
	Difficulty int64  `json:"difficulty"`
}

var DIFFICULTY = 3
var Minerate = 3000

func NewBlock(timestamp int64, lastHash string, hash string, data string, nonce int64, Difficulty int64) *Block {

	return &Block{timestamp, lastHash, hash, data, nonce, Difficulty}
}

// Genesis Block
// Origin of the blockchain
// Hardcoded values in block
// first real block will have lastHash of Genesis Block
func GenesisBlock() *Block {
	// timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	return NewBlock(123456, "check123", "BlockTry2", "firstblock", 1, int64(DIFFICULTY))
}

// THis is the function to be called when we have to append a new block to the blockchain
// For that We would be needing lastBlock's Hash

//Mineblock is to create a miners block , which then will be appended to a Blockchain
func Mineblock(b *Block, data string) *Block {
	var timestamp int64
	var nonce int64
	nonce = 0
	var hash string
	lastHash := b.Hash
	difficulty := b.Difficulty
	for {
		nonce++
		timestamp = time.Now().UnixNano() / int64(time.Millisecond)
		// fmt.Println("NONCEIS", nonce)

		// Nonce is being used here. this changing value is here to generate a hash which has number of leading zeroes equal to the difficulty
		difficulty = b.adjustDifficulty(timestamp)
		hash = utils.NewSHA256(timestamp, lastHash, data, nonce, int64(difficulty))
		// fmt.Println("DIFFICULTY", difficulty)
		// fmt.Println(hash[0:difficulty], strings.Repeat("0", int(difficulty)))
		if hash[0:difficulty] == strings.Repeat("0", int(difficulty)) {
			// fmt.Println(hash[0:DIFFICULTY], strings.Repeat("0", int(DIFFICULTY)))
			// fmt.Println("Difficulty added to ", data, ":::::", difficulty)
			return NewBlock(timestamp, lastHash, hash, data, nonce, difficulty)

		}

	}

}

func (b *Block) adjustDifficulty(timestampNow int64) int64 {
	// fmt.Println(b.Data, difficulty, timestampNow, b.Timestamp)

	if b.Timestamp+int64(Minerate) > timestampNow {
		return b.Difficulty + 1
	}
	return b.Difficulty - 1

}
