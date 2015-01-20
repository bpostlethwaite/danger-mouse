package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type tcpserver struct {
	dng *danger
}

func newTcpServer(dng *danger) *tcpserver {
	return &tcpserver{dng: dng}
}

func (s *tcpserver) up() error {

	// Listen for incoming connections.
	l, err := net.Listen("tcp", "localhost: "+s.dng.tcpPort)
	if err != nil {
		// log here
		os.Exit(1)
	}

	// Close the listener when the application closes.
	defer l.Close()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			// log here
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

		p, err := FromBytes(scanner.Bytes())
		if err != nil {
			fmt.Println("Error parsing client message", err.Error())
			break
		}

		s.dng.Cmd <- p
		resp := <-s.dng.Res

		_, err = conn.Write(resp.ToBytePack())
		if err != nil {
			s.dng.handleErr(err)
		}

		break // we only read one command from each connection
	}

}
