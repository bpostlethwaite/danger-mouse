package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"

	"../packet/"
)

type SimulacraConfig struct {
	httpPort int
	tcpPort  int
}

type tcpclient struct {
	port string
}

func NewTCPClient(conf SimulacraConfig) *tcpclient {
	// checking and validation
	tcpPort := strconv.Itoa(conf.tcpPort)

	return &tcpclient{
		port: tcpPort,
	}
}

func (tcp *tcpclient) Send(p packet.Packet) {

	servAddr := "localhost:" + tcp.port
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()

	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(p.ToBytePack())
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		rp, err := packet.FromBytes(scanner.Bytes())
		if err != nil {
			fmt.Println("Error parsing client message", err.Error())
			os.Exit(1)
		}

		println(rp.Arg)
		break // one response at a time
	}

}

func (tcp *tcpclient) PrintUsage(errStr string) {
	fmt.Println("ARGS!!!!!!!!!!!: " + errStr)
}
