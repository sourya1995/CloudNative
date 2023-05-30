from socket import *

class Server:
    def run(self):
        s = socket(AF_INET, SOCK_STREAM)
        s.bind((HOST, PORT))
        s.listen(1)
        (conn, addr) = s.accept()
        while True:
            data = conn.recv(1024)
            if not data: break
            conn.send(data+b"*")
        conn.close()

class Client:
    def run(self):
        s = socket(AF_INET, SOCK_STREAM)
        s.connect((HOST, PORT))
        s.send(b"hello, world!")
        data = s.recv(1024)
        print(data)
        s.send(b" ")
        s.close()