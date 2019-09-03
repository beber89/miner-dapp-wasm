package main

import (
	"fmt"
	"syscall/js"

	"github.com/beber89/miner-dapp-wasm/wallet"
	"github.com/beber89/miner-dapp-wasm/web"
)

var aliceWlt = wallet.NewWallet("Alice")

func reward(v js.Value, i []js.Value) interface{} {
	aliceWlt.Reward(10)
	web.ViewChain(aliceWlt)
	return js.ValueOf(0)
}

func registerCallbacks() {
	js.Global().Set("reward", js.FuncOf(reward))
}

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}
