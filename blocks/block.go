package blocks

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"

	"github.com/blockchain/consensus"
	"github.com/blockchain/utils"
)

/*
1- Find Proof Of Work Of The Block
2- Validat before adding the block
3- Add The Block To Blockchain if validated
*/

type Block struct {
	Data       []byte
	Hash       []byte
	Difficulty int64
	PoW        int64
}

func (b *Block) HashBlock(LH []byte) []byte {
	finalData := bytes.Join([][]byte{LH, b.Data, toHex(b.Difficulty)}, []byte{})
	nonce, hash := consensus.Work(finalData)
	b.PoW = nonce
	b.Hash = hash[:]
	b.Difficulty = consensus.Difficulty
	return hash[:]
}

func (b *Block) Serialize() []byte {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)

	utils.HandleError(err)
	return buffer.Bytes()
}

func Desrialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode((&block))

	utils.HandleError(err)

	return &block
}

func toHex(nonce int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, nonce)
	utils.HandleError((err))
	return buffer.Bytes()
}