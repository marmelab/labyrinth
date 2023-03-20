package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type getAccessibleTilesRequestBody struct {
	Board *model.Board `json:"board"`
}

type GetAccessibleTilesResponse struct {
	IsShortestPath bool              `json:"isShortestPath"`
	Coordinates    model.Coordinates `json:"coordinates"`
}

func getAccessibleTilesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /get-accessible-tiles', expected 'POST /get-accessible-tiles'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody getAccessibleTilesRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/get-accessible-tiles' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
		return
	}

	if requestBody.Board.State != model.GameStateMovePawn {
		log.Printf("POST '/get-accessible-tiles' - invalid state: %v", requestBody.Board.State)
		http.Error(w, "invalid state, expected MOVE_PAWN", http.StatusBadRequest)
		return
	}

	coordinates, isShortestPath := requestBody.Board.GetAccessibleTiles()
	writeJsonResponse(w, http.StatusOK, &GetAccessibleTilesResponse{
		IsShortestPath: isShortestPath,
		Coordinates:    coordinates,
	})
}
