package main

import (
	"fmt"
	"reflect"
	"strconv"

	"./packet"
)

type SimulacraConfig struct {
	httpPort int
	tcpPort  int
}

type simulacra struct {
	Cmd      chan packet.Packet
	Res      chan packet.Packet
	httpPort string
	tcpPort  string
	active   *Action
	ping     bool
	memdb    [][]byte
}

func NewSimulacra(conf SimulacraConfig) *simulacra {

	// checking and validation
	tcpPort := strconv.Itoa(conf.tcpPort)
	httpPort := strconv.Itoa(conf.httpPort)

	return &simulacra{
		Cmd:      make(chan packet.Packet, 0),
		Res:      make(chan packet.Packet, 0),
		tcpPort:  tcpPort,
		httpPort: httpPort,
		active:   nil,
		memdb:    make([][]byte, 0),
	}
}

func (s *simulacra) Run() {

	tcp := newTcpServer(s)
	go tcp.up()

	http := newHttpServer(s)
	go http.up()

	for p := range s.Cmd {
		s.Res <- s.cmdRouter(p)
	}
}

func (s *simulacra) cmdRouter(p packet.Packet) packet.Packet {

	a, err := PacketToAction(p)

	if err != nil {
		return packet.ErrorPacket(err.Error())
	}

	if s.active != nil {
		return packet.ErrorPacket("simulacra is currently busy")
	}

	// lock simulacra until handler has run
	s.active = &a

	go func() {

		// unlock simulacra once handler returns
		defer func() {
			s.active = nil
		}()

		a.act(s)
	}()

	// return success packet
	return packet.ResponsePacket(reflect.TypeOf(a).Name())
}

func (s *simulacra) getStatus() string {
	if s.active == nil {
		return "simulacra is idle"
	}
	return fmt.Sprintf("simulacra is running command", *s.active)
}
