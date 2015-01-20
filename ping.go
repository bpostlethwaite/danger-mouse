package main

import (
	"fmt"
	"strconv"
)

type Ping struct {
	Code int
}

func newPing(p Packet) (Action, error) {
	m := Ping{}

	if p.Cmd != "ping" {
		return m, fmt.Errorf("wrong command for type MemUp")
	}

	i, err := strconv.Atoi(p.Arg)
	if err != nil {
		return m, err
	}

	if i < 1 || i > 999 {
		return m, fmt.Errorf("ping status code argument must be within range (0, 1000)")
	}

	m.Code = i

	return m, nil
}

func (m Ping) act(dng *danger) error {
	dng.ping = m.Code
	return nil
}
