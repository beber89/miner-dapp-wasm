package web

import (
	"syscall/js"
)

// renderPage reloads elements of the page when change to wallet occurs
func RenderPage() {
	wlt := *GetWallet()
	document := js.Global().Get("document")
	chainView := (*wlt.GetBlockchain()).Draw()
	if chn := document.Call("getElementById", "chain"); chn != js.Null() {
		document.Call("getElementById", "blockchain-placeholder").Call("removeChild", chn)
	}
	document.Call("getElementById", "blockchain-placeholder").Call("appendChild", chainView)
	networthLabel := document.Call("getElementById", "networth")
	networthLabel.Set("innerHTML", wlt.Networth())
}
