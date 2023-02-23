package model

// TileExit represents the possible exits that are available for the tile.
type TileExit int

const (
	TileExitTop TileExit = iota
	TileExitRight
	TileExitBottom
	TileExitLeft
)

var (
	oppositeTileExit = map[TileExit]TileExit{
		TileExitTop:    TileExitBottom,
		TileExitRight:  TileExitLeft,
		TileExitBottom: TileExitTop,
		TileExitLeft:   TileExitRight,
	}
)

func (t TileExit) Opposite() TileExit {
	return oppositeTileExit[t]
}

func (t TileExit) ExitCoordinate(line, row int) *Coordinate {
	switch t {
	case TileExitTop:
		return &Coordinate{Line: line - 1, Row: row}
	case TileExitRight:
		return &Coordinate{Line: line, Row: row + 1}
	case TileExitBottom:
		return &Coordinate{Line: line + 1, Row: row}
	default:
		return &Coordinate{Line: line, Row: row - 1}
	}
}

// TileExists represents a list of tile exists
type TileExits []TileExit

// Contains returns whether the TileExits list contain sthe targeted exit.
func (t TileExits) Contains(target TileExit) bool {
	for _, exit := range t {
		if exit == target {
			return true
		}
	}
	return false
}
