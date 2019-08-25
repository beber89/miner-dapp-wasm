package fabricnet

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

// Node represents a user in the network
//  toIP is IP address of server to connect to
//  toPort is Port address of server to connect to
//  response is value unassignable by user, only read
type Node struct {
	toIP     string
	toPort   uint16
	response string
}

// GetResponse Gets the response attribute of the Node struct
func (nd Node) GetResponse() string {
	return nd.response
}

// NewNode Wrapper for creation of Node struct
func NewNode(toIP string, toPort uint16) Node {
	return Node{toIP, toPort, ""}
}

// Connect connects node to remote server via TCP
func (nd *Node) Connect() {
	// connect to this socket
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", nd.toIP, nd.toPort))
	for {
		// read in input from stdin
		// reader := bufio.NewReader(os.Stdin)
		// fmt.Print("Text to send: ")
		// text, _ := reader.ReadString('\n')
		// // send to socket
		// fmt.Fprintf(conn, text+"\n")
		fmt.Fprintf(conn, "Connect\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		nd.response = message
		fmt.Print("Message from server: " + message)
		time.Sleep(1 * time.Second)
	}
}
