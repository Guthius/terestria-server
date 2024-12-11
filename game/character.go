package game

type Character struct {
	Name   string
	Sprite string
	Level  string
	X      int
	Y      int
}

// NewCharacter creates a new character with the given name.
func NewCharacter(name string) *Character {
	return &Character{
		Name:   name,
		Sprite: ")char1",
		Level:  "start_map",
		X:      0,
		Y:      0,
	}
}
