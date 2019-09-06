package web

import (
	"github.com/beber89/miner-dapp-wasm/wallet"
)

var (
	// User is the name of the owner of the Wallet
	User           = ""
	walletInstance wallet.Wallet
	isInitialized  = false
)

// GetWallet constructor for wallet
func GetWallet() *wallet.Wallet {
	if !isInitialized {
		if User != "" {
			walletInstance = wallet.NewWallet(User)
		} else {
			// Alice is the default user
			walletInstance = wallet.NewWallet("Alice")
		}
		isInitialized = true
	}
	return &walletInstance
}
