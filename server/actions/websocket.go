package actions

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// allowedOrigin := "http://your-nextjs-app-domain"
		// requestOrigin := r.Header.Get("Origin")
		// return allowedOrigin == requestOrigin
		return true 
	},
}

var clients = make(map[*websocket.Conn]bool)

func SendWebSocketMessageToClients(c *gin.Context) {
	// Set headers to allow all origins
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	// If it's a preflight request, return immediately with a 200 status
	if c.Request.Method == http.MethodOptions {
		fmt.Println("Handling preflight request")
		c.AbortWithStatus(http.StatusOK)
		return
	}

    // Broadcast the message to all connected clients
    fmt.Println("SendWebSocketMessageToClients called")

	data := c.Query("blogg_input")

	fmt.Println("data: ", data)

    for client := range clients {
        fmt.Println("Sending message to client")
        err := client.WriteMessage(websocket.TextMessage, []byte(data))
        if err != nil {
            fmt.Println("Error sending message:", err)
            // Handle the error, e.g., remove the client from the list
            delete(clients, client)
        }
    }
}

func SendWebsocket(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Print("upgrade failed: ", err)
		return
	}

	clients[conn] = true

}