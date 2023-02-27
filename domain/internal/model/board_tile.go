package model

// Rotation represents a tile rotation on a board.
type Rotation int

const (
	Rotation0 Rotation = 0

	Rotation90 Rotation = 90

	Rotation180 Rotation = 180

	Rotation270 Rotation = 270
)

type ShapeRotationExit map[Rotation]TileExits
type ShapeExit map[Shape]ShapeRotationExit

var (
	shapeExitMap = ShapeExit{
		ShapeI: {
			Rotation0:   {TileExitRight, TileExitLeft},
			Rotation90:  {TileExitTop, TileExitBottom},
			Rotation180: {TileExitRight, TileExitLeft},
			Rotation270: {TileExitTop, TileExitBottom},
		},
		ShapeT: {
			Rotation0:   {TileExitTop, TileExitRight, TileExitLeft},
			Rotation90:  {TileExitTop, TileExitRight, TileExitBottom},
			Rotation180: {TileExitRight, TileExitBottom, TileExitLeft},
			Rotation270: {TileExitTop, TileExitBottom, TileExitLeft},
		},
		ShapeV: {
			Rotation0:   {TileExitBottom, TileExitLeft},
			Rotation90:  {TileExitTop, TileExitLeft},
			Rotation180: {TileExitTop, TileExitRight},
			Rotation270: {TileExitRight, TileExitBottom},
		},
	}
)

// BoardTile represents a tile that is placed on a board with a given rotation.
type BoardTile struct {

	// Tile is the underlying tile.
	Tile *Tile `json:"tile"`

	// Rotation is the tile rotation.
	Rotation Rotation `json:"rotation"`
}

// GetExits return sthe possible exits for that tile as a bitmask.
func (bt BoardTile) GetExits() TileExits {
	return shapeExitMap[bt.Tile.Shape][bt.Rotation]
}
