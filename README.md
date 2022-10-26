# A Basic Proof Of Work Golang Blockchain - Integrated With SQLITE

# Features
    - Hashing and adding new blocks
    - Integrated with SQLITE
    - Proof of work consensus
    - Good for getting a basic idea of how blockchain works in general

# Run Command:
go run .
to run main.go, this will create a Blockchain.go file in DB folder with a table called 'Blocks'
each row of this table has this columns :
- Number => block number which is also used as primary key
- Hash => hash of the block which is calculated using SHA256 of block data + difficulty + hash of last block 
- Data => a byte array that contains informations contained in the block
- Difficulty => Block diffculty which affects our consensus and how hard it is to calculate a hash
- Nonce => Proof of work number calculated which is also used to validate the block

# Folders  
    - blockchain => includes the core blockchain logic, for adding blocks, starting a new blockchain
    - blocks => includes the core logic for a block, Hashing data, Calculating PoW
    - consensus => includes logic for Proof of work consensus algorithm
    - DB => includes a SQLITE database which stores blocks
    - utils => extra and general functions, like error handling function

# Info
    - this is a core and basic implementation, i will be completing this project and adding more advanced features
