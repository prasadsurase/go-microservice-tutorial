package main

import (
	"context"
	"go-microservice-tutorial/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// })

	// http.HandleFunc("/endpoint", func(rw http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(rw, "Some endpoint")
	// })
	l := log.New(os.Stdout, "Main: ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	mux := http.NewServeMux()
	mux.Handle("/hello", hh)
	mux.Handle("/bye", gh)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	l.Println("Received terminate, graceful shutdown", <-sigChan)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
