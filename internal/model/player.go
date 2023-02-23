package model

type Color int

const (
	ColorBlue Color = iota
	ColorGreen
	ColorRed
	ColorYellow
)

// Player represents a player.
type Player struct {

	// Color is the player color
	Color Color `json:"color"`

	// position is the player current position
	Position *Coordinate `json:"position"`

	// Player targets
	Targets []Treasure `json:"targets"`

	// Player score
	Score int `json:"score"`
}

func (p Player) Name() string {
	switch p.Color {
	case ColorBlue:
		return "Blue"
	case ColorGreen:
		return "Green"
	case ColorRed:
		return "Red"
	default:
		return "Yellow"
	}
}
