package blockchain_test

import (
	"testing"
	"time"

	"github.com/beber89/miner-dapp-wasm/blockchain"
)

func TestTransactionMessage(t *testing.T) {
	trMsg1 := blockchain.TransactionMessage{
		blockchain.Transaction{"Alice", "Bob", 40},
		1010}
	t.Log(trMsg1.Serialize())
	trMsg2 := blockchain.TransactionMessage{}
	trMsg2.Deserialize(trMsg1.Serialize())
	t.Log(trMsg2)
}

func TestBlockchainScenario(t *testing.T) {
	chn := blockchain.NewBlockchain()
	time.Sleep(1 * time.Second)
	if len(chn.GetChain()) != 1 {
		t.Error("Expected block to be added to chain ")
	} else {
		t.Logf("Genesis Block:\n%x\n", chn.GetChain()[0])
	}

	tr := blockchain.Transaction{"Alice", "Bob", 40}
	chn.RequestTransaction(tr, 1010)
	if len(chn.GetChain()) != 2 {
		t.Error("Expected block to be added to chain ")
	} else {
		t.Logf("New Block:\n%x\n", chn.GetChain()[1])
	}
}
