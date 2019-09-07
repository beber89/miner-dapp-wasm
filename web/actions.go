package web

import (
	"strconv"
	"syscall/js"
)

// RegisterCallbacks map javascript function into their go counterparts
func RegisterCallbacks() {
	js.Global().Set("reward", js.FuncOf(reward))
	js.Global().Set("buyCommodity", js.FuncOf(buyCommodity))
	js.Global().Set("networth", js.FuncOf(networth))
	js.Global().Set("changeUser", js.FuncOf(changeUser))
	js.Global().Set("sendSliderValToWasm", js.FuncOf(sendSliderValToWasm))
}

var (
	merchant string
	price    float64
)

func changeUser(this js.Value, args []js.Value) interface{} {
	if args[0].String() != "0" {
		document := js.Global().Get("document")
		document.Call("getElementById", "BuyBtn").Set("disabled", js.ValueOf(false))
		options := document.Call("getElementsByTagName", "option")
		for i := range []int{0, 1, 2} {
			options.Index(i).Set("disabled", js.ValueOf(true))
		}
		User = args[0].String()
		// this call initializes the User's wallet hence so are the node connections
		GetWallet()
		if User == "Alice" {
			merchant = "Bob"
		} else {
			merchant = "Alice"
		}
		if merchant == "Alice" {
			price = 10
			document.Call("getElementById", "BuyBtn").Set("innerHTML", "Buy Bananas  🍌")
		} else {
			price = 20
			document.Call("getElementById", "BuyBtn").Set("innerHTML", "Buy Apples  🍎")
		}
	}
	return js.Null()
}

func buyCommodity(this js.Value, args []js.Value) interface{} {
	if User == "" {
		return js.Null()
	}
	success := GetWallet().PayTo(merchant, price)
	if !success {
		js.Global().Call("alert", "Insufficient Crypsys")
	}
	return js.Null()
}

func networth(this js.Value, args []js.Value) interface{} {
	if User == "" {
		return js.Null()
	}
	return js.ValueOf(GetWallet().Networth())
}

func reward(this js.Value, args []js.Value) interface{} {
	if User == "" {
		return js.Null()
	}
	go func() {
		GetWallet().Reward(10)
	}()
	return js.Null()
}

func sendSliderValToWasm(this js.Value, args []js.Value) interface{} {
	if User == "" {
		return js.Null()
	}
	diff, _ := strconv.Atoi(args[0].String())
	GetWallet().SetDifficulty(uint8(diff))
	return js.Null()
}
