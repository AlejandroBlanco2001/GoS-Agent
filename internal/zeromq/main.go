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

	err = client.Connect("tcp://localhost:5555")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}

	// Send message
	_, err = client.Send(msg, 0)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	// Receive reply
	reply, err := client.Recv(0)
	if err != nil {
		fmt.Println("Error receiving reply:", err)
		return
	}

	fmt.Println("Received:", reply)
}
