package blockchain

import (
	"fmt"
	"strconv"
	"strings"
)

const diff uint8 = 22

type Blockchain struct {
	chain []Block
}

func (chn Blockchain) GetChain() []Block {
	return chn.chain
}

func (chn *Blockchain) listener(payload string) {
	// mining happen creating new block
	// string payload shall be converted to transaction struct
	flds := strings.Fields(payload)
	timestamp, _ := strconv.Atoi(flds[1])
	blk := NewBlock(chn.chain[len(chn.chain)-1].GetHash(), payload, diff, uint64(timestamp))
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", blk), blk)
	chn.chain = append(chn.chain, blk)
}

func NewBlockchain() Blockchain {
	// In order to initialize user and connect to tracker
	GetObserver()
	gblock := NewGenesisBlock("Genesis", 1000)
	var chain []Block
	chain = append(chain, gblock)
	chn := Blockchain{chain}
	GetObserver().node.SetNewTransactionCallback(chn.listener)
	return chn
}

func (chn *Blockchain) RequestTransaction(
	trans string, timestamp uint64) {
	// observer send request to chainfabric
	msg := fmt.Sprintf("TRANSACTION %d "+trans+"\n", timestamp)
	GetObserver().node.SendMessage(msg)
	blk := NewBlock(chn.chain[len(chn.chain)-1].GetHash(), msg, diff, timestamp)
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", blk), blk)
	chn.chain = append(chn.chain, blk)
}
