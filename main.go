package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	version string = "0.1"
	address string = "127.0.0.1"
	port    string = "2222"
)

func print_help() {
	fmt.Printf("go-echo version: %s\n", version)
	fmt.Println("usage: go-echo [-h] [-H HOST_NAME] [-p PORT]")
	os.Exit(0)
}

// Handles connection, returns what was sent to a socket and status via a channel
func handleConnection(conn *net.TCPConn, c chan string) {
	log.Printf("New connection from %s", conn.RemoteAddr().String())
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		echoMessage, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			if err == io.EOF {
				log.Printf("Client closed the connection %s", conn.RemoteAddr().String())
			} else {
				log.Printf("Some problem with reading from client %s", conn.RemoteAddr().String())
			}
			c <- fmt.Sprintf("done serving %s", conn.RemoteAddr().String())
			return
		}
		_, err = conn.Write([]byte(echoMessage))
	}
	c <- fmt.Sprintf("done serving %s", conn.RemoteAddr().String())
	return
}

func main() {

	addressFlag := flag.String("H", address, "address to listen on default: localhost")
	portFlag := flag.String("p", port, "port to listen on, default: 2055")
	helpFlag := flag.Bool("h", false, "help message")
	flag.Parse()

	if *helpFlag != false {
		print_help()
	}

	connString := *addressFlag + ":" + *portFlag
	listenAddress, err := net.ResolveTCPAddr("tcp4", connString)

	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", listenAddress)
	defer ln.Close()

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Listening on ", connString)
	c := make(chan string)
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn, c)
		log.Print(<-c)
	}
}
