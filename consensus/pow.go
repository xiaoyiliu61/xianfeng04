package consensus

import (
	"XianfengChain04/chain"
	"fmt"
)

type PoW struct {
	Block chain.Block
}

func(pos PoW)FindNonce()int64{
	fmt.Println("这是共识机制使用PoS算法的实现")
	return 0
}

