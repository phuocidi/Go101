package main

import (
	"log"
	"net"
)

// run netcat in terminal
// nc -lk 1902

func main() {
	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f) // send log mesg to the network connection

	logger.Println("This is a regular mesage")
	logger.Panicln("This is a panic")
}
