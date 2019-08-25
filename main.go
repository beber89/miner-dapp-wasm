package main

import (
	"fmt"
	"time"

	"github.com/beber89/miner-dapp-wasm/fabricnet"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	var bobNode = fabricnet.NewNode("127.0.0.1", 8081)
	go bobNode.Connect()

	time.Sleep(3 * time.Second)
	fmt.Println("[main] Response")
	fmt.Println(bobNode.GetResponse())
}
