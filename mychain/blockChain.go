package mychain

import (
	"time"
)

type BlockChain interface {
	AddBlock(block *Block)
	Verify() bool
}

type MyChain struct {
	chain      []*Block
	difficulty int
}

func NewMyChain() BlockChain {
	return &MyChain{
		chain:      []*Block{genesisBlock()},
		difficulty: 3,
	}
}

func genesisBlock() *Block {
	b := &Block{
		Timestamp:    time.Now(),
		Data:         make([]byte, 0),
		previousHash: "000empty",
	}
	b.hash = b.ComputeHash()

	return b
}

func (c *MyChain) AddBlock(block *Block) {
	block.previousHash = c.chain[len(c.chain)-1].hash
	block.Mine(c.difficulty)
	c.chain = append(c.chain, block)
}

func (m *MyChain) Verify() bool {
	for i := 1; i < len(m.chain); i++ {
		currBlock := m.chain[i]
		prevBlock := m.chain[i-1]

		if currBlock.hash != currBlock.ComputeHash() {
			return false
		}
		if currBlock.previousHash != prevBlock.ComputeHash() {
			return false
		}
	}

	return true
}
