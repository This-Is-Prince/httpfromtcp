package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	buffer := make([]byte, 8)

	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				os.Exit(0)
			}
			log.Fatal("Error reading file:", err)
		}

		chunk := buffer[:n]

		fmt.Printf("read: %s\n", chunk)
	}
}
