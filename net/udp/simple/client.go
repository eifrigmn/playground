package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "37", "port")
//go run timeclient.go -host time.nist.gov
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		return
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	var msg = "hello"
	for {
		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed:", err)
			return
		}
		data := make([]byte, 1024)
		cnt, err := conn.Read(data)
		if err != nil {
			fmt.Println("failed to read UDP msg because of ", err)
			return
		}
		fmt.Printf("Read From UDP Server %s\n", data[:cnt])
		time.Sleep(1*time.Second)
	}

}
