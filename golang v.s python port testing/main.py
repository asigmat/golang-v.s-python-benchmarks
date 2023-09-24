import threading
import time
import socket


class scan:
    def scan(self, ports, target):
        self.l = []
        self.loop = False

        def scan_port(port):
            _ = time.time()
            if port == ports:
                self.loop = True
            try:
                s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
                s.connect((target, port))
                self.l.append(port)
                s.close()
            except:
                pass

        for i in range(ports + 1):
            t = threading.Thread(target=scan_port, args=(i,))
            t.daemon = True
            t.start()
        while 1:
            if self.loop:
                time.sleep(1)
                return self.l
            else:
                time.sleep(0.1)


print(scan().scan(5000, "youtube.com"))