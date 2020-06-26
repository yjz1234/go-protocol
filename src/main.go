package main

import (
	"log"
	"protocol/src/protocol"
)

func main() {
	ipmin, ipmax := protocol.ParserNetMask("192.168.68.100/30")
	log.Printf("%s,%s", ipmin, ipmax)
}
