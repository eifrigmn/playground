package main

import (
	"github.com/eifrigmn/playground/net/tcp/config"
	"github.com/eifrigmn/playground/net/tcp/server"
)

func main(){
	cfg := config.LoadConfig()
	s := server.NewServerV2("S2", cfg)
	s.Serve()
}
