package mychain

import (
	"bytes"
	"encoding/binary"
	"time"
)

type Block interface {
	ComputeHash() string
	Mine(difficulty int) Block
}

type MyBlock struct {
	Timestamp    time.Time
	Data         []byte
	Hash         string
	PreviousHash string
	nonce        int64
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
	// TODO difficulty
	expected := "000"
	for b.Hash[:difficutly] == expected {
		b.nonce++
		b.Hash = b.ComputeHash()
	}

	return b
}
