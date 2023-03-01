package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type rotateRemainingRequestBody struct {
	Board    *model.Board `json:"board"`
	Rotation string       `json:"rotation"`
}

func rotateRemainingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /rotate-remaining', expected 'POST /rotate-remaining'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody rotateRemainingRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/rotate-remaining' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
	}

	switch requestBody.Rotation {
	case "CLOCKWISE":
		requestBody.Board.RotateRemainingTileClockwise()
	case "ANTICLOCKWISE":
		requestBody.Board.RotateRemainingTileAntiClockwise()
	default:
		log.Printf("POST '/rotate-remaining' - Unsupported direction: %v", requestBody.Rotation)
		http.Error(w, fmt.Sprintf("unsupported direction: %v", requestBody.Rotation), http.StatusInternalServerError)
	}

	log.Println(requestBody.Board.RemainingTile)

	writeJsonResponse(w, http.StatusOK, requestBody.Board)
}
