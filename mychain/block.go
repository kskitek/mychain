package mychain

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

type Block interface {
	ComputeHash() string
	Mine(difficulty int) Block
	GetHash() string
	GetPreviousHash() string
}

type MyBlock struct {
	Timestamp    time.Time
	Data         []byte
	Hash         string
	PreviousHash string
	nonce        int64
}

func NewMyBlock(timestamp time.Time, data []byte) *MyBlock {
	b := &MyBlock{
		Timestamp: timestamp,
		Data:      data,
	}

	b.Hash = b.ComputeHash()

	return b
}

func (b *MyBlock) GetHash() string {
	return b.Hash
}

func (b *MyBlock) GetPreviousHash() string {
	return b.PreviousHash
}

func (b *MyBlock) ComputeHash() string {
	buff := &bytes.Buffer{}

	t, _ := b.Timestamp.MarshalBinary()
	buff.Write(t)

	buff.Write(b.Data)

	buff.Write([]byte(b.PreviousHash))

	nonce := make([]byte, 8)
	binary.BigEndian.PutUint64(nonce, uint64(b.nonce))
	buff.Write(nonce)

	return buff.String()
}

func (b *MyBlock) Mine(difficutly int) Block {
	fmt.Printf("Mining %v\n", b)
	fmt.Printf("Mining %v\n", b.Hash)
	// TODO difficulty
	expected := "000"
	for b.Hash[:difficutly] == expected {
		b.nonce++
		b.Hash = b.ComputeHash()
	}

	return b
}
