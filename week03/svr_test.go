package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
	mhttp "week03/http"
	"week03/transport"
)

func TestManualStopApp(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello, GohperCon")
	})

	addrs := []string{
		":8080",
		":8081",
		":8082",
	}

	svrs := make([]transport.Server, len(addrs))
	for i, addr := range addrs {
		svrs[i] = mhttp.NewServer(addr, mux)
	}

	app := NewApp(WithServer(svrs...))

	time.AfterFunc(time.Second*5, func() {
		_ = app.Stop()
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func TestKillStopApp(t *testing.T) {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Hello, GohperCon")
	})

	addrs := []string{
		":8080",
		":8081",
		":8082",
	}

	svrs := make([]transport.Server, len(addrs))
	for i, addr := range addrs {
		svrs[i] = mhttp.NewServer(addr, mux)
	}

	app := NewApp(WithServer(svrs...))

	time.AfterFunc(time.Second*5, func() {
		app.SimulateKillSignal()
	})

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
