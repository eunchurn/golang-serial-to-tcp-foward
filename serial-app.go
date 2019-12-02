package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tarm/serial"
)

const (
	CONN_HOST = "0.0.0.0"
	CONN_PORT = "8002"
	CONN_TYPE = "tcp"
)

func main() {
	// port := os.Args[1]
	// baud, err := strconv.Atoi(os.Args[2])
	port := "/dev/ttySAC4"
	baud := 2500000
	c := &serial.Config{Name: port, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go SerialToTCP(s, conn)
	}
}

func SerialToTCP(s *serial.Port, conn net.Conn) {
	buf := make([]byte, 128)
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		conn.Write(buf[:n])
	}
}
