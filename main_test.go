package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"testing"
	"time"
)

const connstring = "localhost:2222"

var echotests = []string{
	"some string!",
	"another stest string man....",
	"🤖,",
	"通话/普通",
	"my pony is over\t\t\t\tthe mooon"}

func TestUDPServer(t *testing.T) {
	go setupUDPServer(connstring)
	time.Sleep(2 * time.Second)
	conn, err := net.Dial("udp4", connstring)
	if err != nil {
		t.Error("Expected to connect to ", connstring)
	}
	for _, tt := range echotests {
		msg := fmt.Sprintf("%s\n", tt)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			t.Error("Sending the datagram failed")
		}
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			t.Error("No data received from server", err)
		}
		if strings.Trim(status, "\n") != tt {
			t.Errorf("expected %s to match %s", strings.Trim(status, "\n"), tt)
		}
	}
}

func TestTCPServer(t *testing.T) {
	go setupTCPServer(connstring)
	conn, err := net.Dial("tcp", connstring)
	if err != nil {
		t.Error("Expected to connect to ", connstring)
	}
	for _, tt := range echotests {
		msg := fmt.Sprintf("%s\n", tt)
		_, err = conn.Write([]byte(msg))
		if err != nil {
			t.Error("Sending the datagram failed")
		}
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			t.Error("No data received from server", err)
		}
		if strings.Trim(status, "\n") != tt {
			t.Errorf("expected %s to match %s", strings.Trim(status, "\n"), tt)
		}
	}
}
