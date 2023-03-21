package model

import (
	"math"
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
		bestDistance               = math.MaxFloat64
	)
	for _, accessibleTile := range accessibleTiles {
		var (
			lineDistance = math.Pow(float64(accessibleTile.Line)-float64(targetCoordinate.Line), 2)
			rowDistance  = math.Pow(float64(accessibleTile.Row)-float64(targetCoordinate.Row), 2)
			distance     = math.Sqrt(lineDistance + rowDistance)
		)

		if distance < bestDistance {
			bestDistance = distance
			bestCoordinate = accessibleTile
		}
	}

	return bestCoordinate, nil
}
