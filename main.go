package main

import "github.com/blockchain/blockchain"
import 	"github.com/blockchain/blocks"

func main(){
	//Block
	block1 := blocks.Block{[]byte("gooz"), []byte{}, 0, 0,}
	block2 := blocks.Block{[]byte("kir"), []byte{}, 0, 0,}
	block3 := blocks.Block{[]byte("kos"), []byte{}, 0, 0,}
	block4 := blocks.Block{[]byte("kon"), []byte{}, 0, 0,}
	block5 := blocks.Block{[]byte("mame"), []byte{}, 0, 0,}
	block1.HashBlock()
	block2.HashBlock()
	block3.HashBlock()
	block4.HashBlock()
	block5.HashBlock()
	chain := blockchain.Blockchain{
		[]*blocks.Block{},
		 []byte{},
	}
	chain.AddBlock(&block1)
	chain.AddBlock(&block2)
	chain.AddBlock(&block3)
	chain.AddBlock(&block4)
	chain.AddBlock(&block5)
	chain.PrintChain()
	
}