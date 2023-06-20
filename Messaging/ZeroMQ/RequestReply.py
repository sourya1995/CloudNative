import zmq

def server():
    context = zmq.Context()
    socket = context.socket(zmq.REP)
    socket.bind("tcp://*:12345")

    while True:
        message = socket.recv()
        if not "STOP" in str(message):
            reply = str(message.decode())+'*'
            socket.send(reply.encode())
        else:
            break

def client():
    context = zmq.Context()
    socket = context.socket(zmq.REQ)
    socket.connect("tcp://localhost:12345")
    socket.send(b"Hello, world!")
    message = socket.recv()
    socket.send(b"STOP")
    print(message.decode())