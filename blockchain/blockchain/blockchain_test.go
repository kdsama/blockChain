package blockchain

import (
	"blockchain/blockchain/blocks"
	"fmt"
	"testing"
)

//TestBlockchain to test the blockchain
func TestBlockchain(t *testing.T) {
	want := "firstblock"
	newChain := NewBlockChain()
	got := newChain.chain[0].Data

	if want != got {
		t.Errorf("Wanted data of genesis to be %v, but got %v", want, got)
	}

	if want != blocks.GenesisBlock().Data {
		t.Errorf("Wanted data of genesis to be %v, but got %v", want, got)
	}

	t.Run("New added block should be at the end of the blockchain", func(t *testing.T) {

		newChain.addBlock("Yahoo")
		want := "Yahoo"
		got := newChain.chain[len(newChain.chain)-1].Data

		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

}

func TestIsValidChain(t *testing.T) {
	t.Run("Genesis BLock: Verification should return true", func(t *testing.T) {
		newChain := NewBlockChain()
		want := true
		l := newChain.chain
		newChain2 := NewBlockChain()
		got := newChain2.isValidChain(l)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("GenesisBlock: Verification should return false", func(t *testing.T) {
		newChain := NewBlockChain()
		want := false
		l := newChain.chain
		newChain2 := NewBlockChain()
		l[0].Data = "NotFirstBlock"
		got := newChain2.isValidChain(l)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Changing Data of SecondBlock,should not be valid", func(t *testing.T) {
		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.addBlock("Second")
		chain2.addBlock("ass")
		chain1.addBlock("third")
		chain2.addBlock("third")
		want := false
		got := chain1.isValidChain(chain2.chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Similar Data, should return true", func(t *testing.T) {

		chain1 := NewBlockChain()

		chain1.addBlock("second")

		chain1.addBlock("third")

		chain1.addBlock("third")

		chain1.addBlock("third")

		want := true
		got := chain1.isValidChain(chain1.chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})
}
func TestReplaceChain(t *testing.T) {
	t.Run("Try to replace chain with an invalidChain", func(t *testing.T) {

		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.addBlock("Second")
		chain2.addBlock("ass")
		chain1.addBlock("third")
		chain2.addBlock("third")
		chain2.addBlock("fourth")
		fmt.Println(len(chain1.chain), len(chain2.chain))
		want := errInvalidChain
		_, got := chain1.replaceChain(chain2.chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Not Long Enough Chain", func(t *testing.T) {

		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.addBlock("Second")
		chain2.addBlock("ass")
		chain1.addBlock("third")

		want := errNotLongEnough
		_, got := chain1.replaceChain(chain2.chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Input a valid Chain", func(t *testing.T) {
		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.addBlock("Second")

		chain1.addBlock("third")
		chain2.chain = chain1.chain
		chain2.addBlock("Fourth")
		want := len(chain2.chain)
		replaceChainResult, err := chain1.replaceChain(chain2.chain)
		if err != nil {
			t.Error("Didnot Expect an error")
		}
		got := len(replaceChainResult.chain)
		if got != want {
			t.Errorf("Expected total length to be %v, but got %v", want, got)
		}
	})
}
