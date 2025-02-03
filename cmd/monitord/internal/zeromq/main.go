package zero

import (
	"alejandroblanco2001/scanneros/internal/platform/logger"
	"encoding/json"

	"alejandroblanco2001/scanneros/internal/terminal"

	zmq "github.com/pebbe/zmq4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Reciever struct {
	socket   *zmq.Socket
	logger   *logger.EchoHandler
	terminal *terminal.Terminal
}

func NewReciever(logger *logger.EchoHandler, terminal *terminal.Terminal) *Reciever {
	socket, _ := zmq.NewSocket(zmq.REP)
	socket.Connect("tcp://localhost:5555")

	return &Reciever{
		socket:   socket,
		logger:   logger,
		terminal: terminal,
	}
}

// Remember we are using REQ/REP sockets, so we need to send a message after receiving one.
// if not, the receiving socket will have an state that wil unable him to receive more messages.
func (r *Reciever) Recieve() {
	defer r.socket.Close()

	for {
		if msg, err := r.socket.Recv(0); err != nil {
			r.logger.Log("Error receiving message: " + err.Error())
			panic(err)
		} else {
			r.logger.Log("Received message: " + msg)
			if msg == "Adapters" {
				r.logger.Log("Sending adapters...")
				response := r.terminal.GetOpenConnectionStatistics()
				responseBinary, err := json.Marshal(response)

				if err != nil {
					r.logger.Log("Error marshalling response: " + err.Error())
					continue
				}

				r.socket.Send(string(responseBinary), 0)
			} else {
				r.logger.Log("Unknown command")
				r.socket.Send("Unknown command", 0)
			}
		}
	}
}

var Module = fx.Options(
	fx.Provide(NewReciever, logger.NewEchoHandler, zap.NewExample, terminal.NewTerminal),
	fx.Invoke(func(r *Reciever) {
		go r.Recieve()
	}),
)
