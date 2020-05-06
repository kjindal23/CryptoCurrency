package main

import (
	"context"
	"cryptoerver"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()
	// our cryptoerver service
	srv := cryptoerver.NewService()
	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := cryptoerver.Endpoints{
		GetCurrencyEndpoint: cryptoerver.MakeGetCurrencyEndpoint(srv),
		StatusEndpoint:      cryptoerver.MakeStatusEndpoint(srv),
	}

	// HTTP transport
	go func() {
		log.Println("crypto Server is listening on port:", *httpAddr)
		handler := cryptoerver.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
