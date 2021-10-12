package main

import (
	"fmt"
	"net"
	"time"
)

func main(){
	fmt.Println("client start...")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:5001")
	if err != nil {
		fmt.Println("Client start error", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello..."))
		if err != nil {
			fmt.Println("Client write error", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Client read buff error", err)
			return
		}
		fmt.Printf("server call back buf = %s, cnt = %d\n", buf, cnt)
		time.Sleep(time.Second)
	}
}

