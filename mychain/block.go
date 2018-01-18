package mychain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"time"
)

type Block struct {
	Timestamp    time.Time
	Data         []byte
	hash         string
	previousHash string
	nonce        int64
}

func NewBlock(timestamp time.Time, data []byte) *Block {
	b := &Block{
		Timestamp: timestamp,
		Data:      data,
	}

	b.hash = b.ComputeHash()

	return b
}

func (b *Block) ComputeHash() string {
	buff := &bytes.Buffer{}

	t, _ := b.Timestamp.MarshalBinary()
	buff.Write(t)

	buff.Write(b.Data)

	buff.Write([]byte(b.previousHash))

	nonce := make([]byte, 8)
	binary.BigEndian.PutUint64(nonce, uint64(b.nonce))
	buff.Write(nonce)

	// return buff.String()
	h := sha256.New()
	h.Write(buff.Bytes())

	return hex.EncodeToString(h.Sum(nil))
}

func (b *Block) Mine(difficutly int) {
	// TODO difficulty
	expected := "000"
	for b.hash[:difficutly] != expected {
		b.nonce++
		b.hash = b.ComputeHash()
	}
}
