package blocks

import (
	"fmt"
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

}
