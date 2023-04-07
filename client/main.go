package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	fmt.Print("Name ? ")
	var name string
	fmt.Scan(&name)

	conn, err := net.Dial("tcp", "10.56.148.72:1234")

	if err != nil {
		color.Red(err.Error())
	}

	color.Red("Connected !")

	var msg string

	defer conn.Close()

	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			color.Red("Server has closed.")
			exec.Command("pause")
			return
		}

		color.Cyan("%s", string(buf[:n]))

		fmt.Print(">> ")
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		msg = sc.Text()
		if strings.Trim(msg, " ") == "quit" {
			conn.Close()
			os.Exit(0)
		}
		conn.Write([]byte(name + ">> " + msg))
	}
}
