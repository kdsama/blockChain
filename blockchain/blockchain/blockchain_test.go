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
	got := newChain.Chain[0].Data

	if want != got {
		t.Errorf("Wanted data of genesis to be %v, but got %v", want, got)
	}

	if want != blocks.GenesisBlock().Data {
		t.Errorf("Wanted data of genesis to be %v, but got %v", want, got)
	}

	t.Run("New added block should be at the end of the blockchain", func(t *testing.T) {

		newChain.AddBlock("Yahoo")
		want := "Yahoo"
		got := newChain.Chain[len(newChain.Chain)-1].Data

		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

}

func TestIsValidChain(t *testing.T) {
	t.Run("Genesis BLock: Verification should return true", func(t *testing.T) {
		newChain := NewBlockChain()
		want := true
		l := newChain.Chain
		newChain2 := NewBlockChain()
		got := newChain2.isValidChain(l)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("GenesisBlock: Verification should return false", func(t *testing.T) {
		newChain := NewBlockChain()
		want := false
		l := newChain.Chain
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
		chain1.AddBlock("Second")
		chain2.AddBlock("ass")
		chain1.AddBlock("third")
		chain2.AddBlock("third")
		want := false
		got := chain1.isValidChain(chain2.Chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Similar Data, should return true", func(t *testing.T) {

		chain1 := NewBlockChain()

		chain1.AddBlock("second")

		chain1.AddBlock("third")

		chain1.AddBlock("third")

		chain1.AddBlock("fourth")
		// fmt.Println(chain1.Chain)
		want := true
		got := chain1.isValidChain(chain1.Chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})
}
func TestReplaceChain(t *testing.T) {
	t.Run("Try to replace chain with an invalidChain", func(t *testing.T) {

		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.AddBlock("Second")
		chain2.AddBlock("ass")
		chain1.AddBlock("third")
		chain2.AddBlock("third")
		chain2.AddBlock("fourth")
		want := errInvalidChain
		_, got := chain1.ReplaceChain(chain2.Chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Not Long Enough Chain", func(t *testing.T) {

		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.AddBlock("Second")
		chain2.AddBlock("ass")
		chain1.AddBlock("third")

		want := errNotLongEnough
		_, got := chain1.ReplaceChain(chain2.Chain)
		if want != got {
			t.Errorf("Expected data to be %v , but got %v", want, got)
		}

	})

	t.Run("Input a valid Chain", func(t *testing.T) {
		chain1 := NewBlockChain()
		chain2 := NewBlockChain()
		chain1.AddBlock("Second")

		chain1.AddBlock("third")
		chain2.Chain = chain1.Chain
		chain2.AddBlock("Fourth")
		want := len(chain2.Chain)
		// fmt.Printf("%v \n %v", chain1.Chain[1], chain2.Chain[1])
		replaceChainResult, err := chain1.ReplaceChain(chain2.Chain)
		if err != nil {
			fmt.Println(err)
			// t.Error("Didnot Expect an error")
			// t.Error(fmt.Sprintf("%v", err))
		}
		got := len(replaceChainResult.Chain)
		if got != want {
			t.Errorf("Expected total length to be %v, but got %v", want, got)
		}
	})
}
