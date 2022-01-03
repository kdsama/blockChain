package blocks

import (
	"blockchain/utils"
	"fmt"
	"strings"
	"testing"
)

func TestBlock(t *testing.T) {
	data := "firstblock"
	want := data
	// timestamp := utils.MakeTimestamp()
	block := GenesisBlock()

	t.Run("Data can be seen on the block", func(t *testing.T) {
		got := block.Data
		if want != got {
			t.Errorf("Wanted %s , but got %s", want, got)
		}
	})
	block2 := Mineblock(block, "New Data")
	fmt.Println(block2, block)
	t.Run("Check if hash from previous block is same as lastHash of this block", func(t *testing.T) {
		want := block.Hash
		got := block2.LastHash
		fmt.Println(want, got)
		// fmt.Println(block2.LastHash, block2.Hash, block.LastHash)
		if want != got {
			t.Errorf("Wanted %s , but got %s", want, got)
		}
	})

	t.Run("Proof of Work : Check if nonce condition is working fine or not", func(t *testing.T) {
		want := block2.Hash[0:block2.Difficulty]
		got := strings.Repeat("0", int(block2.Difficulty))
		if want != got {
			t.Errorf("Mismatch for proof of work, wanted %s, but got %s", want, got)
		}
	})

	t.Run("Adjust Difficulty Mechanism, should return a smaller difficulty level", func(t *testing.T) {
		bl := GenesisBlock()
		bl1 := Mineblock(bl, "NewData")
		timestamp := utils.MakeTimestamp()
		got := bl1.adjustDifficulty(int64(timestamp) + 36000)
		want := bl1.Difficulty - 1
		if got != want {
			t.Errorf("Expected difficulty to be %d, but got %d", want, got)
		}
	})
	t.Run("Adjust Difficulty Mechanism, should return a bigger difficulty level", func(t *testing.T) {
		bl := GenesisBlock()
		bl1 := Mineblock(bl, "NewData")
		timestamp := utils.MakeTimestamp()
		got := bl1.adjustDifficulty(int64(timestamp) + 1)
		want := bl1.Difficulty + 1
		if got != want {
			t.Errorf("Expected difficulty to be %d, but got %d", want, got)
		}
	})

}
