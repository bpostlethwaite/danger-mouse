package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type tcpclient struct {
	port string
}

func NewTCPClient(conf DangerConfig) *tcpclient {
	// checking and validation
	port := strconv.Itoa(conf.TcpPort)

	return &tcpclient{
		port: port,
	}
}

func (tcp *tcpclient) Send(p Packet) {

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
		rp, err := FromBytes(scanner.Bytes())
		if err != nil {
			fmt.Println("Error parsing client message", err.Error())
			os.Exit(1)
		}
		if rp.Cmd == "error" {
			println(rp.Arg)
		}
		break // one response at a time
	}

}

func (tcp *tcpclient) Run(args []string) {

	if len(args) > 0 {

		p := Packet{}
		p.Cmd = args[0]

		if len(args) == 2 {
			p.Arg = args[1]
		}

		tcp.Send(p)

	} else {
		tcp.PrintUsage("")
	}

}

func (tcp *tcpclient) PrintUsage(errStr string) {
	fmt.Println("ARGS!!!!!!!!!!!: " + errStr)
}
