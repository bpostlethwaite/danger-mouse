package main

import (
	"fmt"
	"net/http"
	"os"
)

type httpserver struct {
	dng *danger
}

func newHttpServer(dng *danger) *httpserver {
	return &httpserver{dng: dng}
}

func (s *httpserver) mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, s.dng.getStatus())
}

func (s *httpserver) pingHandler(w http.ResponseWriter, r *http.Request) {
	if s.dng.ping == 200 {
		fmt.Fprintf(w, "pong")
	} else {
		w.WriteHeader(s.dng.ping)
		fmt.Fprintf(w, "")
	}
}

func (s *httpserver) up() {

	http.HandleFunc("/", s.mainHandler)
	http.HandleFunc("/ping", s.pingHandler)

	err := http.ListenAndServe(":"+s.dng.httpPort, nil)
	if err != nil {
		// log here
		os.Exit(1)
	}
}
