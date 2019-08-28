package main

import (
	"fmt"

	"github.com/beber89/miner-dapp-wasm/blockchain"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	// var bobNode = fabricnet.NewNode("127.0.0.1", 8081)
	// go bobNode.Connect()

	// time.Sleep(5 * time.Second)
	// fmt.Println("[main] Response")
	// fmt.Println(bobNode.GetResponse())
	var gblock = blockchain.NewGenesisBlock("Hello Block", 1000)
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", gblock), gblock)

	fmt.Printf("blockhash: %x\n", gblock.GetHash())

	var nblock = blockchain.NewBlock(gblock.GetHash(), "Hello Block", 21, 1010)
	fmt.Printf("%s\n hashes to\n %x\n",
		fmt.Sprintf("%v", nblock), nblock)
}
