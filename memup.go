package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"

	"./packet"
)

type MemUp struct {
	Megs int
}

func newMemUp(p packet.Packet) (Action, error) {
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

	m.Megs = i

	return m, nil
}

func (m MemUp) act(dng *danger) {

	megs := float64(m.Megs)

	getheap := func() float64 {
		mem := runtime.MemStats{}
		runtime.ReadMemStats(&mem)
		return float64(mem.HeapAlloc) / MB
	}

	moremem := func() bool {
		if getheap() < megs {
			return true
		}
		return false
	}

	for moremem() {
		rem := int(getheap() - megs)
		if rem < 1 {
			rem = 1
		}

		junk := make([]byte, rem*MB)

		for i := 0; i < len(junk); i++ {
			junk[i] = uint8(rand.Intn(255))
		}

		dng.memdb = append(dng.memdb, junk)

		fmt.Println("memory now at", getheap(), "megabytes")

		time.Sleep(250 * time.Millisecond)
	}

}
