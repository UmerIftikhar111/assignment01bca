package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a single block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain is a slice of blocks.
var Blockchain []*Block

// NewBlock creates a new block and adds it to the blockchain.
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CreateHash(block)
	Blockchain = append(Blockchain, block)
	return block
}

// CreateHash calculates the hash of a block.
func CreateHash(b *Block) string {
	data := fmt.Sprintf("%s%d%s", b.Transaction, b.Nonce, b.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}

// DisplayBlocks prints all blocks in the blockchain.
func DisplayBlocks() {
	for i, block := range Blockchain {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n\n", block.Hash)
	}
}

// ChangeBlock changes the transaction of a specific block.
func ChangeBlock(blockIndex int, newTransaction string) {
	if blockIndex >= 0 && blockIndex < len(Blockchain) {
		Blockchain[blockIndex].Transaction = newTransaction
		Blockchain[blockIndex].Hash = CreateHash(Blockchain[blockIndex])
	}
}

// VerifyChain verifies the integrity of the blockchain.
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			return false
		}
	}
	return true
}

// CalculateHash calculates the hash of a string.
func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}
