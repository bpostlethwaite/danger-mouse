package main

import (
	"fmt"
)

type MemDown struct {
}

func newMemDown(p Packet) (Action, error) {
	m := MemDown{}

	if p.Cmd != "memdown" {
		return m, fmt.Errorf("wrong command for type MemDown")
	}

	return m, nil
}

func (m MemDown) act(dng *danger) error {
	dng.memcache.create()
	return nil
}
