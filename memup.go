package main

import (
	"fmt"
	"strconv"
)

type MemUp struct {
	Megs int64
}

func newMemUp(p Packet) (Action, error) {
	m := MemUp{}

	if p.Cmd != "memup" {
		return m, fmt.Errorf("wrong command for type MemUp")
	}

	i, err := strconv.Atoi(p.Arg)
	if err != nil {
		return m, err
	}

	if i < 1 || i > 1000 {
		return m, fmt.Errorf("memup megabyte argument must be within range [1, 1000]")
	}

	m.Megs = int64(i)

	return m, nil
}

func (m MemUp) act(dng *danger) error {
	dng.memcache.resize(m.Megs)
	return nil
}
