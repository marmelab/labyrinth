package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type movePawnHintRequestBody struct {
	Board *model.Board `json:"board"`
}

type movePawnHintResponse struct {
	Hint *model.Coordinate `json:"hint"`
}

func movePawnHintHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /move-pawn-hint', expected 'POST /move-pawn-hint'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody movePawnHintRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/move-pawn-hint' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
		return
	}

	hint, err := requestBody.Board.GetMovePawnHint()
	if err != nil {
		writeJsonResponse(w, http.StatusOK, &movePawnHintResponse{
			Hint: nil,
		})
		return
	}

	writeJsonResponse(w, http.StatusOK, &movePawnHintResponse{
		Hint: hint,
	})
}
