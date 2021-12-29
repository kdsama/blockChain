package blocks

type Block struct {
	Timestamp int64
	LastHash  string
	Hash      string
	Data      string
}

func NewBlock(timestamp int64, lastHash string, hash string, data string) *Block {

	return &Block{timestamp, lastHash, hash, data}
}
