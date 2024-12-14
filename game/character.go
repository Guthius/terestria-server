package game

type Character struct {
	Name      string
	Sprite    string
	Map       string
	Direction int
	X         int
	Y         int
}

// NewCharacter creates a new character with the given name.
func NewCharacter(name string) *Character {
	return &Character{
		Name:      name,
		Sprite:    ")char1",
		Map:       "start_map",
		Direction: DirDown,
		X:         1,
		Y:         5,
	}
}
