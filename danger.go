package main

import (
	"fmt"
	"reflect"
	"strconv"

	"./packet"
)

type DangerConfig struct {
	httpPort int
	tcpPort  int
}

type danger struct {
	Cmd      chan packet.Packet
	Res      chan packet.Packet
	httpPort string
	tcpPort  string
	active   *Action
	ping     int
	memdb    [][]byte
}

func NewDanger(conf DangerConfig) *danger {

	// checking and validation
	tcpPort := strconv.Itoa(conf.tcpPort)
	httpPort := strconv.Itoa(conf.httpPort)

	return &danger{
		Cmd:      make(chan packet.Packet, 0),
		Res:      make(chan packet.Packet, 0),
		tcpPort:  tcpPort,
		httpPort: httpPort,
		active:   nil,
		ping:     200,
		memdb:    make([][]byte, 0),
	}
}

func (s *danger) Run() {

	tcp := newTcpServer(s)
	go tcp.up()

	http := newHttpServer(s)
	go http.up()

	for p := range s.Cmd {
		s.Res <- s.cmdRouter(p)
	}
}

func (s *danger) cmdRouter(p packet.Packet) packet.Packet {

	a, err := PacketToAction(p)

	if err != nil {
		return packet.ErrorPacket(err.Error())
	}

	if s.active != nil {
		return packet.ErrorPacket("danger is currently busy")
	}

	// lock danger until handler has run
	s.active = &a

	go func() {

		// unlock danger once handler returns
		defer func() {
			s.active = nil
		}()

		a.act(s)
	}()

	// return success packet
	return packet.ResponsePacket(reflect.TypeOf(a).Name())
}

func (s *danger) getStatus() string {
	if s.active == nil {
		return "danger is idle"
	}
	cmd := reflect.TypeOf(*s.active).Name()
	return fmt.Sprintf("danger is running command %s with %+v\n", cmd, *s.active)
}
