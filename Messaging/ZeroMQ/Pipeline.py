import zmq

def producer():
    context = zmq.Context()
    socket = context.socket(zmq.PUSH)
    socket.bind("tcp://127.0.0.1:12345")

    while True:
        workload = random.randint(1, 100)
        socket.send(pickle.dumps(workload))
        time.sleep(workload/NWORKERS)

def worker(id):
    context = zmq.Context()
    socket = context.socket(zmq.PULL)
    socket.connect("tcp://localhost:12345")

    while True:
        work = pickle.loads(socket.recv())
        time.sleep(work)