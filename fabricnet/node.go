package fabricnet

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

// Node represents a user in the network
//  toIP is IP address of server to connect to
//  toPort is Port address of server to connect to
//  response is value unassignable by user, only read
type Node struct {
	toIP       string
	toPort     uint16
	response   string
	connection net.Conn
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

// SendResponse Sends the response attribute of the Node struct over the net
func (nd Node) SendResponse(nnc uint64) {
	fmt.Printf("sending to net nnc is %d\n", nnc)
	fmt.Fprintf(nd.connection, fmt.Sprintf("%d", nnc))
}

// ResponseEmpty checks if the response is empty string
func (nd Node) ResponseEmpty() bool {
	return nd.response == ""
}

// NewNode Wrapper for creation of Node struct
func NewNode(toIP string, toPort uint16) Node {
	return Node{toIP, toPort, "", nil}
}

// Connect connects node to remote server via TCP
func (nd *Node) Connect() {
	// connect to this socket
	nd.connection, _ = net.Dial("tcp", fmt.Sprintf("%s:%d", nd.toIP, nd.toPort))
	fmt.Fprintf(nd.connection, "Connect\n")
	for {
		// listen for reply
		fmt.Println("Shall read from server: ")
		message, _ := bufio.NewReader(nd.connection).ReadString('\n')
		fmt.Println("Did read from server: " + message)
		nd.response = message
		fmt.Println("Message from server: " + message)
		time.Sleep(1 * time.Second)
	}
}
