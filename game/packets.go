package game

import "github.com/guthius/terestria-server/net"

// sendLoginResult sends a packet to the specified player with the login result.
func sendLoginResult(player *Player, result int, message string) {
	writer := net.NewPacketWriter()

	writer.WriteInteger(MsgLogin)
	writer.WriteByte(byte(result))
	if result != 0 {
		writer.WriteString(message)
	}

	player.Send(writer.Bytes())
}

// sendJoinGame sends a packet to the specified player to join the game.
func sendJoinGame(player *Player) {
	if player.Character == nil {
		return
	}

	writer := net.NewPacketWriter()

	writer.WriteInteger(MsgJoinGame)
	writer.WriteLong(player.Conn.Id())
	writer.WriteLong(player.Character.X)
	writer.WriteLong(player.Character.Y)
	writer.WriteString(player.Character.Level)

	player.Send(writer.Bytes())
}
