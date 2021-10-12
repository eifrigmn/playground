package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
)

func main(){
	cert, err := tls.LoadX509KeyPair("./certs/server.pem", "./certs/server.key")
	if err != nil {
		fmt.Println("Server Load Certs error", err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", "127.0.0.1:5003", config)
	if err != nil {
		fmt.Println("Server Listen Error", err)
		return
	}

	fmt.Println("Server Listen At", ln.Addr().String())
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Server Handle Conn Error", err)
			return
		}
		fmt.Println("Client Connected!", conn.RemoteAddr().String())
		go handleServConn(conn)
	}
}

func handleServConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF{
				fmt.Println("Remote Client Closed!", conn.RemoteAddr().String())
			} else {
				fmt.Println("Server Handle conn error", err)
			}
			return
		}
		fmt.Println("Server Got Message ", string(line))
	}
}
