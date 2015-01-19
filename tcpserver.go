package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"./packet"
)

type tcpserver struct {
	dng *danger
}

func newTcpServer(dng *danger) *tcpserver {
	return &tcpserver{dng: dng}
}

func (s *tcpserver) up() {

	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost: "+s.dng.tcpPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	fmt.Println("TCP server listening on port: " + s.dng.tcpPort)

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		go s.handleRequest(conn)
	}
}

// Handles incoming requests.
func (s *tcpserver) handleRequest(conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {

		p, err := packet.FromBytes(scanner.Bytes())
		if err != nil {
			fmt.Println("Error parsing client message", err.Error())
			break
		}

		s.dng.Cmd <- p
		resp := <-s.dng.Res

		_, err = conn.Write(resp.ToBytePack())
		if err != nil {
			panic(err.Error())
		}

		break // we only read one command from each connection
	}

}
