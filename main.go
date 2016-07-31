package main

import (
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	addr_parts := strings.Split(addr, ":")
	ip := ""
	if len(addr_parts) > 0 {
		ip = addr_parts[0]
		conn.Write([]byte(ip))
	}
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":6666")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
		}
		go handleConnection(conn)
	}
}
