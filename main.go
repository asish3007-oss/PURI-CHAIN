package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
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

var Blockchain []Block

func calculateHash(block Block) string {
    record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) Block {
    var newBlock Block
    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = time.Now().String()
    newBlock.Data = data
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Nonce = 0

    for {
        newBlock.Hash = calculateHash(newBlock)
        if newBlock.Hash[:4] == "0000" { // mining difficulty
            break
        } else {
            newBlock.Nonce++
        }
    }

    return newBlock
}

func main() {
    genesisBlock := Block{0, time.Now().String(), "Genesis Block", "0", "", 0}
    genesisBlock.Hash = calculateHash(genesisBlock)
    Blockchain = append(Blockchain, genesisBlock)

    fmt.Println("Mining Block 1...")
    newBlock := generateBlock(genesisBlock, "PURI-CHAIN Block")
    Blockchain = append(Blockchain, newBlock)

    fmt.Println(Blockchain)
}
