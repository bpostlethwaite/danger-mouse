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
	w.WriteHeader(200)
	fmt.Fprintf(w, s.sim.getStatus())
}

func (s *httpserver) pingHandler(w http.ResponseWriter, r *http.Request) {
	if s.sim.ping == 200 {
		fmt.Fprintf(w, "pong")
	} else {
		w.WriteHeader(s.sim.ping)
		fmt.Fprintf(w, "")
	}
}

func (s *httpserver) up() {

	http.HandleFunc("/", s.mainHandler)
	http.HandleFunc("/ping", s.pingHandler)

	fmt.Println("HTTP servxxer listening on port: " + s.sim.httpPort)

	err := http.ListenAndServe(":"+s.sim.httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
