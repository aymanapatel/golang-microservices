package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/AymanArif/golang-microservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create handlers
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)
	sm.Handle("/products", ph)
	//http.ListenAndServe(":9090", sm)

	// create a new server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	// go func(){...}() is non-blocking function call
	// Start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)      // signal channel
	signal.Notify(sigChan, os.Interrupt) // Broadcast message when OS.Signal gives os.Interrupt
	signal.Notify(sigChan, os.Kill)      // Broadcast message when OS.Signal gives os.Kill

	sig := <-sigChan // Block channel till message is available to be consumed

	l.Println("Exit started \n----------------Exiting------------------\n")
	l.Println("\n Recieved terminate, graceful shutdown \n ", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
