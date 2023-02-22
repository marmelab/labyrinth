package model

type Color int

const (
	ColorBlue Color = 1 << iota
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
}
