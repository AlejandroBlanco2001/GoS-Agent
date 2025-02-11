import zmq

# Create a ZeroMQ context
context = zmq.Context()

# Create a REP socket and bind to the same port where monitord is listening
socket = context.socket(zmq.REP)
socket.connect("tcp://localhost:5555")

print("Intercepting response from shell...")

while True:
    # Receive a message from the REP socket (this will be the response from monitord)
    message = socket.recv_string()
    print("Intercepted response:", message)

    # You can process or log the response before returning it
    socket.send_string("")
