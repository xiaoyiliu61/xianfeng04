package main

import (
	"XianfengChain04/chain"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	blockChain:=chain.CreatechainWithGensis([]byte("hello world"))
	blockChain.CreateNewBlock([]byte("hello"))

    fmt.Println("区块链中区块的个数",len(blockChain.Blocks))

    fmt.Println("区块0的哈希值：",blockChain.Blocks[0])
    fmt.Println("区块1的哈希值：",blockChain.Blocks[1])
}
