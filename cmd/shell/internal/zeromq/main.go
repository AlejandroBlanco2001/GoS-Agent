package zero

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func NewClient(msg string) {
	client, err := zmq.NewSocket(zmq.REQ)

	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}

	defer client.Close()

	err = client.Bind("tcp://*:5555")

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	fmt.Printf("Sending message")

	// Send message
	client.Send(msg, 0)

	// Receive message
	reply, err := client.Recv(0)

	if err != nil {
		fmt.Println("Error receiving message:", err)
		return
	}

	fmt.Println("Received reply:", reply)
}
