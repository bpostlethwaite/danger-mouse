package main

import (
	"fmt"
)

type Action interface {
	act(dng *danger) error
}

func PacketToAction(p Packet) (Action, error) {
	var act Action
	var err error

	switch p.Cmd {
	case "memup":
		act, err = newMemUp(p)
	case "memdown":
		act, err = newMemDown(p)
	case "ping":
		act, err = newPing(p)
	case "cpu":
		act, err = newCpu(p)
	case "dbup":
		act, err = newDBUp(p)
	case "dbdown":
		act, err = newDBDown(p)

	default:
		err = fmt.Errorf("Unrecognized command")
	}

	return act, err
}
