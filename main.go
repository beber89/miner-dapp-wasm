package main

import (
	"fmt"

	"github.com/beber89/miner-dapp-wasm/web"
)

func main() {
	c := make(chan struct{}, 0)
	fmt.Println("Hello, WebAssembly!")
	web.RegisterCallbacks()
	<-c
}
