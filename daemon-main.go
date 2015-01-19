package main

import "flag"

const MB = 1000000

func main() {

	var port int
	flag.IntVar(&port, "port", 8080, "port for server to listen on")
	flag.Parse()

	conf := DangerConfig{
		httpPort: port,
		tcpPort:  3344,
	}

	dng := NewDanger(conf)

	dng.Run()

}
