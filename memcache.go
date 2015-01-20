package main

import (
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"
)

type MemCache struct {
	cache [][]byte
}

func (mem *MemCache) create() {

	mem.cache = make([][]byte, 0)

	// force garbage collection and attempt to return memory back to OS
	debug.FreeOSMemory()
}

func (mem *MemCache) resize(mb int64) {

	nbyte := mb * bytesPerMB

	moremem := func() bool {
		if mem.size() < nbyte {
			return true
		}
		return false
	}

	for moremem() {
		rem := nbyte - mem.size()

		if rem < 1 {
			return
		}

		junk := make([]byte, rem)

		for i := 0; i < len(junk); i++ {
			junk[i] = uint8(rand.Intn(255))
		}

		mem.cache = append(mem.cache, junk)

		time.Sleep(250 * time.Millisecond)
	}

}

func (mem *MemCache) size() int64 {
	stat := runtime.MemStats{}
	runtime.ReadMemStats(&stat)

	return int64(stat.HeapAlloc)
}
