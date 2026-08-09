# Go Multi-Client TCP Chat Server

A concurrent multi-client chat server built in Go to understand TCP networking, goroutines, channels, and concurrent server architecture.

The server allows multiple clients to connect simultaneously, send messages, and receive broadcasts from other connected clients.

## Features

- TCP server using Go's `net` package
- Multiple clients connected simultaneously
- One goroutine per client connection
- Channel-based communication between client handlers and the server manager
- Event-driven client management
- Message broadcasting
- TCP read timeouts
- Graceful client disconnection handling
- Newline-based application-level message framing

## Architecture

The server uses a simple event-driven architecture.

<img width="357" height="687" alt="image" src="https://github.com/user-attachments/assets/2ec99e4c-588f-4c40-bf58-b6f019d3bfa4" />


How to Run
1. Clone the repository
git clone <your-repository-url>
cd go-multiclient-chat-server
2. Start the server
go run server.go

You should see:

Waiting for clients...
3. Start clients

Open two or more terminals and run:

go run client.go

Each client connects to:

127.0.0.1:5000
4. Send messages

Type a message in one client.

The server broadcasts the message to the connected clients.

Example Output
Server
Waiting for clients...
Client connected: 127.0.0.1:53834
Total clients: 1

Client connected: 127.0.0.1:53836
Total clients: 2

Message from 127.0.0.1:53834: "hello\n"
Message from 127.0.0.1:53836: "hi\n"

Client disconnected: 127.0.0.1:53834
Total clients: 1
Client
You: hello

Received: 127.0.0.1:53834: hello

You: hi
What I Learned

This project was built primarily as a learning project to understand how concurrent network servers work in Go.

Key concepts learned:

TCP client/server communication
TCP connections and ports
Blocking network I/O
Goroutines
Channels
Channel ownership
select
Concurrent client handling
Race conditions
Mutexes
Event-driven architecture
TCP read deadlines and timeouts
Connection lifecycle
Application-level message framing
bufio.Reader
Separation of client handling and server state management

