package model

type Treasure rune

// NoTreasure is when tile has no treasure in it.
const NoTreasure Treasure = '.'

func (t Treasure) MarshalJSON() ([]byte, error) {
	return []byte(`"` + string(t) + `"`), nil
}

func (t *Treasure) UnmarshalJSON(data []byte) error {
	runes := []rune(string(data))
	*t = Treasure(runes[1])
	return nil
}

// Shape reprents a tile shape.
type Shape int

const (
	// ShapeI represents an I shape.
	ShapeI Shape = iota

	// ShapeT represents a T shape.
	ShapeT

	// ShapeV represents a V shape.
	ShapeV
)

// Tile represents a tile.
type Tile struct {

	// Shape is the tile shape.
	Shape Shape `json:"shape"`

	// Treasure is the optional tile treasure if present, '.' otherwise.
	Treasure Treasure `json:"treasure"`
}
