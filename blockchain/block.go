package blockchain

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"syscall/js"
	"time"
)

// Block is the building entity for the blockchain
type Block struct {
	id            uint64
	lastBlockHash [32]byte
	nonce         uint64
	payload       Transaction
	difficulty    uint8
	timestamp     uint64
	hash          [32]byte
}

// NewBlock constructor for Block
func NewBlock(
	lastBlockHash [32]byte, payload Transaction,
	difficulty uint8, timestamp uint64) Block {
	var slice = make([]byte, 32)
	var hash [32]byte
	copy(hash[:], slice)
	blk := Block{getBlockIDGenerator().generate(),
		lastBlockHash,
		0,
		payload,
		difficulty,
		timestamp,
		hash}
	blk.mine()
	return blk
}

// GetNonce returns the nonce that solved the block
func (blk Block) GetNonce() uint64 {
	return blk.nonce
}

// GetLastBlockHash returns the hash of the previous block
func (blk Block) GetLastBlockHash() [32]byte {
	return blk.lastBlockHash
}

// NewGenesisBlock constructor for Block if Genesis
func NewGenesisBlock(timestamp uint64) Block {
	var slice = make([]byte, 32)
	var hash [32]byte
	copy(hash[:], slice)
	emptyTransaction := Transaction{"Coinbase", "Coinbase", 0}
	blk := Block{getBlockIDGenerator().generate(),
		hash,
		0,
		emptyTransaction,
		0,
		timestamp,
		hash}
	return blk
}

// GetHash ...
func (blk *Block) GetHash() [32]byte {
	return blk.hash
}

// GetTransaction returns the transaction described by this block
func (blk Block) GetTransaction() Transaction {
	return blk.payload
}

// doHash calculates the hash of the block
// and assigns the corresponding hash field
func (blk *Block) doHash() {
	type toBeHashed struct {
		id            uint64
		lastBlockHash [32]byte
		nonce         uint64
		payload       Transaction
		difficulty    uint8
		timestamp     uint64
	}
	h := sha256.New()
	var blkToBeHashed = toBeHashed{
		blk.id, blk.lastBlockHash,
		blk.nonce, blk.payload,
		blk.difficulty, blk.timestamp}
	byteBlk := fmt.Sprintf("%v", blkToBeHashed)
	h.Write([]byte(byteBlk))
	// transferring bits from sum into blk.hash
	copy(blk.hash[:], h.Sum(nil))
}

func (blk *Block) hashValid() bool {
	hash := blk.hash
	// put in mind Little Endian
	// converting the 8 most significant bytes of hash to one number int
	hashAsInt := uint64(0)
	for i := uint8(1); i <= uint8(8); i++ {
		hashAsInt = uint64(hash[32-i])<<((8-i)*8) + hashAsInt
	}
	diff := (uint64(1)<<63)>>(blk.difficulty-1) - 1
	return hashAsInt <= diff
}

func (blk *Block) tryNonce(nnc uint64) bool {
	blk.nonce = nnc
	blk.doHash()
	return blk.hashValid()
}

func (blk *Block) mine() {
	rnd := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for nnc, i := uint64(0), 0; nnc <= ^uint64(0); nnc, i = nnc+uint64(rnd.Intn(5)+1), i+1 {
		if !blk.tryNonce(nnc) {
			if i%500 == 0 {
				fmt.Println("Go sleep ")
				<-time.After(70 * time.Microsecond)
				fmt.Println("woke from sleep  ")
				if !GetObserver().node.ResponseEmpty() {
					if blk.tryNonce(GetObserver().node.GetResponse()) {
						js.Global().Call("alert", fmt.Sprintf("nonce mined by other user: %d", nnc))
						break
					}
				}
			}
			continue
			// select {
			// case resp := <-GetObserver().node.Chnl:
			// 	fmt.Println("got signal from chnl")
			// 	nnnc, _ := strconv.Atoi(resp)
			// 	if blk.tryNonce(uint64(nnnc)) {
			// 		js.Global().Call("alert", fmt.Sprintf("nonce mined by other user: %d", nnnc))
			// 		break
			// 	} else {
			// 		continue
			// 	}
			// default:
			// 	continue
			// }
		}
		js.Global().Call("alert", fmt.Sprintf("nonce mined successfully by this user %d\n", blk.nonce))
		// found hash > notify other nodes
		(GetObserver().node).SendResponse(nnc)
		break
	}
}

// This section of block.go concerned with drawing the block onto the webpage ---

type blockShape struct {
	cells []js.Value
}

func (v *blockShape) newCellContent(content string) {
	document := js.Global().Get("document")
	cell := document.Call("createElement", "div")
	cell.Set("className", js.ValueOf("cell color-lightblue"))
	cell.Set("innerHTML", content)
	v.cells = append(v.cells, cell)
}

// Draw returns the js Object which is going to be added to the dom
func (blk Block) Draw() js.Value {
	shape := blockShape{make([]js.Value, 0)}
	tr := blk.GetTransaction()
	shape.newCellContent(fmt.Sprintf("%s  ->  %s: %0.2f", tr.From, tr.To, tr.Amount))
	shape.newCellContent(fmt.Sprintf("%x", blk.GetLastBlockHash()))
	shape.newCellContent(fmt.Sprintf("%d", blk.GetNonce()))
	shape.newCellContent(fmt.Sprintf("%x", blk.GetHash()))

	document := js.Global().Get("document")
	div := document.Call("createElement", "div")
	div.Set("className", js.ValueOf("block"))
	for _, c := range shape.cells {
		div.Call("appendChild", c)
	}
	return div
}

// ------------------------------------------------------------------------------------------
