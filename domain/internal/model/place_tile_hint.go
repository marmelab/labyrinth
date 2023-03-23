package model

import (
	"fmt"
	"math"
)

type PlaceTileHint struct {
	Direction Direction `json:"direction"`
	Index     int       `json:"index"`
	Rotation  Rotation  `json:"rotation"`
}

var (
	placeTileHintsRotations = []Rotation{Rotation0, Rotation90, Rotation180, Rotation270}
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

	if b.LastInsertion != nil {
		boardCopy.LastInsertion = &TileInsertion{
			Direction: b.LastInsertion.Direction,
			Index:     b.LastInsertion.Index,
		}
	}

	return boardCopy
}

func (b *Board) GetPlaceTileHint() (*Board, *PlaceTileHint) {
	var (
		foundShortestPath                = false
		largestScore                     = -1.0
		bestBoardCopy     *Board         = nil
		bestPlaceTileHint *PlaceTileHint = nil
	)

	for _, insertion := range b.getAvailableInsertions() {
		for _, rotation := range placeTileHintsRotations {
			boardCopy := b.Copy()
			boardCopy.RemainingTile.Rotation = rotation
			if err := boardCopy.InsertTileAt(insertion.Direction, insertion.Index); err != nil {
				fmt.Printf("Failed to insert tile: %v\n", err)
			}

			_, isShortestPath := boardCopy.GetAccessibleTiles()
			hint := &PlaceTileHint{
				Direction: insertion.Direction,
				Index:     insertion.Index,
				Rotation:  rotation,
			}

			score := boardCopy.getPlaceTileScore()
			if (isShortestPath && !foundShortestPath) ||
				((isShortestPath == foundShortestPath) && score > largestScore) {
				foundShortestPath = isShortestPath || foundShortestPath
				largestScore = score
				bestBoardCopy = boardCopy
				bestPlaceTileHint = hint
			}
		}
	}

	return bestBoardCopy, bestPlaceTileHint
}

func (b *Board) getRemainingTileShapeScore() float64 {
	switch b.RemainingTile.Tile.Shape {
	case ShapeI:
		return 0.5
	case ShapeV:
		return 1.0
	default:
		return 0.0
	}
}

// getPlaceTileScore returns the board score that optimizes the number of available tiles for the player while minimizing the number of available tiles for the next player
func (b *Board) getPlaceTileScore() float64 {
	var (
		accessibleTiles, _                      = b.GetAccessibleTiles()
		accessibleTilesCount                    = float64(len(accessibleTiles))
		isOnlyRemainingPlayer                   = len(b.RemainingPlayers) == 1
		maxOpponentAccessibleTilesCount float64 = 1.0
		weights                                 = b.GetCurrentPlayer().GetWeights()
		maxScore                        float64 = 0
	)

	coordinate, err := b.GetMovePawnHint()
	if err != nil {
		fmt.Printf("Failed to get move pawn hint: %v\n", err)
	}
	for _, insertion := range b.getAvailableInsertions() {
		for _, rotation := range placeTileHintsRotations {
			boardCopy := b.Copy()
			boardCopy.MoveCurrentPlayerTo(coordinate.Line, coordinate.Row)
			boardCopy.RemainingTile.Rotation = rotation
			boardCopy.InsertTileAt(insertion.Direction, insertion.Index)

			opponentAccessibleTiles, isShortestPath := boardCopy.GetAccessibleTiles()
			// This is the best move if we are the only remaining player
			if isOnlyRemainingPlayer && isShortestPath {
				return math.MaxFloat64
			}

			opponentAccessibleTilesCount := float64(len(opponentAccessibleTiles))

			score := weights.getPlaceTileScore(
				accessibleTilesCount, opponentAccessibleTilesCount, boardCopy.getRemainingTileShapeScore())

			if score >= maxScore {
				maxScore = score
				maxOpponentAccessibleTilesCount = opponentAccessibleTilesCount
			}
		}
	}

	// This ensures that if we are the only player, we want to maximize the possibilities for all turns
	if isOnlyRemainingPlayer {
		return accessibleTilesCount * maxOpponentAccessibleTilesCount
	}

	return maxScore
}
