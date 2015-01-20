package main

import (
	"fmt"
	"strconv"
)

type DBUp struct {
	Megs int64
}

func newDBUp(p Packet) (Action, error) {
	d := DBUp{}

	if p.Cmd != "dbup" {
		return d, fmt.Errorf("wrong command for type MemUp")
	}

	i, err := strconv.Atoi(p.Arg)
	if err != nil {
		return d, err
	}

	if i < 1 || i > 1000 {
		return d, fmt.Errorf("dbup megabyte argument must be within range [1, 1000]")
	}

	d.Megs = int64(i)

	return d, nil
}

func (d DBUp) act(dng *danger) error {
	return dng.db.resize(d.Megs)
}
