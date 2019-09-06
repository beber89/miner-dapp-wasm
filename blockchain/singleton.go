package blockchain

import "github.com/beber89/miner-dapp-wasm/chainfabric"

// blockIDGenerator is meant to emulate as a static class member for Block struct
type blockIDGenerator struct {
	lastID uint64
}

type observer struct {
	node *chainfabric.Node
}

var (
	blockIDGeneratorInstance blockIDGenerator
	observerInstance         observer
)

// GetObserver constructor for observer
func GetObserver() observer {
	if observerInstance == (observer{}) {
		nd := chainfabric.NewNode("127.0.0.1", 8081)
		observerInstance = observer{&nd}
		if success := observerInstance.node.Connect(); !success {
			panic("Could not connect to tracker")
		}
	}
	return observerInstance
}

// getBlockIDGenerator constructor for BlockIDGenerator
func getBlockIDGenerator() blockIDGenerator {
	if blockIDGeneratorInstance == (blockIDGenerator{}) {
		blockIDGeneratorInstance = blockIDGenerator{0}
	}
	return blockIDGeneratorInstance
}
func (gen blockIDGenerator) generate() uint64 {
	gen.lastID++
	return gen.lastID
}
