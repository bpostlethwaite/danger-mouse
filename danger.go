package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type danger struct {
	Cmd      chan Packet
	Res      chan Packet
	db       *DB
	httpPort string
	tcpPort  string
	active   *Action
	ping     int
	memcache *MemCache
}

func NewDanger(conf DangerConfig) *danger {

	// should validate port ranges
	tcpPort := strconv.Itoa(conf.TcpPort)
	httpPort := strconv.Itoa(conf.HttpPort)

	// should to validate path here
	dbfile := conf.DBFile

	dng := &danger{
		Cmd:      make(chan Packet, 0),
		Res:      make(chan Packet, 0),
		tcpPort:  tcpPort,
		httpPort: httpPort,
		active:   nil,
		ping:     200,
	}

	dng.db = &DB{path: dbfile}
	if err := dng.db.create(); err != nil {
		dng.handleErr(err)
	}

	dng.memcache = &MemCache{}
	dng.memcache.create()

	return dng
}

func (dng *danger) Run() {

	tcp := newTcpServer(dng)
	go tcp.up()

	http := newHttpServer(dng)
	go http.up()

	for p := range dng.Cmd {
		dng.Res <- dng.cmdRouter(p)
	}
}

func (dng *danger) cmdRouter(p Packet) Packet {

	a, err := PacketToAction(p)

	if err != nil {
		return ErrorPacket(err.Error())
	}

	if dng.active != nil {
		return ErrorPacket("danger is currently busy")
	}

	// lock danger until handler has run
	dng.active = &a

	go func() {

		// unlock danger once handler returns
		defer func() {
			dng.active = nil
		}()

		if err := a.act(dng); err != nil {
			dng.handleErr(err)
		}
	}()

	// return success packet
	return ResponsePacket(reflect.TypeOf(a).Name())
}

func (dng *danger) getStatus() string {
	if dng.active == nil {
		return "danger is idle"
	}
	cmd := reflect.TypeOf(*dng.active).Name()
	return fmt.Sprintf("danger is running command %s with %+v\n", cmd, *dng.active)
}

func (dng *danger) handleErr(err error) {
	// log the error and continue
	panic(err.Error())
}
