package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Guardian/proxy/api"
)

func main() {

	l := log.New(os.Stdout, "---> Proxy Service ", log.LstdFlags)

	port := ":" + os.Getenv("PORT_PROXY_SERVICE")
	srv := api.NewServer(port)

	go func() {
		l.Println("Starting server on port " + port)

		err := srv.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	log.Println("Got signal: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)
}
