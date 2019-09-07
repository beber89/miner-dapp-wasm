package main

import (
	"fmt"
	"sync"

	"github.com/beber89/miner-dapp-wasm/web"
)

var wg sync.WaitGroup

func main() {
	// c := make(chan struct{}, 0)
	wg.Add(1)
	fmt.Println("Hello, WebAssembly!")
	web.RegisterCallbacks()
	wg.Wait()
}
