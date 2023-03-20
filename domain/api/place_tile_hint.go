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
	}

	updatedBoard, hint := requestBody.Board.GetPlaceTileHint()
	if hint == nil {
		writeJsonResponse(w, http.StatusOK, &BoardResponse{
			Board:   requestBody.Board,
			Actions: []*Action{},
		})
		return
	}

	writeJsonResponse(w, http.StatusOK, &BoardResponse{
		Board: updatedBoard,
		Actions: []*Action{
			newRotateRemainingAction("", hint.Rotation),
			newPlaceTileAction(hint.Direction, hint.Index),
		},
	})
}
