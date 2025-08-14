package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	buffer := make([]byte, 8)
	var s strings.Builder

	for {
		n, err := file.Read(buffer)
		chunk := buffer[:n]
		newLineIndex := len(chunk)

		if err != nil {
			if err == io.EOF {
				s.Write(chunk)
				if s.Len() > 0 {
					fmt.Printf("read: %s\n", s.String())
					s.Reset()
				}
				os.Exit(0)
			}
			log.Fatal("Error reading file:", err)
		}

		for idx, char := range chunk {
			if char == '\n' {
				newLineIndex = idx
				break
			}
		}

		s.Write(chunk[:newLineIndex])

		if newLineIndex < len(chunk) {
			fmt.Printf("read: %s\n", s.String())
			s.Reset()

			if newLineIndex+1 < len(chunk) {
				s.Write(chunk[newLineIndex+1:])
			}
		}
	}
}
