package model

import (
	"math"
)

var (
	maxTileDistance float64 = math.Sqrt(98)
	maxEdgeDistance float64 = 3
)

func (b *Board) GetMovePawnHint() (*Coordinate, error) {
	if b.State != GameStateMovePawn {
		return nil, ErrInvalidAction
	}

	accessibleTiles, isShortestPath := b.GetAccessibleTiles()
	if isShortestPath {
		lastTile := accessibleTiles[len(accessibleTiles)-1]
		return lastTile, nil
	}

	targetCoordinate := b.GetCurrentTargetCoordinate()
	if targetCoordinate == nil {
		targetCoordinate = &Coordinate{3, 3}
	}

	var (
		bestCoordinate *Coordinate = nil
		bestScore      float64     = 0
	)
	for _, accessibleTile := range accessibleTiles {
		score := b.getAccessibleTileScore(targetCoordinate, accessibleTile)
		if score >= bestScore {
			bestCoordinate = accessibleTile
			bestScore = score
		}
	}

	return bestCoordinate, nil
}

func (b *Board) isFixedTile(tile *Coordinate) bool {
	return ((tile.Line & 1) == 0) && ((tile.Row & 1) == 0)
}

func (b *Board) getEndTurnShapeScore(coordinates *Coordinate) float64 {
	tile := b.Tiles[coordinates.Line][coordinates.Row]
	switch tile.Tile.Shape {
	case ShapeI:
		return 0.0
	case ShapeV:
		return 0.5
	default:
		return 1.0
	}
}

func (b *Board) getEdgeDistanceScore(tile *Coordinate) float64 {
	var (
		topDistance    = float64(tile.Row)
		rightDistance  = float64(len(b.Tiles) - 1 - tile.Line)
		bottomDistance = float64(len(b.Tiles) - 1 - tile.Row)
		leftDistance   = float64(tile.Line)

		yDistance = math.Min(topDistance, bottomDistance)
		xDistance = math.Min(rightDistance, leftDistance)
		distance  = math.Min(yDistance, xDistance)
	)

	return (distance / maxEdgeDistance)
}

func (b *Board) getAccessibleTileScore(targetCoordinate *Coordinate, accessibleTile *Coordinate) float64 {
	var (
		currentPlayer         = b.GetCurrentPlayer()
		currentPlayerPosition = currentPlayer.Position
		weights               = currentPlayer.GetWeights()

		targetLineDistance = math.Pow(float64(accessibleTile.Line)-float64(targetCoordinate.Line), 2)
		targetRowDistance  = math.Pow(float64(accessibleTile.Row)-float64(targetCoordinate.Row), 2)
		targetDistance     = math.Sqrt(targetLineDistance+targetRowDistance) / maxTileDistance

		playerLineDistance = math.Pow(float64(accessibleTile.Line)-float64(currentPlayerPosition.Line), 2)
		playerRowDistance  = math.Pow(float64(accessibleTile.Row)-float64(currentPlayerPosition.Row), 2)
		playerDistance     = math.Sqrt(playerLineDistance+playerRowDistance) / maxTileDistance
	)

	return weights.getMovePawnScore(
		targetDistance,
		playerDistance,
		b.isFixedTile(accessibleTile),
		b.getEdgeDistanceScore(accessibleTile),
		b.getEndTurnShapeScore(accessibleTile))
}
