package blockchain

type Blockchain struct {
	chain                  []Block
	newTransactionCallback func(string)
}

func (chn *Blockchain) listener(payload string) {
	// mining happen creating new block
}

func NewBlockchain() Blockchain {
	gblock := NewGenesisBlock("Genesis", 1000)
	var chain []Block
	chain = append(chain, gblock)
	chn := Blockchain{chain, nil}
	chn.newTransactionCallback = chn.listener
	return chn
}

func (chn *Blockchain) RequestTransaction(
	payload string, timestamp uint64) {
	// observer send request to chainfabric
}
