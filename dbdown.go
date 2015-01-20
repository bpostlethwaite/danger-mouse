package main

import (
	"fmt"
)

type DBDown struct {
}

func newDBDown(p Packet) (Action, error) {
	d := DBDown{}

	if p.Cmd != "dbdown" {
		return d, fmt.Errorf("wrong command for type MemUp")
	}

	return d, nil
}

func (d DBDown) act(dng *danger) error {
	return dng.db.create()
}
