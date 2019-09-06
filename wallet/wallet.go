package wallet

import (
	"github.com/beber89/miner-dapp-wasm/blockchain"
)

// Wallet is the interface struct to the blockchain for the user
type Wallet struct {
	username   string
	blockchain blockchain.Blockchain
}

// NewWallet creates a Wallet struct
func NewWallet(username string) Wallet {
	return Wallet{username, blockchain.NewBlockchain()}
}

// SetDifficulty sets the difficulty value for the PoW done by block mining
func (wlt Wallet) SetDifficulty(diff uint8) {
	blockchain.SetDifficulty(diff)
}

// Reward transfers amt to owner of wallet from coinbase
func (wlt *Wallet) Reward(amt float64) {
	tr := blockchain.Transaction{"Coinbase", wlt.username, amt}
	wlt.blockchain.RequestTransaction(tr)
}

// Networth traverses the blockchain to calculate user's balance
func (wlt Wallet) Networth() float64 {
	// Calculate transaction To "username" minus From "username"
	var netsum float64
	for _, blk := range wlt.blockchain.GetChain() {
		var trnsxn = blk.GetTransaction()
		if trnsxn.To == wlt.username {
			netsum = netsum + trnsxn.Amount
		}
		if trnsxn.From == wlt.username {
			netsum = netsum - trnsxn.Amount
		}
	}
	return netsum
}

// PayTo transfers amt to the username specified from the owner of this wallet
func (wlt *Wallet) PayTo(username string, amt float64) bool {
	if wlt.Networth() < amt {
		return false
	}
	tr := blockchain.Transaction{wlt.username, username, amt}
	// TODO: Generate Time here as uint64
	wlt.blockchain.RequestTransaction(tr)
	return true
}

// GetBlockchain temp func used only for testing
func (wlt Wallet) GetBlockchain() blockchain.Blockchain {
	return wlt.blockchain
}
