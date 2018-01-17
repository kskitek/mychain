package mychain

import (
	"time"
)

type BlockChain interface {
	AddBlock(block Block)
	Verify() bool
}

type MyChain struct {
	chain      []Block
	difficulty int
}

func NewMyChain() BlockChain {
	return &MyChain{
		chain:      []Block{genesisBlock()},
		difficulty: 3,
	}
}

func genesisBlock() *MyBlock {
	b := &MyBlock{
		Timestamp:    time.Now(),
		Data:         make([]byte, 0),
		PreviousHash: "empty",
	}
	b.Hash = b.ComputeHash()

	return b
}

func (c *MyChain) AddBlock(block Block) {
	// TODO block needs to get hash of last block in chain..
	b := block.Mine(c.difficulty)
	c.chain = append(c.chain, b)
}

func (m *MyChain) Verify() bool {
	for i := 1; i < len(m.chain); i++ {
		currBlock := m.chain[i]
		prevBlock := m.chain[i-1]

		if currBlock.GetHash() != currBlock.ComputeHash() {
			return false
		}
		if currBlock.GetPreviousHash() != prevBlock.ComputeHash() {
			return false
		}
	}

	return true
}
