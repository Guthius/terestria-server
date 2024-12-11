package main

import (
	"fmt"
	"log"
	"time"

	"github.com/guthius/terestria-server/config"
	"github.com/guthius/terestria-server/game"
	"github.com/guthius/terestria-server/net"
)

// handleClientConnected handles the client connection.
func handleClientConnected(id int, client *net.Conn) {
	fmt.Printf("[%04d] Client connected from %s\n", id, client.RemoteAddr())
	game.CreatePlayer(id, client)
}

// handleClientDisconnected handles the client disconnection.
func handleClientDisconnected(id int, client *net.Conn) {
	fmt.Printf("[%04d] Client disconnected from %s\n", id, client.RemoteAddr())
	game.DestroyPlayer(id)
}

// handleDataReceived handles the data received from the client.
func handleDataReceived(id int, _ *net.Conn, bytes []byte) {
	game.HandleDataReceived(id, bytes)
}

func main() {
	networkConfig := net.Config{
		Address:              config.GameAddr,
		MaxConnections:       config.MaxConnections,
		OnClientConnected:    handleClientConnected,
		OnClientDisconnected: handleClientDisconnected,
		OnDataReceived:       handleDataReceived,
	}

	err := net.Start(networkConfig)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for range ticker.C {
		game.Update()
	}
}
