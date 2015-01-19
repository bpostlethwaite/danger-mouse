package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"./packet"
)

type tcpserver struct {
	sim *simulacra
}

func newTcpServer(sim *simulacra) *tcpserver {
	return &tcpserver{sim: sim}
}

func (s *tcpserver) up() {

	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost: "+s.sim.tcpPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	fmt.Println("TCP server listening on port: " + s.sim.tcpPort)

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

		s.sim.Cmd <- p
		resp := <-s.sim.Res

		_, err = conn.Write(resp.ToBytePack())
		if err != nil {
			panic(err.Error())
		}

		break // we only read one command from each connection
	}

}
