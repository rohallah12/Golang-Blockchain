package main

import (
	"github.com/blockchain/blockchain"
	"github.com/blockchain/blocks"
	"github.com/blockchain/consensus"
)

func main() {
	//Block
	// block1 := blocks.Block{[]byte("data1"), []byte{}, 0, 0,}
	// block2 := blocks.Block{[]byte("data2"), []byte{}, 0, 0,}
	// block3 := blocks.Block{[]byte("data3"), []byte{}, 0, 0,}
	// block4 := blocks.Block{[]byte("data4"), []byte{}, 0, 0,}
	// block5 := blocks.Block{[]byte("data5"), []byte{}, 0, 0,}
	// block1.HashBlock()
	// block2.HashBlock()
	// block3.HashBlock()
	// block4.HashBlock()
	// block5.HashBlock()
	// chain := blockchain.Blockchain{
	// 	[]*blocks.Block{},
	// 	 []byte{},
	// }
	// chain.InitChain()
	// chain.AddBlock(&block1)
	// chain.AddBlock(&block2)
	// chain.AddBlock(&block3)
	// chain.AddBlock(&block4)
	// chain.AddBlock(&block5)
	// chain.PrintChain()
	chain := blockchain.Blockchain{
		"", []byte{},
	}
	block1 := blocks.Block{[]byte("data3"), []byte{}, consensus.Difficulty, 0}
	chain.StartChain("./DB/Blockchain.db")
	block1.HashBlock(chain.LastHash)
	chain.AddBlock(&block1)
	chain.PrintChain()
}
