package blockchain

import (
	"fmt"

	"github.com/blockchain/blocks"
	"github.com/blockchain/consensus"
)

type Blockchain struct {
	Blocks   []*blocks.Block
	LastHash []byte
}

func (chain *Blockchain) AddBlock(block *blocks.Block) {
	//First validate block details
	isValidBlock := consensus.Validate(block.PoW, block.Data, block.Hash)
	if isValidBlock {
		chain.Blocks = append(chain.Blocks, block)
	} else {
		fmt.Println("Block is not validated!")
	}
}

func (chain *Blockchain) PrintChain() {
	fmt.Println("\n")
	for i, block := range chain.Blocks {
		fmt.Printf("\nBlock Number #%d\n", i)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Difficulty: %d\n", block.Difficulty)
		fmt.Printf("Proof Of Work: %d\n", block.PoW)
		fmt.Printf("Data Hex: %x\n", block.Data)
		fmt.Println("===============================")
	}
}
