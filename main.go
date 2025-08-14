package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)

		buffer := make([]byte, 8)
		var s strings.Builder

		for {
			n, err := f.Read(buffer)
			if err != nil {
				if err == io.EOF {
					s.Write(buffer[:n])
					break
				} else {
					log.Fatal("Error reading file:", err)
				}
			}

			chunk := buffer[:n]
			if idx := bytes.IndexByte(chunk, '\n'); idx != -1 {
				s.Write(chunk[:idx])
				chunk = chunk[idx+1:]
				lines <- s.String()
				s.Reset()
			}

			s.Write(chunk)
		}

		if s.Len() > 0 {
			lines <- s.String()
			s.Reset()
		}
	}()

	return lines
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error listening on port :42069 :", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}

		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
	}
}
