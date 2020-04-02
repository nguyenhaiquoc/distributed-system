# coding = utf-8
"""
    @author: quochai.kstn
"""


import socket
import selectors
import logging

FORMAT = "%(asctime)-15s: %(message)s"
logging.basicConfig(level=logging.INFO, format=FORMAT)

HOSTNAME = "slowwly.robertomurray.co.uk"
PORT = 80
ADDRESS = (HOSTNAME, PORT)
BUFFER_SIZE = 4096


def nonblocking_request():
    """Send HTTP request to a slow website.
       The slow website will send response back after 10 second!
    """
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    client.connect(ADDRESS)
    client.setblocking(False)

    poller = selectors.DefaultSelector()

    # register client object with poller. We interest in both Write and Read event.
    poller.register(client, selectors.EVENT_READ | selectors.EVENT_WRITE)

    header = b"GET /delay/1000/url/http://www.google.co.uk HTTP/1.1\r\nHost: slowwly.robertomurray.co.uk\r\n\r\n"
    sent_so_far = 0
    while True:
        for fd, event in poller.select():
            if event == selectors.EVENT_WRITE:
                logging.info("send data to remote host")
                sent_length = client.send(header[sent_so_far:])
                sent_so_far = sent_so_far + sent_length

                # Check if all data is sent? Then we are no longer interested in write
                if sent_so_far == len(header):
                    poller.modify(client, selectors.EVENT_READ)

            elif event == selectors.EVENT_READ:
                data_from_network = client.recv(BUFFER_SIZE)
                logging.info("receive data from network")
                logging.info(data_from_network)
                if not data_from_network:
                    logging.info("unregister")
                    poller.unregister(client)


if __name__ == "__main__":
    nonblocking_request()
