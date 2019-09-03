package blockchain

import (
	"fmt"
	"time"
)

// const diff uint8 = 22
const diff uint8 = 16

type Blockchain struct {
	chain []Block
}

func (chn Blockchain) GetChain() []Block {
	return chn.chain
}

func (chn *Blockchain) listener(msg string) {
	// mining happen creating new block
	// string payload shall be converted to transaction struct
	var trMsg TransactionMessage
	trMsg.Deserialize(msg)
	blk := NewBlock(chn.chain[len(chn.chain)-1].GetHash(), trMsg.Trans, diff, trMsg.Timestamp)
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", blk), blk)
	chn.chain = append(chn.chain, blk)
}

func NewBlockchain() Blockchain {
	// In order to initialize user and connect to tracker
	GetObserver()
	gblock := NewGenesisBlock(1000)
	var chain []Block
	chain = append(chain, gblock)
	chn := Blockchain{chain}
	GetObserver().node.SetNewTransactionCallback(chn.listener)
	return chn
}

func (chn *Blockchain) RequestTransaction(trans Transaction) {
	// observer send request to chainfabric
	trMsg := TransactionMessage{trans, uint64(time.Now().Unix())}
	GetObserver().node.SendMessage("TRANSACTION " + trMsg.Serialize())
	blk := NewBlock(chn.chain[len(chn.chain)-1].GetHash(), trMsg.Trans, diff, uint64(time.Now().Unix()))
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", blk), blk)
	chn.chain = append(chn.chain, blk)
}
