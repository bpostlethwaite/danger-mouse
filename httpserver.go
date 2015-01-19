package main

import (
	"fmt"
	"log"
	"net/http"
)

type httpserver struct {
	sim *simulacra
}

func newHttpServer(sim *simulacra) *httpserver {
	return &httpserver{sim: sim}
}

func (s *httpserver) mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, s.sim.getStatus())
}

func (s *httpserver) up() {

	http.HandleFunc("/", s.mainHandler)

	fmt.Println("HTTP servxxer listening on port: " + s.sim.httpPort)

	err := http.ListenAndServe(":"+s.sim.httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
