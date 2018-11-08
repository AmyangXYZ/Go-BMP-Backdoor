package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	port := flag.String("port", "1337", "the port your server listen on")
	flag.Parse()

	ln, err := net.Listen("tcp", ":"+*port)
	fmt.Println("Listening on", *port, "...")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()
	conn, _ := ln.Accept()
	fmt.Println("Connection from", conn.RemoteAddr().String())
	for {
		// scan payload
		fmt.Print(">>> ")
		reader := bufio.NewReader(os.Stdin)
		payload, _ := reader.ReadString('\n')

		// write payload to server_send.bmp
		err := write(payload, "server_send.bmp")
		if err != nil {
			fmt.Println(err)
			return
		}

		// send server_send.bmp
		err = send(conn, "server_send.bmp")
		if err != nil {
			fmt.Println(err)
			return
		}

		// save server_recv.bmp
		err = recv(conn, "server_recv.bmp")
		if err != nil {
			fmt.Println(err)
			return
		}

		// read output from server_recv.bmp
		ret, err := read("server_recv.bmp")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(ret)
	}
}
