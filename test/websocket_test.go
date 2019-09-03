package dapp_test

import (
	"testing"

	"github.com/beber89/miner-dapp-wasm/chainfabric"
)

func TestConnection(t *testing.T) {
	nd := chainfabric.NewNode("127.0.0.1", 8081)
	t.Logf("Node is %v", nd)
	nd.Connect()
	// t.Logf("connection success is %t", success)
}
