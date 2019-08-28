package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/beber89/miner-dapp-wasm/fabricnet"
)

// blockIDGenerator is meant to emulate as a static class member for Block struct
type blockIDGenerator struct {
	lastID uint64
}

type observer struct {
	node fabricnet.Node
}

var instance blockIDGenerator
var observerInstance observer

// GetObserver constructor for observer
func GetObserver() observer {
	if observerInstance == (observer{}) {
		observerInstance = observer{fabricnet.NewNode("127.0.0.1", 8081)}
		go observerInstance.node.Connect()
	}
	return observerInstance
}

// getBlockIDGenerator constructor for BlockIDGenerator
func getBlockIDGenerator() blockIDGenerator {
	if instance == (blockIDGenerator{}) {
		instance = blockIDGenerator{0}
	}
	return instance
}
func (gen blockIDGenerator) generate() uint64 {
	gen.lastID++
	return gen.lastID
}

// Block is the building entity for the blockchain
type Block struct {
	id            uint64
	lastBlockHash [32]byte
	nonce         uint64
	payload       string
	difficulty    uint8
	timestamp     uint64
	hash          [32]byte
}

// NewBlock constructor for Block
func NewBlock(
	lastBlockHash [32]byte, payload string,
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

// NewGenesisBlock constructor for Block if Genesis
func NewGenesisBlock(
	payload string, timestamp uint64) Block {
	var slice = make([]byte, 32)
	var hash [32]byte
	copy(hash[:], slice)
	blk := Block{getBlockIDGenerator().generate(),
		hash,
		0,
		payload,
		0,
		timestamp,
		hash}
	return blk
}

// GetHash ...
func (blk *Block) GetHash() [32]byte {
	return blk.hash
}

// doHash calculates the hash of the block
// and assigns the corresponding hash field
func (blk *Block) doHash() {
	type toBeHashed struct {
		id            uint64
		lastBlockHash [32]byte
		nonce         uint64
		payload       string
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
	for nnc := uint64(0); nnc <= ^uint64(0); nnc++ {
		if !blk.tryNonce(nnc) {
			if !GetObserver().node.ResponseEmpty() {
				if blk.tryNonce(GetObserver().node.GetResponse()) {
					fmt.Printf("nonce is %d\n", blk.nonce)
					break
				}
			}
			continue
		}
		fmt.Printf("I got it, nonce is %s\n", fmt.Sprintf("%d", blk.nonce))
		// found hash > notify other nodes
		(GetObserver().node).SendResponse(nnc)
		break
	}
}
