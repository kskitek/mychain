package main

import (
	"fmt"
	"kskitek/mychain/mychain"
	"time"
)

func main() {
	bc := mychain.NewMyChain()

	b1 := &mychain.MyBlock{
		Data:      []byte("test"),
		Timestamp: time.Now(),
	}

	bc.AddBlock(b1)
	fmt.Println(bc.Verify())

}
