package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type placeTileHintRequestBody struct {
	Board *model.Board `json:"board"`
}

type placeTileHint struct {
	InsertDirection       model.Direction `json:"insertDirection"`
	InsertIndex           int             `json:"insertIndex"`
	RemainingTileRotation model.Rotation  `json:"remainingTileRotation"`
}

type placeTileHintResponse struct {
	Hint    *placeTileHint `json:"hint"`
	Actions []*Action      `json:"actions"`
}

func placeTileHintHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /place-tile-hint', expected 'POST /place-tile-hint'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody placeTileHintRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/place-tile-hint' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
		return
	}

	_, hint := requestBody.Board.GetPlaceTileHint()
	if hint == nil {
		writeJsonResponse(w, http.StatusOK, &placeTileHintResponse{
			Hint:    nil,
			Actions: []*Action{},
		})
		return
	}

	writeJsonResponse(w, http.StatusOK, &placeTileHintResponse{
		Hint: &placeTileHint{
			InsertDirection:       hint.Direction,
			InsertIndex:           hint.Index,
			RemainingTileRotation: hint.Rotation,
		},
		Actions: []*Action{
			newRotateRemainingAction("", hint.Rotation),
		},
	})
}
