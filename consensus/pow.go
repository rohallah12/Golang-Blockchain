package consensus

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"

	"github.com/blockchain/utils"
)

const Difficulty = 12

/*
================Calculating Proof Of Work================
1- finding a number below 2 ** (256 - difficulty)
2- attaching that number to block PoW field
3- validating the block to check if the hash is below target number or not
*/

/*
How To Work?

	1- Incraese nonce each time and hash it with block byte data and compare against target
*/
func Work(data []byte) (int64, [32]byte) {
	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-Difficulty))

	var nonce int64
	var value big.Int
	var intHash [32]byte

	nonce = 0
	for nonce < math.MaxInt64 {
		data := initData(nonce, data)
		intHash = sha256.Sum256(data)
		value.SetBytes(intHash[:])
		if value.Cmp(target) == -1 {
			break
		} else {
			nonce += 1
		}
		fmt.Printf("\rPoW : %x", intHash)
	}

	return nonce, intHash
}

func Validate(pow int64, data []byte, target_hash []byte) bool {
	//calculating both target and pow number
	target := big.NewInt(1)
	target = target.Lsh(target, uint(256-Difficulty))

	//First check if pow is less than target
	if pow < target.Int64() {
		return false
	}

	//Second check if the nonce hashed with block data gives us the block hash (the block sha256 checksum without the nonce)
	//to do this we calculate the sha256 hash of block's data
	hashData := sha256.Sum256(initData(pow, data))

	if bytes.Compare(hashData[:], target_hash) != 0 {
		return false
	}

	return true
}

func initData(nonce int64, block_data []byte) []byte {
	nonceBytes := toHex(nonce)
	init_data := [][]byte{
		nonceBytes,
		block_data,
	}
	data := bytes.Join(init_data, []byte{})
	return data
}

func toHex(nonce int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, nonce)
	utils.HandleError((err))
	return buffer.Bytes()
}
