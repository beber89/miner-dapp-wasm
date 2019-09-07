package chainfabric

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

// Node represents a user in the network
//  toIP is IP address of server to connect to
//  toPort is Port address of server to connect to
//  response is value unassignable by user, only read
type Node struct {
	toIP                   string
	toPort                 uint16
	response               string
	ws                     js.Value
	newTransactionCallback func(string)
}

// GetResponse Gets the response attribute of the Node struct
func (nd Node) GetResponse() uint64 {
	if nnc, err := strconv.Atoi(strings.TrimSpace(nd.response)); err == nil {
		nnc := uint64(nnc)
		nd.response = ""
		return nnc
	}
	nd.response = ""
	return 0
}

// SetNewTransactionCallback ...
func (nd *Node) SetNewTransactionCallback(cb func(string)) {
	nd.newTransactionCallback = cb
}

// SendResponse Sends the response attribute of the Node struct over the net
func (nd Node) SendResponse(nnc uint64) {
	nd.ws.Call("send", js.ValueOf(fmt.Sprintf("%d", nnc)))
}

// SendMessage Sends the response attribute of the Node struct over the net
func (nd Node) SendMessage(msg string) {
	nd.ws.Call("send", js.ValueOf(msg))
}

// ResponseEmpty checks if the response is empty string
func (nd Node) ResponseEmpty() bool {
	return nd.response == ""
}

// NewNode Wrapper for creation of Node struct
func NewNode(toIP string, toPort uint16) Node {
	return Node{toIP, toPort, "", js.Value{}, nil}
}

// Connect connects node to remote server via TCP
func (nd *Node) Connect() bool {
	// connect to this socket
	nd.ws = js.Global().Get("WebSocket").New(fmt.Sprintf("ws://%s:%d/ws", nd.toIP, nd.toPort))

	nd.ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return nil
	}))
	nd.ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		message := args[0].Get("data").String()
		fmt.Println("message received ")
		fmt.Println(message)
		flds := strings.Fields(message)
		if flds[0] == "TRANSACTION" {
			// execute callback to start mining
			nd.newTransactionCallback(flds[1])
		} else {
			// receive nonce from tracker
			nd.response = message
		}
		return nil
	}))
	return true
}
