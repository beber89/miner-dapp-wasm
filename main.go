package main

import (
	"fmt"

	"github.com/beber89/miner-dapp-wasm/blockchain"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	var gblock = blockchain.NewGenesisBlock(1000)
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", gblock), gblock)

	fmt.Printf("blockhash: %x\n", gblock.GetHash())

	tr := blockchain.Transaction{"Alice", "Bob", 50}
	var nblock = blockchain.NewBlock(gblock.GetHash(), tr, 21, 1010)
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", nblock), nblock)
}
