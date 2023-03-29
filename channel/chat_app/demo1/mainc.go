//Chat Client
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	//time.Sleep(10*t)
	io.Copy(conn, os.Stdin)
}
