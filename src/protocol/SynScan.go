package protocol

import (
	"golang.org/x/net/ipv4"
	"net"
	"syscall"
)

func SynScan() {
	raw_socket, _ := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_TCP)
	addr := syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}
	sysHeader := ipv4.Header{
		Version:  4,
		Len:      20,
		TotalLen: 20, //20 bytes for ip
		TTL:      64,
		Protocol: 6, //Tcp
		Dst:      net.IPv4(127, 0, 0, 1),
		Src:      net.IPv4(192, 168, 68, 1),
	}
	payload, _ := sysHeader.Marshal()
	syscall.Sendto(raw_socket, payload, 0, &addr)

}
