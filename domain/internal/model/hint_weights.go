package model

import (
	"math/rand"
	"time"
)

const (
	MaximumAccessibleTiles = 49
)

type HintWeights struct {
	// Place Tile Weights
	PlayerAccessibleTilesWeight   float64
	OpponentAccessibleTilesWeight float64
	RemainingTileShapeWeight      float64

	// Move Pawn Weights
	TargetDistanceWeight float64
	PlayerDistanceWeight float64
	FixedTileWeight      float64
	MovableTileWeight    float64
	CloseEdgeWeight      float64
	TileShapeWeight      float64
}

func (w HintWeights) getPlaceTileScore(playerAccessibleTile, opponentAccessibleTiles, remainingTileShapeScore float64) float64 {

	var (
		totalWeights = w.PlayerAccessibleTilesWeight +
			w.OpponentAccessibleTilesWeight +
			w.RemainingTileShapeWeight

		playerAccessibleTileScore   = playerAccessibleTile * w.PlayerAccessibleTilesWeight / MaximumAccessibleTiles
		opponentAccessibleTileScore = opponentAccessibleTiles * w.OpponentAccessibleTilesWeight / MaximumAccessibleTiles
		remainingTileScore          = remainingTileShapeScore * w.RemainingTileShapeWeight
	)

	return (playerAccessibleTileScore + opponentAccessibleTileScore + remainingTileScore) / totalWeights
}

func (w HintWeights) getMovePawnScore(targetDistance, playerDistance float64, isFixedTile bool, edgeDistance, tileShape float64) float64 {

	// We are not close to the target, we should go to an edge
	if targetDistance > 0.5 {
		var (
			wrongPathTotalWeights = w.MovableTileWeight +
				w.CloseEdgeWeight +
				w.TileShapeWeight

			closeEdgeScore   = (1 - edgeDistance) * w.CloseEdgeWeight
			tileShapeScore   = tileShape * w.TileShapeWeight
			movableTileScore = 0.0
		)

		if !isFixedTile {
			movableTileScore = w.MovableTileWeight
		}

		return (float64(movableTileScore) + closeEdgeScore + tileShapeScore) / wrongPathTotalWeights
	}

	var (
		goodPathTotalWeights = w.TargetDistanceWeight +
			w.PlayerDistanceWeight +
			w.FixedTileWeight +
			w.CloseEdgeWeight +
			w.TileShapeWeight

		targetDistanceScore = targetDistance * w.TargetDistanceWeight
		playerDistanceScore = playerDistance * w.PlayerDistanceWeight
		fixedTileScore      = 0.0
		edgeDistanceScore   = edgeDistance * w.CloseEdgeWeight
		tileShapeScore      = tileShape * w.TileShapeWeight
	)

	if isFixedTile {
		fixedTileScore = w.FixedTileWeight
	}

	return (targetDistanceScore +
		playerDistanceScore +
		fixedTileScore +
		edgeDistanceScore +
		tileShapeScore) / goodPathTotalWeights
}

func (w *HintWeights) Copy() *HintWeights {
	return &HintWeights{
		PlayerAccessibleTilesWeight:   w.PlayerAccessibleTilesWeight,
		OpponentAccessibleTilesWeight: w.OpponentAccessibleTilesWeight,
		RemainingTileShapeWeight:      w.RemainingTileShapeWeight,

		TargetDistanceWeight: w.TargetDistanceWeight,
		FixedTileWeight:      w.FixedTileWeight,
		MovableTileWeight:    w.MovableTileWeight,
		CloseEdgeWeight:      w.CloseEdgeWeight,
		TileShapeWeight:      w.TileShapeWeight,
	}
}

func NewRandomHintWeights() *HintWeights {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	return &HintWeights{
		PlayerAccessibleTilesWeight:   float64(generator.Intn(100)),
		OpponentAccessibleTilesWeight: float64(generator.Intn(100)),
		RemainingTileShapeWeight:      float64(generator.Intn(100)),

		TargetDistanceWeight: float64(generator.Intn(100)),
		PlayerDistanceWeight: float64(generator.Intn(100)),
		FixedTileWeight:      float64(generator.Intn(100)),
		MovableTileWeight:    float64(generator.Intn(100)),
		CloseEdgeWeight:      float64(generator.Intn(100)),
		TileShapeWeight:      float64(generator.Intn(100)),
	}
}

// NewBestHintWeights are trained using AI
func NewBestHintWeights() *HintWeights {
	return &HintWeights{16, 48, 87, 47, 16, 79, 50, 75, 90}
}
