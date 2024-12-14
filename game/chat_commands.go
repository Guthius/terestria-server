package game

func handleChatCommand(player *Player, command string) {
	if command == "/help" {
		player.SendNotification("Available commands:")
		player.SendNotification("/help - Display this help message")
		player.SendNotification("/who - List all players in the room")
	} else if command == "/who" {
		player.Room.SendNotification("Players in the room:")
		for _, p := range player.Room.Players {
			if p.IsPlaying() {
				player.Room.SendNotification(p.Character.Name)
			}
		}
	} else {
		player.SendNotification("Unknown command. Type /help for a list of available commands.")
	}
}
