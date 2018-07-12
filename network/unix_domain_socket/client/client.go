package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	c, err := net.Dial("unix", "/tmp/go.sock")
	if err != nil {
		log.Fatal("Dail error", err)
	}
	defer c.Close()

	go reader(c)
	for {
		msg := "hi"
		_, err := c.Write([]byte(msg))
		if err != nil {
			log.Fatal("Write error", err)
			break
		}
		fmt.Println("CLent send ", msg)
		time.Sleep(1e9)
	}
}

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		_, err := r.Read(buf)
		if err != nil {
			return
		}
		fmt.Println("Client got:", string(buf))
	}
}
