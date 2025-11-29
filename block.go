package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp +
		block.Data + block.PrevHash + strconv.Itoa(block.Nonce)

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string, difficulty int) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Nonce = 0

	for {
		hash := calculateHash(newBlock)
		if hash[:difficulty] == prefix(difficulty) {
			newBlock.Hash = hash
			break
		}
		newBlock.Nonce++
	}

	return newBlock
}

func prefix(d int) string {
	p := ""
	for i := 0; i < d; i++ {
		p += "0"
	}
	return p
}
