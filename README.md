# Real-Time WebSocket Chat Server in Golang

## 📌 Overview
This project is a **real-time chat server** built using **Golang** and **WebSockets**. It allows multiple users to connect, send messages, and receive updates instantly through WebSocket connections.

---

## 🔹 What is WebSocket?
### **WebSocket vs. HTTP**
- **HTTP**: A request-response protocol where the client sends a request, and the server responds once. It is not ideal for real-time communication.
- **WebSocket**: A persistent, full-duplex communication channel where both the client and server can send messages to each other **at any time**. Perfect for real-time applications like chat.

### **How WebSocket Works**
1. The **client** initiates a WebSocket handshake with the server.
2. If the handshake is successful, a **persistent connection** is established.
3. Both the **client and server** can send messages to each other at any time.
4. The connection remains **open** until either the client or server closes it.

---

## 🚀 Features
- **Real-time messaging**: Messages are sent and received instantly.
- **Broadcasting**: Every message is sent to **all connected users**.
- **Connection tracking**: Logs when a user **connects** or **disconnects**.
- **Goroutines & Channels**: Efficiently handles multiple connections concurrently.
- **Structured logging**: Uses the Zap logger for clean and informative logs.

---

## 📂 Project Structure
```
/realtime-chat
│── /cmd
│    ├── /server
│    │    ├── main.go               # Entry point
│── /internal
│    ├── /config                    # Configuration management
│    │    ├── read_config.go
│    ├── /model                     # Data models
│    │    ├── message_model.go
│    ├── /app                       # Core WebSocket logic
│    │    ├── hub.go                # Manages connected clients & messages
│    │    ├── websocket.go          # WebSocket connection handling
│    ├── /controller                # API controller
│    │    ├── ws_controller.go      # WebSocket route controller
│── /pkg
|    ├── /logger                    # Logger setup
│         ├── logger.go
│── go.mod
│── go.sum
```

## 🔹 How the WebSocket System Works
### **1️⃣ The Hub (`hub.go`)**
- **Manages all active clients**.
- **Broadcasts messages** to all connected users.
- **Handles user connections & disconnections**.

```go
type Hub struct {
    Clients    map[*Client]bool  // Connected clients
    Broadcast  chan model.Message // Message queue
    Register   chan *Client // New clients joining
    Unregister chan *Client // Clients leaving
}
```

### **2️⃣ Handling WebSocket Connections (`websocket.go`)**
- Reads messages from the client.
- Sends messages to the **Hub** for broadcasting.

```go
func HandleWebSocket(hub *Hub, conn *websocket.Conn) {
    client := &Client{Conn: conn, Send: make(chan model.Message)}
    hub.Register <- client // Register client

    go client.ReadMessages(hub)  // Read incoming messages
    go client.WriteMessages()    // Send messages to the client
}
```

### **3️⃣ Broadcasting Messages (`hub.go`)**
- When a client sends a message, it is **written** to `hub.Broadcast`.
- The **Hub listens** for new messages and sends them to **all active clients**.

```go
case message := <-h.Broadcast:
    for client := range h.Clients {
        select {
        case client.Send <- message:
        default:
            delete(h.Clients, client) // Remove inactive clients
            close(client.Send)
            client.Conn.Close()
        }
    }
```

## 🚀 How to Run This Project
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/your-username/realtime-chat.git
cd realtime-chat
```

### **2️⃣ Install Dependencies**
```sh
go mod tidy
```

### **3️⃣ Start the WebSocket Server**
```sh
go run main.go
```

### **4️⃣ Test WebSockets in Postman**
1️⃣ Open Postman.
2️⃣ Select **"WebSocket Request"**.
3️⃣ Connect to:
```
ws://localhost:8080/ws
```
4️⃣ Send a JSON message:
```json
{
    "sender": "Abhi",
    "content": "Hello, World!"
}
```
5️⃣ Open another WebSocket connection and verify **real-time message delivery**.

## 📌 Example WebSocket Flow
```
Client A connects ✅
Client B connects ✅
Abhi sends "Hello, World!" 📩
📢 Broadcasting "Hello, World!" to all clients
Bob receives "Hello, World!" ✅
```

## 🛠 Technologies Used
- **Golang** - Backend language
- **Gin** - HTTP framework
- **Gorilla WebSocket** - WebSocket library
- **Zap Logger** - Structured logging

## 🔹 Future Enhancements
✅ **Private Messaging** (1-on-1 chat instead of broadcasting).<br>
✅ **User Authentication** (JWT-based authentication for security).<br>
✅ **Message History** (Save chat logs in a database).<br>
✅ **Online Presence** (Show which users are online).

## 🎯 Conclusion
This project provides a **clean and maintainable WebSocket implementation in Golang**. It is designed to be **scalable**, using channels and goroutines for handling concurrent connections efficiently.