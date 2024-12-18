package game

import (
	"encoding/binary"
	"fmt"

	"github.com/guthius/terestria-server/net"
)

const (
	_        = iota
	MsgLogin = iota
	MsgJoinGame
	MsgAddPlayer
	MsgRemovePlayer
	MsgMovePlayer
	MsgSetPlayerPosition
	MsgSetPlayerDirection
	MsgChangeMap
	MsgAttack
	MsgChat
	MsgNotification

	MaxMessageId = iota
)

type PacketHandler func(player *Player, reader *net.PacketReader)

var PacketHandlers [MaxMessageId]PacketHandler

func init() {
	PacketHandlers[MsgLogin] = handleLogin
	PacketHandlers[MsgMovePlayer] = handleMovePlayer
	PacketHandlers[MsgSetPlayerDirection] = handleSetPlayerDirection
	PacketHandlers[MsgAttack] = handleAttack
	PacketHandlers[MsgChat] = handleChat
}

// HandleDataReceived handles the data received from the specified player.
func HandleDataReceived(id int, bytes []byte) {
	const headerSize = 2

	fmt.Printf("[%04d] Received %d bytes from player\n", id, len(bytes))

	player, ok := players[id]
	if !ok {
		return
	}

	player.Buffer = append(player.Buffer, bytes...)
	if len(player.Buffer) < headerSize {
		return
	}

	buf := player.Buffer
	off := 0

	// Handle all packets in the buffer
	for len(buf) >= headerSize {
		size := int(binary.LittleEndian.Uint16(buf))
		if len(buf) < size+headerSize {
			return
		}
		off += headerSize
		buf = buf[headerSize:]

		reader := net.NewPacketReader(buf[:size])
		handlePacket(player, reader)

		off += size
		buf = buf[size:]
	}

	// Move the bytes that are remaining to the front of the buffer
	bytesLeft := len(player.Buffer) - off
	if bytesLeft > 0 {
		copy(player.Buffer, player.Buffer[off:])
	}

	player.Buffer = player.Buffer[:bytesLeft]
}

// handlePacket handles the specified packet for the player.
func handlePacket(player *Player, reader *net.PacketReader) {
	if reader.Remaining() < 2 {
		return
	}

	packetId := reader.ReadInteger()
	if packetId < 0 || packetId >= MaxMessageId {
		return
	}

	packetHandler := PacketHandlers[packetId]
	if packetHandler == nil {
		return
	}

	packetHandler(player, reader)
}

// handleLogin handles the login packet for the player.
func handleLogin(player *Player, reader *net.PacketReader) {
	if player.IsPlaying() {
		return
	}

	name := reader.ReadString()
	if len(name) < 3 {
		return
	}

	player.Character = NewCharacter(name)
	player.Room = room

	sendLoginResult(player, 0, "")
	sendJoinGame(player)

	player.Room.AddPlayer(player)
}

// handleMovePlayer handles the move player packet for the player.
func handleMovePlayer(player *Player, reader *net.PacketReader) {
	if player.Room == nil {
		return
	}

	dir := reader.ReadByte()

	player.Room.MovePlayer(player, int(dir))
}

// handleSetPlayerDirection handles the set player direction packet for the player.
func handleSetPlayerDirection(player *Player, reader *net.PacketReader) {
	if player.Room == nil {
		return
	}

	dir := reader.ReadByte()

	player.Room.SetPlayerDirection(player, int(dir))
}

// handleAttack handles the attack packet for the player.
func handleAttack(player *Player, reader *net.PacketReader) {
	if player.Room == nil {
		return
	}

	dir := reader.ReadByte()

	player.Room.Attack(player, int(dir))
}

// handleChat handles the chat packet for the player.
func handleChat(player *Player, reader *net.PacketReader) {
	if player.Room == nil {
		return
	}

	message := reader.ReadString()

	player.Room.Chat(player, message)
}
