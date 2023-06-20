import zmq, time
import multiprocessing

def server():
    context = zmq.Context()
    socket = context.socket(zmq.PUB) #creates publisher socket
    socket.bind("tcp://*:12345") #bind socket to address
    while True:
        time.sleep(5)
        t = "TIME" + time.asctime()
        socket.send(t.encode())

def client():
    context = zmq.Context()
    socket = context.socket(zmq.SUB) 
    socket.connect("tcp://localhost:12345")
    socket.setsockopt(zmq.SUBSCRIBE, b"TIME")

    for i in range(5):
        time = socket.recv()
        print(time.decode())
