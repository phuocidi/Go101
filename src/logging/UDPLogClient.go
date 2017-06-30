package main

import (
	"log"
	"net"
	"time"
)

// run netcat in terminal
// nc -luk 1902

func main() {
	timeout := 30 * time.Second
	conn, err := net.DialTimeout("udp", "localhost:1902", timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f) // send log mesg to the network connection

	logger.Println("This is a regular mesage")
	logger.Panicln("This is a panic")
}
