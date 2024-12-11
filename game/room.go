package game

import (
	"fmt"

	"github.com/guthius/terestria-server/net"
)

type Room struct {
	Players []*Player
}

// NewRoom creates a new room.
func NewRoom() *Room {
	return &Room{
		Players: make([]*Player, 0, 20),
	}
}

// Send a packet with the specified bytes to all players on the level.
func (room *Room) Send(bytes []byte) {
	for _, p := range room.Players {
		p.Send(bytes)
	}
}

// SendExcept sends a packet with the specified bytes to all players on the level except the specified player.
func (room *Room) SendExcept(bytes []byte, except *Player) {
	for _, p := range room.Players {
		if p == except {
			continue
		}
		p.Send(bytes)
	}
}

// AddPlayer adds a player to the room.
func (room *Room) AddPlayer(player *Player) {
	room.Players = append(room.Players, player)

	fmt.Printf("[%04d] Player %s joined the room\n", player.Conn.Id(), player.Character.Name)

	for _, other := range room.Players {
		if !other.IsPlaying() || other == player {
			continue
		}

		writer := net.NewPacketWriter()
		writer.WriteInteger(MsgAddPlayer)
		writer.WriteLong(other.Conn.Id())
		writer.WriteString(other.Character.Name)
		writer.WriteString(other.Character.Sprite)
		writer.WriteLong(other.Character.X)
		writer.WriteLong(other.Character.Y)

		player.Send(writer.Bytes())
	}

	writer := net.NewPacketWriter()
	writer.WriteInteger(MsgAddPlayer)
	writer.WriteLong(player.Conn.Id())
	writer.WriteString(player.Character.Name)
	writer.WriteString(player.Character.Sprite)
	writer.WriteLong(player.Character.X)
	writer.WriteLong(player.Character.Y)

	room.Send(writer.Bytes())
}

// RemovePlayer removes the specified Player from the level.
func (room *Room) RemovePlayer(player *Player) {
	for i, p := range room.Players {
		if p == player {
			room.Players = append(room.Players[:i], room.Players[i+1:]...)
			break
		}
	}

	fmt.Printf("[%04d] Player %s left the room\n", player.Conn.Id(), player.Character.Name)

	writer := net.NewPacketWriter()
	writer.WriteInteger(MsgRemovePlayer)
	writer.WriteLong(player.Conn.Id())

	room.Send(writer.Bytes())
}

// MovePlayer moves the specified player in the specified direction.
func (room *Room) MovePlayer(player *Player, dir int) {
	switch dir {
	case DirUp:
		player.Character.Y--
	case DirDown:
		player.Character.Y++
	case DirLeft:
		player.Character.X--
	case DirRight:
		player.Character.X++
	}

	fmt.Printf("[%04d] Player %s moved to (%d, %d)\n", player.Conn.Id(), player.Character.Name, player.Character.X, player.Character.Y)

	writer := net.NewPacketWriter()
	writer.WriteInteger(MsgMovePlayer)
	writer.WriteLong(player.Conn.Id())
	writer.WriteByte(byte(dir))

	room.SendExcept(writer.Bytes(), player)
}
