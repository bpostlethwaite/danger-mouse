package main

import (
	"flag"

	"../packet/"
)

func main() {

	flag.Parse()

	args := flag.Args()
	p := packet.Packet{}

	tcp := NewTCPClient(SimulacraConfig{
		tcpPort: 3344,
	})

	if len(args) > 0 {

		p.Cmd = args[0]
		if len(args) == 2 {
			p.Arg = args[1]
		}

		tcp.Send(p)

	} else {
		tcp.PrintUsage("")
	}

	return
}
