# coding = utf-8
"""
    @author: quochai.kstn
"""


import socket
import logging

FORMAT = "%(asctime)-15s: %(message)s"
logging.basicConfig(level=logging.INFO, format=FORMAT)

HOSTNAME = "slowwly.robertomurray.co.uk"
PORT = 80
ADDRESS = (HOSTNAME, PORT)
BUFFER_SIZE = 4096


def blocking_request():
    """Send HTTP request to a slow website.
       The slow website will send response back after 10 second!
    """
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect(ADDRESS)
    header = b"GET /delay/10000/url/http://www.google.co.uk HTTP/1.1\r\nHost: slowwly.robertomurray.co.uk\r\n\r\n"
    logging.info("send request")
    client.sendall(header)
    data = ""

    """
        This loop pattern
    """
    while True:
        # client.recv is a blocking call, it will wait until data available in OS buffer.
        # In this example, it should block for > 10 seconds.
        network_data = client.recv(BUFFER_SIZE)
        if not network_data:
            break
        data = data + network_data

    logging.info(data)
    logging.info("receive response")


if __name__ == "__main__":
    blocking_request()
