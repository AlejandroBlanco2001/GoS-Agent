package zero

import (
	"alejandroblanco2001/scanneros/internal/platform/logger"

	zmq "github.com/pebbe/zmq4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Reciever struct {
	socket *zmq.Socket
	logger *logger.EchoHandler
}

func NewReciever(logger *logger.EchoHandler) *Reciever {
	socket, _ := zmq.NewSocket(zmq.REP)
	socket.Bind("tcp://*:5555")

	return &Reciever{
		socket: socket,
		logger: logger,
	}
}

func (r *Reciever) Recieve() string {
	r.logger.Log("Recieving messages...")
	msg, _ := r.socket.Recv(0)
	return msg
}

var Module = fx.Options(
	fx.Provide(NewReciever, logger.NewEchoHandler, zap.NewExample),
	fx.Invoke(func(r *Reciever) {
		go r.Recieve()
	}),
)
