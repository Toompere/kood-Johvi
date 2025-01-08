import argparse
import re
import socket

def check_upd_ports(host, start, end):
    for port in range(start, end+1):
        sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
        sock.settimeout(2)
        try:
            sock.sendto(b'test', (host, port))
            sock.recvfrom(1024)
        except (socket.timeout, socket.error):
            print(f'Port {port} is closed') 
        else:
            print(f'Port {port} is open')
        finally:
            sock.close()

def check_tcp_ports(host, start, end):
    for port in range(start, end+1):
        sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        sock.settimeout(2)
        try:
            sock.connect((host, port))
        except (socket.timeout, socket.error):
            print(f'Port {port} is closed') 
        else:
            print(f'Port {port} is open') 
        finally:
            sock.close()

def main():
    parser = argparse.ArgumentParser(usage="tinyscanner [OPTIONS] [HOST] [PORT]", add_help=False)

    parser.add_argument('-p', type=str, required=True, help='Range of ports to scan')
    parser.add_argument('-u', type=str, help='UDP scan')
    parser.add_argument('-t', type=str, help='TCP scan')
    parser.add_argument('--help', action='help', help='Show this message and exit.')

    args = parser.parse_args()

    if not args.u and not args.t:
        print('Specify either UDP(-u) or TCP(-t)')
        return

    if args.p.isdigit():
        start, end = int(args.p), int(args.p)
    else:
        match = re.match(r'^(\d+)-(\d+)$', args.p)
        if match:
            start, end = int(match.group(1)), int(match.group(2))
            if start > end:
                print('Invalid range')
                return
        else:
            print('Invalid port number(s)')
            return

    if args.u:
        check_upd_ports(args.u, start, end)
    if args.t:
        check_tcp_ports(args.t, start, end)
    
if __name__ == "__main__":
    main()