package mychain

import (
	"time"
)

type BlockChain interface {
	AddBlock(block Block)
	Validate() bool
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
	b := block.Mine(c.difficulty)
	c.chain = append(c.chain, b)
}

func (m *MyChain) Validate() bool {
	return true
}
