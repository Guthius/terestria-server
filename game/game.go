package game

import "github.com/guthius/terestria-server/net"

var players map[int]*Player
var room *Room

func init() {
	players = make(map[int]*Player)
	room = NewRoom()
}

// Update is the main game loo
func Update() {
}

// CreatePlayer creates a new player with the specified id and connection.
func CreatePlayer(id int, conn *net.Conn) {
	if _, ok := players[id]; !ok {
		delete(players, id)
	}
	players[id] = NewPlayer(conn)
}

// DestroyPlayer destroys the player with the specified id.
func DestroyPlayer(id int) {
	if player, ok := players[id]; ok {
		if player.Room != nil {
			player.Room.RemovePlayer(player)
		}
		delete(players, id)
	}
}
