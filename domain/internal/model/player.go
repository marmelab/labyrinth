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

	// Weigths is the hint weights that are applied.
	Weights *HintWeights `json:"weights"`
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

func (p *Player) GetWeights() *HintWeights {
	if p.Weights == nil {
		p.Weights = NewBestHintWeights()
	}

	return p.Weights
}

func (p *Player) Copy() *Player {
	if p.Weights == nil {
		p.Weights = NewBestHintWeights()
	}

	playerCopy := &Player{
		Color: p.Color,
		Position: &Coordinate{
			Line: p.Position.Line,
			Row:  p.Position.Row,
		},
		Targets: make([]Treasure, len(p.Targets)),
		Score:   p.Score,
		Weights: p.Weights.Copy(),
	}
	copy(playerCopy.Targets, p.Targets)
	return playerCopy
}
