package main

import (
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

const buffersize int64 = 1024

func send(conn net.Conn, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	fileInfo, _ := file.Stat()
	fileSize := fillString(strconv.FormatInt(fileInfo.Size(), 10), 10)

	conn.Write([]byte(fileSize))

	sendBuffer := make([]byte, buffersize)
	for {
		_, err = file.Read(sendBuffer)
		if err == io.EOF {
			break
		}
		conn.Write(sendBuffer)
	}
	return nil
}

func recv(conn net.Conn, filename string) error {
	bufferFileSize := make([]byte, 10)
	conn.Read(bufferFileSize)
	fileSize, _ := strconv.ParseInt(strings.Trim(string(bufferFileSize), ":"), 10, 64)

	newFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer newFile.Close()
	var receivedBytes int64
	for {
		if (fileSize - receivedBytes) < buffersize {
			io.CopyN(newFile, conn, (fileSize - receivedBytes))
			conn.Read(make([]byte, (receivedBytes+buffersize)-fileSize))
			break
		}
		io.CopyN(newFile, conn, buffersize)
		receivedBytes += buffersize
	}
	return nil
}

func fillString(retunString string, toLength int) string {
	for {
		lengtString := len(retunString)
		if lengtString < toLength {
			retunString = retunString + ":"
			continue
		}
		break
	}
	return retunString
}
