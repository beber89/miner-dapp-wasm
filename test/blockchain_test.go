package blockchain_test

import (
	"testing"
	"time"

	"github.com/beber89/miner-dapp-wasm/blockchain"
)

func TestBlockchainScenario(t *testing.T) {
	chn := blockchain.NewBlockchain()
	time.Sleep(1 * time.Second)
	if len(chn.GetChain()) != 1 {
		t.Error("Expected block to be added to chain ")
	} else {
		t.Logf("Genesis Block:\n%x\n", chn.GetChain()[0])
	}

	chn.RequestTransaction("Alice Bob 40", 1010)
	if len(chn.GetChain()) != 2 {
		t.Error("Expected block to be added to chain ")
	} else {
		t.Logf("New Block:\n%x\n", chn.GetChain()[1])
	}

}
