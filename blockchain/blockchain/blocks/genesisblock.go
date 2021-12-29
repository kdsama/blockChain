package blocks

// Genesis Block
// Origin of the blockchain
// Hardcoded values in block
// first real block will have lastHash of Genesis Block
func GenesisBlock() *Block {
	return NewBlock(123456, "BlockTry1", "BlockTry2", "firstblock")
}
