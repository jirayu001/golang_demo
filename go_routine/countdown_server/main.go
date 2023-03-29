package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		countingDownHandler(conn)
	}

}
func countingDownHandler(conn net.Conn) {

	defer func() {
		io.WriteString(conn, "Your connection will be closed by server")
		conn.Close()
	}()
	count := 5
	for {
		io.WriteString(conn, "Hello 123\n")
		time.Sleep(time.Second)
		count--
		if count == 0 {
			io.WriteString(conn, "Enter number : ")

			break
		}
	}
}
