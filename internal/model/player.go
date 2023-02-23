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

	// Line is the player current line.
	Line int `json:"line"`

	// Row is the current player row.
	Row int `json:"row"`

	// Player hand
	Targets []Treasure `json:"hand"`

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
