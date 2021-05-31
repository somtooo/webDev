package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	println("Listening on port 8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()

		if err != nil {
			log.Panicln(err)
		}

		io.WriteString(conn, "\nHello from tcp server\n")
		// conn.Close()
		go handle(conn)

	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("Code got here")
}
