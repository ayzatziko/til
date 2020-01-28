package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	s := http.Server{Addr: ":8080", Handler: http.DefaultServeMux}
	http.DefaultServeMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	})

	http.DefaultServeMux.HandleFunc("/panic/common", func(w http.ResponseWriter, r *http.Request) {
		panic("common")
	})

	http.DefaultServeMux.HandleFunc("/panic/aborterr", func(w http.ResponseWriter, r *http.Request) {
		panic(http.ErrAbortHandler)
	})

	log.Println("server started at " + s.Addr)
	go s.ListenAndServe()
	<-c
	s.Shutdown(context.Background())
}
