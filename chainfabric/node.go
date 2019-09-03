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
	fmt.Println("is not an integer.")
	return 0
}

// SetNewTransactionCallback ...
func (nd *Node) SetNewTransactionCallback(cb func(string)) {
	nd.newTransactionCallback = cb
}

// SendResponse Sends the response attribute of the Node struct over the net
func (nd Node) SendResponse(nnc uint64) {
	fmt.Printf("sending to net nnc is %d\n", nnc)
	nd.ws.Call("send", js.ValueOf(fmt.Sprintf("%d", nnc)))
}

// SendMessage Sends the response attribute of the Node struct over the net
func (nd Node) SendMessage(msg string) {
	fmt.Printf("sending to net msg is %s\n", msg)
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
	// conn, err := websocket.Dial(fmt.Sprintf("ws://%s:%d", nd.toIP, nd.toPort))
	nd.ws = js.Global().Get("WebSocket").New(fmt.Sprintf("ws://%s:%d/ws", nd.toIP, nd.toPort))
	// conn, err := websocket.Dial("ws://127.0.0.1:8081/ws")
	// if err != nil {
	// 	fmt.Println("[Node Connect] Could not connect to tracker")
	// 	fmt.Println(err)
	// 	return false
	// }
	nd.ws.Call("addEventListener", "open", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("open")
		// ws.Call("send", js.ValueOf("1232"))
		// ws.Call("send", js.TypedArrayOf([]byte{123}))
		return nil
	}))
	nd.ws.Call("addEventListener", "message", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("message")
		message := args[0].Get("data").String()
		flds := strings.Fields(message)
		if flds[0] == "TRANSACTION" {
			// execute callback to start mining
			fmt.Println("message recieved is transaction")
			fmt.Printf("message is: %s\n", message)
			nd.newTransactionCallback(flds[1])
		} else {
			// send nonce to tracker
			fmt.Println("message recieved is a solved nonce")
			fmt.Printf("message is: %s\n", message)
			nd.response = message
		}
		return nil
	}))
	// nd.connection = conn
	fmt.Println("Connected to tracker !")
	// fmt.Fprintf(nd.connection, "Connect\n")
	return true
}

// StartListening waits for data from tracker
// func (nd *Node) StartListening() {
// 	for {
// 		// listen for reply
// 		fmt.Println("Shall read from server: ")
// 		buf := make([]byte, 1024)
// 		n, err := nd.connection.Read(buf) // Blocks until a WebSocket frame is received.
// 		message := strings.TrimSuffix(string(buf[:n]), "\n")
// 		if err != nil {
// 			panic("Could not read message from websocket")
// 		}
// 		fmt.Println("Did read from server: " + message)
// 		flds := strings.Fields(message)
// 		if flds[0] == "TRANSACTION" {
// 			// execute callback to start mining
// 			nd.newTransactionCallback(flds[1])
// 		} else {
// 			// send nonce to tracker
// 			nd.response = message
// 		}
// 		fmt.Println("Message from server: " + message)
// 		time.Sleep(1 * time.Second)
// 	}
// }
