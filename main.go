package main

import (
	"fmt"
	"syscall/js"

	"github.com/beber89/miner-dapp-wasm/wallet"
)

var aliceWlt = wallet.NewWallet("Alice")

func reward(v js.Value, i []js.Value) interface{} {
	aliceWlt.Reward(10)
	js.Global().Set("output", 10)
	return js.ValueOf(0)
}

func registerCallbacks() {
	js.Global().Set("reward", js.FuncOf(reward))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	// nd := chainfabric.NewNode("127.0.0.1", 8081)
	// nd.Connect()
	<-c
}
