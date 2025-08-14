package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	for str := range getLinesChannel(file) {
		fmt.Printf("read: %s\n", str)
	}
}
