package main

import (
	"fmt"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

func main() {
	lastWinningWeigths := model.NewBestHintWeights()

outer:
	for round := 0; round < 100; round++ {
		fmt.Printf("----------\n")
		fmt.Printf("[Round #%d] Start\n", round)

		board, _ := model.NewBoard(7, 4)

		// We want to keep some weights for optimization
		for _, player := range board.Players {
			switch player.Color {
			case model.ColorBlue:
				player.Weights = lastWinningWeigths
			case model.ColorGreen:
				player.Weights = model.NewRandomHintWeights()

				player.Weights.OpponentAccessibleTilesWeight = lastWinningWeigths.OpponentAccessibleTilesWeight

				player.Weights.TargetDistanceWeight = lastWinningWeigths.TargetDistanceWeight
				player.Weights.MovableTileWeight = lastWinningWeigths.MovableTileWeight
			case model.ColorRed:
				player.Weights = model.NewRandomHintWeights()

				player.Weights.RemainingTileShapeWeight = lastWinningWeigths.RemainingTileShapeWeight

				player.Weights.PlayerDistanceWeight = lastWinningWeigths.PlayerDistanceWeight
				player.Weights.TileShapeWeight = lastWinningWeigths.TileShapeWeight
			default:
				player.Weights = model.NewRandomHintWeights()

				player.Weights.PlayerAccessibleTilesWeight = lastWinningWeigths.PlayerAccessibleTilesWeight

				player.Weights.FixedTileWeight = lastWinningWeigths.FixedTileWeight
				player.Weights.CloseEdgeWeight = lastWinningWeigths.CloseEdgeWeight
			}

			fmt.Printf("[Round #%d] Start Player[%d] Weights=%v\n", round, player.Color, player.Weights)
		}

		maxTurns := 100
		for board.State != model.GameStateEnd && maxTurns > 0 {

			_, placeTileHint := board.GetPlaceTileHint()
			board.RemainingTile.Rotation = placeTileHint.Rotation

			if err := board.InsertTileAt(placeTileHint.Direction, placeTileHint.Index); err != nil {
				fmt.Printf("[Round #%d] Encountered place tile error: %v", round, err)
				continue outer
			}

			movePawnHint, err := board.GetMovePawnHint()
			if err != nil {
				fmt.Printf("[Round #%d] Encountered move pawn hint error: %v", round, err)
				continue outer
			}

			if err := board.MoveCurrentPlayerTo(movePawnHint.Line, movePawnHint.Row); err != nil {
				fmt.Printf("[Round #%d] Encountered move pawn error: %v", round, err)
				continue outer
			}

			for _, player := range board.Players {
				if len(player.Targets) == 0 {
					lastWinningWeigths = player.Weights
					fmt.Printf("[Round #%d] Winner = %v\n", round, lastWinningWeigths)
					continue outer
				}
			}

			maxTurns--
		}

		var (
			bestScore = -1
		)
		for _, player := range board.Players {

			fmt.Printf("[Round #%d] Draw Player[%d] Score=%v Weights=%v\n", round, player.Color, player.Score, player.Weights)
			if player.Score > 3 && bestScore < player.Score {
				bestScore = player.Score
				lastWinningWeigths = player.Weights
			}
		}
		fmt.Printf("[Round #%d] Draw, Best = %v\n", round, lastWinningWeigths)
	}

	fmt.Printf("Best weights = %v\n", lastWinningWeigths)
}
