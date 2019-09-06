package blockchain

import (
	"fmt"
	"syscall/js"
	"time"
)

var diff uint8 = 16

// SetDifficulty sets the difficulty value for the PoW done by block mining
func SetDifficulty(difficulty uint8) {
	diff = difficulty
}

// Blockchain ..
type Blockchain struct {
	chain []Block
}

// GetChain returns the blocks forming the blockchain as an array
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

// NewBlockchain creates a new struct blockchain with default values for genesis block
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

// RequestTransaction initiates a transaction
func (chn *Blockchain) RequestTransaction(trans Transaction) {
	// observer send request to chainfabric
	trMsg := TransactionMessage{trans, uint64(time.Now().Unix())}
	GetObserver().node.SendMessage("TRANSACTION " + trMsg.Serialize())
	blk := NewBlock(chn.chain[len(chn.chain)-1].GetHash(), trMsg.Trans, diff, uint64(time.Now().Unix()))
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", blk), blk)
	chn.chain = append(chn.chain, blk)
	fmt.Printf("chain is %v\n", chn.chain)
}

// This section of blockchain.go concerned with drawing the blockchain onto the webpage

// Draw returns the js Object which is going to be added to the dom
func (chn Blockchain) Draw() js.Value {
	document := js.Global().Get("document")
	div := document.Call("createElement", "div")
	div.Set("id", js.ValueOf("chain"))
	for _, b := range chn.GetChain() {
		div.Call("appendChild", b.Draw())
	}
	return div
}

// ------------------------------------------------------------------------------------------
