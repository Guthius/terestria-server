package game

import "github.com/guthius/terestria-server/net"

const BufferSize = 4096

type Player struct {
	Conn      *net.Conn
	Character *Character
	Room      *Room
	Buffer    []byte
}

// NewPlayer creates a new player with the specified connection.
func NewPlayer(conn *net.Conn) *Player {
	return &Player{
		Conn:   conn,
		Buffer: make([]byte, 0, BufferSize),
	}
}

// Send the specified bytes to the player.
func (p *Player) Send(bytes []byte) {
	if p == nil || p.Conn == nil {
		return
	}

	size := len(bytes)
	if size == 0 {
		return
	}

	sizeBytes := []byte{byte(size), byte(size >> 8)}

	p.Conn.Send(sizeBytes)
	p.Conn.Send(bytes)
}

// Disconnect closes the connection with the player.
func (player *Player) Disconnect() {
	if player == nil || player.Conn == nil {
		return
	}
	player.Conn.Close()
}

// IsPlaying returns true if the player is in game.
func (player *Player) IsPlaying() bool {
	return player.Character != nil
}
