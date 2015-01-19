package main

import (
	"fmt"

	"./packet"
)

type Action interface {
	act(s *simulacra)
}

func PacketToAction(p packet.Packet) (Action, error) {
	var act Action
	var err error

	switch p.Cmd {
	case "memup":
		act, err = newMemUp(p)
	case "memdown":
		act, err = newMemDown(p)
	case "ping":
		act, err = newPing(p)
	default:
		err = fmt.Errorf("Unrecognized command")
	}

	return act, err
}
