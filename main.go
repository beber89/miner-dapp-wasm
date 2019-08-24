package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Node struct {
	toIP     string
	toPort   uint16
	response string
}

func NewNode(toIP string, toPort uint16) Node {
	return Node{toIP, toPort, ""}
}

func (self *Node) Connect() {
	var url = fmt.Sprintf("http://%s:%d", self.toIP, self.toPort)
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		self.response = bodyString
		fmt.Println("[Node connect] Response")
		fmt.Println(self.response)
	} else {
		fmt.Println("status not ok")
		fmt.Println(resp.StatusCode)
	}
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	var bobNode = NewNode("127.0.0.1", 8081)
	bobNode.Connect()

	fmt.Println("[main] Response")
	fmt.Println(bobNode.response)
}
