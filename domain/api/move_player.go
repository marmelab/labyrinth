package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type movePlayerRequestBody struct {
	Board *model.Board `json:"board"`
	Line  int          `json:"line"`
	Row   int          `json:"row"`
}

func movePlayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /move-player', expected 'POST /move-player'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody movePlayerRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/move-player' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
	}

	path, err := requestBody.Board.MoveCurrentPlayerTo(requestBody.Line, requestBody.Row)
	actions := make([]*Action, 0, 1)
	if err == nil {
		actions = append(actions,
			newMovePawnAction(requestBody.Line, requestBody.Row, path))
	}

	writeJsonResponse(w, http.StatusOK, &BoardResponse{
		Board:   requestBody.Board,
		Actions: actions,
	})
}
