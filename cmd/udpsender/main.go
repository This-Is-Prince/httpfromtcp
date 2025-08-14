package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		log.Fatal("Error resolving udp address:", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Error dialing udp connection:", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, ">")
		data, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error reading string:", err)
		}

		_, err = conn.Write([]byte(data))
		if err != nil {
			log.Fatal("Error writing to udp connection:", err)
		}
	}
}
