package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

func main(){
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:5003", conf)
	if err != nil {
		fmt.Println("Client Dial Server error", err)
		return
	}
	defer conn.Close()
	for {
		_, err = conn.Write([]byte("hello\n"))
		if err != nil {
			fmt.Println("Client Write to Server error", err)
			return
		}
		fmt.Println("Client Write to Server Success!")
		time.Sleep(time.Second)
	}


}
