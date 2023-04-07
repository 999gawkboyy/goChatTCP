package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
)

func acceptConnections(ln net.Listener, connections chan<- net.Conn) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		color.Red("Connected : %s", conn.RemoteAddr().String())
		connections <- conn
	}
}

func handleConnection(conn net.Conn, name string) {
	defer conn.Close()

	buf := make([]byte, 2048)

	// send msg
	var msg string
	for {
		fmt.Print(">> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg = sc.Text()
		if strings.Trim(msg, " ") == "quit" {
			conn.Close()
			os.Exit(0)
		}
		conn.Write([]byte(name + ">> " + msg))

		// recieve msg
		n, err := conn.Read(buf)
		if err != nil {
			color.Red(err.Error())
			return
		}
		color.Cyan("%s", string(buf[:n]))
	}
}

func main() {
	args := os.Args
	if len(args) != 3 {
		color.Red("Usage : ./server <port> <name>")
		os.Exit(0)
	}

	port := args[1]
	name := args[2]

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		color.Red(err.Error())
		os.Exit(0)
	}

	defer listen.Close()

	connections := make(chan net.Conn)

	go acceptConnections(listen, connections)

	for conn := range connections {
		go handleConnection(conn, name)
	}
}
