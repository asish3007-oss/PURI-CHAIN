package main

import (
	"fmt"
)

func createGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: "GENESIS",
		Data:      "Genesis Block",
		PrevHash:  "0",
		Hash:      calculateHash(Block{Index: 0, Timestamp: "GENESIS", Data: "Genesis Block", PrevHash: "0", Nonce: 0}),
		Nonce:     0,
	}
}

func createBlockchain() []Block {
	var chain []Block
	genesis := createGenesisBlock()
	chain = append(chain, genesis)
	return chain
}

func addBlock(data string, difficulty int) {
	oldBlock := Blockchain[len(Blockchain)-1]
	newBlock := generateBlock(oldBlock, data, difficulty)
	Blockchain = append(Blockchain, newBlock)
}

func isChainValid(chain []Block) bool {
	for i := 1; i < len(chain); i++ {
		prev := chain[i-1]
		curr := chain[i]

		if curr.PrevHash != prev.Hash {
			return false
		}

		if curr.Hash != calculateHash(curr) {
			return false
		}
	}
	return true
}

func printChain() {
	for _, block := range Blockchain {
		fmt.Println("------ BLOCK ------")
		fmt.Println("Index:    ", block.Index)
		fmt.Println("Timestamp:", block.Timestamp)
		fmt.Println("Data:     ", block.Data)
		fmt.Println("PrevHash: ", block.PrevHash)
		fmt.Println("Hash:     ", block.Hash)
		fmt.Println("Nonce:    ", block.Nonce)
	}
}
