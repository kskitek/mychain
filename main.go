package main

import (
	"fmt"
	"kskitek/mychain/mychain"
	"time"
)

func main() {
	bc := mychain.NewMyChain()

	b1 := mychain.NewBlock(time.Now(), []byte("test"))

	bc.AddBlock(b1)
	fmt.Println(bc.Verify())

}
