package dapp_test

import (
	"fmt"
	"testing"

	"github.com/beber89/miner-dapp-wasm/wallet"
)

func TestReward(t *testing.T) {
	wlt := wallet.NewWallet("Alice")
	wlt.Reward(100)

	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))
}

func TestPayTo(t *testing.T) {
	wlt := wallet.NewWallet("Alice")
	wlt.Reward(100)
	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))
	wlt.PayTo("Bob", 40)

	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))
}

func TestNetworth(t *testing.T) {
	wlt := wallet.NewWallet("Alice")
	wlt.Reward(100)
	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))
	wlt.PayTo("Bob", 40)
	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))

	t.Logf("Networth of Alice is\n%f crypsy\n\n", wlt.Networth())
}

func TestNotEnoughBalancePayment(t *testing.T) {
	wlt := wallet.NewWallet("Alice")
	wlt.Reward(30)
	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))
	success := wlt.PayTo("Bob", 40)

	t.Logf("chain is\n%s\n\n", fmt.Sprintf("%v", wlt.GetBlockchain().GetChain()))

	t.Logf("payment success value is %t\n\n", success)
}
