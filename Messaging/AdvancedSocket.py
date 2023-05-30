import socket
import pickle

class DBClient:
    def _sendrecv(self, message):
        sock = socket()
        sock.connect((self.host, self.port))
        sock.send(pickle.dumps(message))
        result = pickle.loads(sock.recv(1024))
        sock.close()
        return result
    
    def create(self):
        self.listID = self._sendrecv([CREATE])
        return self.listID
    
    def getValue(self):
        return self._sendrecv([GETVALUE, self.listID])
    
    def appendData(self, data):
        return self._sendrecv([APPEND, data, self.listID])
    

class Server:
    self.setOfLists = {}

    def run(self):
        while True:
            (conn, addr) = self.sock.accept()
            data = conn.recv(1024)
            request = pickle.loads(data)

            if request[0] == CREATE:
                listID = len(self.setOfLists) + 1
                self.setOfLists[listID] = []
                conn.send(pickle.dumps(listID))

            elif request[0] == APPEND:
                listID = request[2]
                data = request[1]
                self.setOfLists[listID].append(data)
                conn.send(pickle.dumps(OK))

            elif request[0] == GETVALUE:
                listID = request[1]
                result = self.setOfLists[listID]
                conn.send(pickle.dumps(result))


class Client:
    def __init__(self, port):
        self.host = 'localhost'
        self.port = port
        self.sock = socket()
        self.sock.bind(self.host, self.port)
        self.sock.listen(2)

    def sendTo(self, host, port, data):
        sock = socket()
        sock.connect((host, port))
        sock.send(pickle.dumps(data))
        sock.close()

    def recvAny(self):
        (conn, addr) = self.sock.accept()
        return conn.recv(1024)
    

def client1():
    c1 = Client(PORTC1)
    dbC1 = DBClient(HOSTS, PORTS)
    dbC1.create()
    dbC1.appendData('Client 1')
    c1.sendTo(HOSTC2, PORTC2, dbC1)


def client2():
    c1 = Client(PORTC2)
    data = c2.recvAny()
    dbC2 = pickle.loads(data)
    dbC2.appendData('Client 2')