package blockchain

import (
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

func (blkchn *Blockchain) mutate(blk Block) {
	blkchn.chain = append(blkchn.chain, blk)
	NotifyListeners()
}

// GetChain returns the blocks forming the blockchain as an array
func (blkchn Blockchain) GetChain() []Block {
	return blkchn.chain
}

func listener(blkchn *Blockchain, msg string) {
	// mining happen creating new block
	// string payload shall be converted to transaction struct
	var trMsg TransactionMessage
	trMsg.Deserialize(msg)
	var blk Block
	go func() {
		blk = NewBlock(blkchn.chain[len(blkchn.chain)-1].GetHash(), trMsg.Trans, diff, trMsg.Timestamp)
	}()
	blkchn.mutate(blk)
}

// NewBlockchain creates a new struct blockchain with default values for genesis block
func NewBlockchain() *Blockchain {
	// In order to initialize user and connect to tracker
	GetObserver()
	gblock := NewGenesisBlock(1000)
	var chain []Block
	chain = append(chain, gblock)
	chn := Blockchain{chain}
	GetObserver().node.SetNewTransactionCallback(func(msg string) {
		listener(&chn, msg)
	})
	return &chn
}

// RequestTransaction initiates a transaction
func (blkchn *Blockchain) RequestTransaction(trans Transaction) {
	// observer send request to chainfabric
	trMsg := TransactionMessage{trans, uint64(time.Now().Unix())}
	GetObserver().node.SendMessage("TRANSACTION " + trMsg.Serialize())
	blk := NewBlock(blkchn.chain[len(blkchn.chain)-1].GetHash(), trMsg.Trans, diff, uint64(time.Now().Unix()))

	blkchn.mutate(blk)
}

// This section of blockchain.go concerned with drawing the blockchain onto the webpage

// Draw returns the js Object which is going to be added to the dom
func (blkchn Blockchain) Draw() js.Value {
	document := js.Global().Get("document")
	div := document.Call("createElement", "div")
	div.Set("id", js.ValueOf("chain"))
	for _, b := range blkchn.GetChain() {
		div.Call("appendChild", b.Draw())
	}
	return div
}

// ------------------------------------------------------------------------------------------
