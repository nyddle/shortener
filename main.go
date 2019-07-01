package main

import (
	"fmt"
	"log"
	"os"
	"shortener/httpserver"
	"shortener/myshortener"
	"shortener/service"
	"strconv"
)

func main() {

	httpPort, err := strconv.Atoi(os.Getenv("SHORTENER_PORT"))
	if err != nil {
		panic(fmt.Sprint("SHORTENER_PORT not defined"))
	}

	shortener := myshortener.NewMyShortener()
	shortenerService := service.NewShortenerService(shortener)
	hs := httpserver.NewHTTPServer(shortenerService, httpPort)

	log.Fatal(hs.Start())
}
