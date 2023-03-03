package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type newRequestBody struct {
	PlayerCount int `json:"playerCount"`
}

// newHandler is in charge of "/new" routes
func newHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /new', expected 'POST /new'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody newRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/rotate-remaining' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
	}

	board, err := model.NewBoard(7, requestBody.PlayerCount)
	if err != nil {
		log.Printf("Failed to initialize board: %v.", err)
		http.Error(w, "failed to initialize board", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, http.StatusOK, board)
}
