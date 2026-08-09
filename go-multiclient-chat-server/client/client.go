package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		buf := make([]byte, 1024)

		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Println("Server connection closed")
				return
			}

			fmt.Printf("\nReceived: %s", buf[:n])
			fmt.Print("You: ")
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("You: ")

		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
	}
}