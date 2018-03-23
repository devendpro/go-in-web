package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		io.WriteString(conn, "Como esta o seu dia?\n")
		io.WriteString(conn, fmt.Sprintf("%v", "Espero  que esteja bem!\n"))
		conn.Close()
	}
}
