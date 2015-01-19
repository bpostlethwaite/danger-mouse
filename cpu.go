package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"./packet"
)

type Cpu struct {
	Seconds time.Duration // how long to burn CPU
}

func newCpu(p packet.Packet) (Action, error) {
	c := Cpu{}

	if p.Cmd != "cpu" {
		return c, fmt.Errorf("wrong command for type MemUp")
	}

	i, err := strconv.Atoi(p.Arg)
	if err != nil {
		return c, err
	}

	if i < 1 {
		return c, fmt.Errorf("cpu seconds argument must be greater than 0")
	}

	c.Seconds = time.Duration(i) * time.Second

	return c, nil
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func (c Cpu) act(s *simulacra) {

	ncpus := runtime.GOMAXPROCS(-1)
	ngos := 4 * ncpus

	done := make(chan struct{})
	defer close(done)

	burner := func() {
		for {
			select {
			case <-done:
				return
			default:
				for j := 0; j < 100000; j++ {
					for i := range []int{1, 2, 3, 4, 5, 6, 7, 8} {
						fact(i)
					}
				}
			}
		}
	}

	for i := 0; i < ngos; i++ {
		go burner()
	}

	time.Sleep(c.Seconds)
}
