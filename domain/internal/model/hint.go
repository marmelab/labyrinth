package model

type PlaceTileHint struct {
	Direction Direction `json:"direction"`
	Index     int       `json:"index"`
	Rotation  Rotation  `json:"rotation"`
}

var (
	placeTileHintsDirections = []Direction{DirectionTop, DirectionRight, DirectionBottom, DirectionLeft}
	placeTileHintsIndexes    = []int{1, 3, 5}
	placeTileHintsRotations  = []Rotation{Rotation0, Rotation90, Rotation180, Rotation270}
)

func (b *Board) Copy() *Board {
	boardCopy := &Board{
		Tiles:                make([][]*BoardTile, len(b.Tiles)),
		RemainingTile:        b.RemainingTile.Copy(),
		Players:              make([]*Player, 0, len(b.Players)),
		RemainingPlayers:     make([]int, len(b.RemainingPlayers)),
		RemainingPlayerIndex: b.RemainingPlayerIndex,
		State:                b.State,
	}

	for line, tileLine := range b.Tiles {
		boardCopy.Tiles[line] = make([]*BoardTile, len(tileLine))
		for row, boardTile := range tileLine {
			boardCopy.Tiles[line][row] = boardTile.Copy()
		}
	}

	for _, player := range b.Players {
		boardCopy.Players = append(boardCopy.Players, player.Copy())
	}

	copy(boardCopy.RemainingPlayers, b.RemainingPlayers)

	return boardCopy
}

func (b *Board) GetPlaceTileHint() (*Board, *PlaceTileHint) {
	var (
		largestAccessibleTiles int            = 0
		bestBoardCopy          *Board         = nil
		bestPlaceTileHint      *PlaceTileHint = nil
	)

	for _, direction := range placeTileHintsDirections {
		for _, index := range placeTileHintsIndexes {
			for _, rotation := range placeTileHintsRotations {
				boardCopy := b.Copy()
				boardCopy.RemainingTile.Rotation = rotation
				boardCopy.InsertTileAt(direction, index)

				accessibleTiles, isShortestPath := boardCopy.GetAccessibleTiles()
				hint := &PlaceTileHint{
					Direction: direction,
					Index:     index,
					Rotation:  rotation,
				}
				if isShortestPath {
					return boardCopy, hint
				}

				if accessibleTileCount := len(accessibleTiles); accessibleTileCount > largestAccessibleTiles {
					largestAccessibleTiles = accessibleTileCount
					bestBoardCopy = boardCopy
					bestPlaceTileHint = hint
				}
			}
		}
	}

	return bestBoardCopy, bestPlaceTileHint
}
