package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
	conn net.Conn
}

type EventType int

const (
	ClientConnected EventType = iota
	ClientMessage
	ClientDisconnected
)

type Event struct {
	Type   EventType
	Client *Client
	Text   string
}

func handleClient(client *Client, events chan<- Event) {
	defer client.conn.Close()

	events <- Event{
		Type:   ClientConnected,
		Client: client,
	}

	reader := bufio.NewReader(client.conn)

	for {
		client.conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		message, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed:", client.conn.RemoteAddr())
			} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				fmt.Println("10-second idle timeout:",
					client.conn.RemoteAddr())
				continue
			} else {
				log.Println("Read error:", err)
			}

			events <- Event{
				Type:   ClientDisconnected,
				Client: client,
			}

			return
		}

		events <- Event{
			Type:   ClientMessage,
			Client: client,
			Text:   message,
		}
	}
}

func manager(events <-chan Event) {

	clients := make(map[*Client]bool)

	for event := range events {

		switch event.Type {

		case ClientConnected:

			clients[event.Client] = true

			fmt.Println(
				"Client connected:",
				event.Client.conn.RemoteAddr(),
			)

			fmt.Println("Total clients:", len(clients))

		case ClientMessage:

			fmt.Printf(
				"Message from %s: %q\n",
				event.Client.conn.RemoteAddr(),
				event.Text,
			)
			
			for client := range clients {

				_, err := client.conn.Write(
					[]byte(
						fmt.Sprintf(
							"%s: %s",
							event.Client.conn.RemoteAddr(),
							event.Text,
						),
					),
				)

				if err != nil {
					log.Println("Write error:", err)
				}
			}

		case ClientDisconnected:

			delete(clients, event.Client)

			fmt.Println(
				"Client disconnected:",
				event.Client.conn.RemoteAddr(),
			)

			fmt.Println("Total clients:", len(clients))
		}
	}
}

func main() {

	listener, err := net.Listen(
		"tcp",
		"127.0.0.1:5000",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	
	events := make(chan Event)

	// Start manager
	go manager(events)

	fmt.Println("Waiting for clients...")

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		client := &Client{
			conn: conn,
		}

		go handleClient(client, events)
	}
}