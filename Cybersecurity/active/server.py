import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)

server_address = ('localhost', 8080)
sock.bind(server_address)

print("UDP server running on port 8080")
print("Ctrl-C to stop the server")

while True:
    data, address = sock.recvfrom(4096)
    
    if data:
        sent = sock.sendto(data, address)
