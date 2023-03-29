//Chat server
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan string

var chMessage = make(chan string)
var chEnteringClient = make(chan client)
var chLeavingClient = make(chan client)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go boardcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			//ig
			log.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	clientName := conn.RemoteAddr()
	chOutGoingMsg := make(chan string)
	chEnteringClient <- chOutGoingMsg
	fmt.Println("Client connected : ", conn.RemoteAddr())
	go cilentWriter(conn, chOutGoingMsg)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		//fmt.Printf("%s => %s\n", clientName, msg)
		chMessage <- fmt.Sprintf("%s => %s\n", clientName, msg)
	}
	chLeavingClient <- chOutGoingMsg
	fmt.Printf("%s disconnected.\n", clientName)

}

func cilentWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		//io.Copy(conn, msg)
		fmt.Fprintf(conn, msg)
	}

}

func boardcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case m := <-chMessage:
			for k := range clients {
				//clientA <-m
				//clientA <-m
				//clientA <-m
				k <- m
			}
		case cli := <-chEnteringClient:
			clients[cli] = true

		case cli := <-chLeavingClient:
			clients[cli] = false
		}
	}

}
