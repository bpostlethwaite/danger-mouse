package main

import (
	"fmt"
	"runtime/debug"

	"./packet"
)

type MemDown struct {
}

func newMemDown(p packet.Packet) (Action, error) {
	m := MemDown{}

	if p.Cmd != "memdown" {
		return m, fmt.Errorf("wrong command for type MemDown")
	}

	return m, nil
}

func (m MemDown) act(s *simulacra) {
	s.memdb = make([][]byte, 0)

	// force a garbage collection and attempts to return memory back to OS
	debug.FreeOSMemory()
}
