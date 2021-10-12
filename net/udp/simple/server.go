package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"time"
)
var ServerPort = flag.String("port", "5002", "port")
func main(){
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s","localhost", *ServerPort))
	if err != nil {
		fmt.Println("Resolve UDP Addr error", err)
		return
	}

	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Listen UDP error", err)
		return
	}
	defer listener.Close()
	fmt.Println("UDP Server Started", udpAddr.String())
	for {
		data := make([]byte, 1024)
		cnt, addr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Println("Read from UDP Error", err)
			break
		}
		fmt.Printf("Read from UDP %s Data: %s\n", addr.String(), data[:cnt])
		buf := bytes.NewBuffer([]byte{})
		buf.WriteString(time.Now().Format("20060102150405"))
		buf.WriteString("->")
		buf.Write(data[:cnt])
		_, err = listener.WriteToUDP(buf.Bytes(), addr)
		if err != nil {
			fmt.Println("Write Back to UDP Client error", err)
			break
		}
		fmt.Println("Write Back UDP Success!")
	}
}
