package main

import (
	"rise-nostr/pkg/app/server"
	"rise-nostr/pkg/config"
)

const confPath = "./config.json"

func main() {

	//cfg, err := config.New(confPath)
	cfg, _ := config.New(confPath)

	svr := server.New(cfg)
	svr.Serve()

}
