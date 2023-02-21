package model

// NoTreasure is when tile has no treasure in it.
const NoTreasure = '.'

// Shape reprents a tile shape.
type Shape int

const (
	// ShapeI represents an I shape.
	ShapeI Shape = iota

	// ShapeT represents a T shape.
	ShapeT Shape = iota

	// ShapeV represents a V shape.
	ShapeV Shape = iota
)

// Tile represents a tile.
type Tile struct {

	// Shape is the tile shape.
	Shape Shape

	// Treasure is the optional tile treasure if present, '.' otherwise.
	Treasure rune
}
