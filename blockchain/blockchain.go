package blockchain

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/blockchain/blocks"
	"github.com/blockchain/consensus"
	"github.com/blockchain/utils"
	_ "github.com/mattn/go-sqlite3"
)

type Blockchain struct {
	DBpath   string
	LastHash []byte
}

func (chain *Blockchain) StartChain(path string) {
	//if chain exists => open it and set Last Hash
	//if chain doesnt exist => create a new one => run InitChain()
	chainExists := chainExist(path)
	if chainExists {
		fmt.Println("Using previous chain!")
		chain.DBpath = path
		chain.updateLastHash()
	} else {
		fmt.Println("chain does not exist, creating a new one!")
		chain.InitChain(path)
	}
}

func (chain *Blockchain) InitChain(path string) *Blockchain {
	genesis_block := blocks.Block{[]byte("Genesis"), []byte{}, 0, 0}
	genesis_block.HashBlock([]byte{})
	chain.createChainDB(path, genesis_block)
	return chain
}

func (chain *Blockchain) createChainDB(path string, genesis blocks.Block) {
	//getting DB
	database, err := sql.Open("sqlite3", path)
	utils.HandleError(err)
	//Creating one if not exist
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS `Blocks` (`number` INTEGER PRIMARY KEY, `hash` VARCHAR(64), `data` VARCHAR(64), `nonce` INTEGER, `difficulty` INTEGER)")
	utils.HandleError(err)
	statement.Exec()

	//Inserting genesis block
	statement, err = database.Prepare("INSERT INTO `Blocks` (hash, data, nonce, difficulty) VALUES(?,?,?,?)")
	utils.HandleError(err)
	statement.Exec(hex.EncodeToString(genesis.Hash), string(genesis.Data), genesis.PoW, genesis.Difficulty)
	database.Close()
	chain.DBpath = path
	chain.LastHash = genesis.Hash
}

func (chain *Blockchain) updateLastHash() {
	database, err := sql.Open("sqlite3", chain.DBpath)
	rows, err := database.Query("SELECT `hash` FROM `Blocks` ORDER BY `number` DESC LIMIT 1;")
	utils.HandleError(err)
	var Hash string
	for rows.Next() {
		rows.Scan(&Hash)
		fmt.Printf("Last hash is %s", Hash)
	}
	LH, err := hex.DecodeString(Hash)
	chain.LastHash = LH
	utils.HandleError(err)
	database.Close()
}

func chainExist(path string) bool {
	fmt.Print("Checking... \n")
	cd, err := sql.Open("sqlite3", path)
	utils.HandleError(err)
	items, err := cd.Query("SELECT * FROM `Blocks`")
	cd.Close()
	if items == nil {
		return false
	}
	return true
}

// Blockchain operations
func (chain *Blockchain) AddBlock(block *blocks.Block) {
	//First validate block details
	finalData := bytes.Join([][]byte{chain.LastHash, block.Data, toHex(block.Difficulty)}, []byte{})
	isValidBlock := consensus.Validate(block.PoW, finalData, block.Hash)
	if isValidBlock {
		//Inserting genesis block
		database, err := sql.Open("sqlite3", chain.DBpath)
		statement, err := database.Prepare("INSERT INTO `Blocks` (hash, data, nonce, difficulty) VALUES(?,?,?,?)")
		utils.HandleError(err)
		statement.Exec(hex.EncodeToString(block.Hash), string(block.Data), block.PoW, block.Difficulty)
		database.Close()
		chain.updateLastHash()
	} else {
		fmt.Println("Block is not validated!")
	}
}

func (chain Blockchain) PrintChain() {
	database, err := sql.Open("sqlite3", chain.DBpath)
	utils.HandleError(err)
	rows, err := database.Query("SELECT * FROM `Blocks`")
	var number int64
	var hash string
	var data string
	var difficulty int64
	var pow int64
	fmt.Println("Printing the chain...")
	for rows.Next() {
		fmt.Println("")
		rows.Scan(&number, &hash, &data, &difficulty, &pow)
		fmt.Printf("Block Number: %d\n", number)
		fmt.Printf("Block Hash: %s\n", hash)
		fmt.Printf("Block Data: %s\n", data)
		fmt.Printf("Block Difficulty: %d\n", difficulty)
		fmt.Printf("Block Nonce: %d\n", pow)
		fmt.Printf("Block Last Hash: %x\n", chain.LastHash)
		fmt.Println("===================================")
	}
}

func toHex(nonce int64) []byte {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, nonce)
	utils.HandleError((err))
	return buffer.Bytes()
}
