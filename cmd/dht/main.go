package main

import (
	"DHT/internal/service"
	"flag"
	"log"
)

func main() {
	var configurationFile string
	flag.StringVar(&configurationFile, "c", "", "configuration file path")
	flag.Parse()

	if len(configurationFile) == 0 {
		log.Fatal("no argument configurationFile!")
	}
	server := service.NewServer(configurationFile)
	server.Serve()
}
