package web

import (
	"fmt"
	"syscall/js"

	"github.com/beber89/miner-dapp-wasm/wallet"

	"github.com/beber89/miner-dapp-wasm/blockchain"
)

type BlockShape struct {
	cells []js.Value
}

type ChainShape struct {
	blocks []BlockShape
}

func (v *ChainShape) NewBlock(blk blockchain.Block) {
	shape := BlockShape{make([]js.Value, 0)}
	tr := blk.GetTransaction()
	shape.NewCellContent(fmt.Sprintf("%s  ->  %s: %0.2f", tr.From, tr.To, tr.Amount))

	shape.NewCellContent(fmt.Sprintf("%x", blk.GetLastBlockHash()))

	shape.NewCellContent(fmt.Sprintf("%d", blk.GetNonce()))

	shape.NewCellContent(fmt.Sprintf("%x", blk.GetHash()))
	v.blocks = append(v.blocks, shape)
}

func (v ChainShape) Draw() {
	document := js.Global().Get("document")
	if chn := document.Call("getElementById", "chain"); chn != js.Null() {
		document.Get("body").Call("removeChild", chn)
	}
	div := document.Call("createElement", "div")
	div.Set("id", js.ValueOf("chain"))
	for _, b := range v.blocks {
		div.Call("appendChild", b.Draw())
	}
	document.Get("body").Call("appendChild", div)
}

func (v *BlockShape) NewCellContent(content string) {
	document := js.Global().Get("document")
	cell := document.Call("createElement", "div")
	cell.Set("className", js.ValueOf("cell color-lightblue"))
	cell.Set("innerHTML", content)
	v.cells = append(v.cells, cell)
}

func (v BlockShape) Draw() js.Value {
	document := js.Global().Get("document")
	div := document.Call("createElement", "div")
	div.Set("className", js.ValueOf("block"))
	for _, c := range v.cells {
		div.Call("appendChild", c)
	}
	return div
}

// ViewChain ...
func ViewChain(wlt wallet.Wallet) {
	shape := ChainShape{make([]BlockShape, 0)}
	for _, b := range wlt.GetBlockchain().GetChain() {
		shape.NewBlock(b)
	}
	shape.Draw()
}

// ViewBlock ...
func ViewBlock(blk blockchain.Block) {
	shape := BlockShape{make([]js.Value, 0)}
	tr := blk.GetTransaction()
	shape.NewCellContent(fmt.Sprintf("%s  ->  %s: %0.2f", tr.From, tr.To, tr.Amount))

	shape.NewCellContent(fmt.Sprintf("%x", blk.GetLastBlockHash()))

	shape.NewCellContent(fmt.Sprintf("%d", blk.GetNonce()))

	shape.NewCellContent(fmt.Sprintf("%x", blk.GetHash()))

	shape.Draw()
}
