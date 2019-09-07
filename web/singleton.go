package web

import (
	"github.com/beber89/miner-dapp-wasm/blockchain"
)

var (
	// User is the name of the owner of the Wallet
	User           = ""
	walletInstance blockchain.Wallet
	isInitialized  = false
)

// GetWallet constructor for wallet
func GetWallet() *blockchain.Wallet {
	if !isInitialized {
		if User != "" {
			walletInstance = blockchain.NewWallet(User)
		} else {
			// Alice is the default user
			walletInstance = blockchain.NewWallet("Alice")
		}
		blockchain.NotifyListeners = RenderPage
		isInitialized = true
	}
	return &walletInstance
}
