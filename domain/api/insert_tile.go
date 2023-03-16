package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type Direction string

const (
	DirectionTop    Direction = "TOP"
	DirectionRight  Direction = "RIGHT"
	DirectionBottom Direction = "BOTTOM"
	DirectionLeft   Direction = "LEFT"
)

type insertTileRequestBody struct {
	Board     *model.Board `json:"board"`
	Direction Direction    `json:"direction"`
	Index     int          `json:"index"`
}

func insertTileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /insert-tile', expected 'POST /insert-tile'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	var requestBody insertTileRequestBody
	if err := parseJsonBody(r, &requestBody); err != nil {
		log.Printf("POST '/insert-tile' - Failed to decode body: %v", err)
		http.Error(w, "failed to decode body", http.StatusInternalServerError)
	}

	var err error = nil
	switch requestBody.Direction {
	case DirectionTop:
		err = requestBody.Board.InsertTileTopAt(requestBody.Index)
	case DirectionRight:
		err = requestBody.Board.InsertTileRightAt(requestBody.Index)
	case DirectionBottom:
		err = requestBody.Board.InsertTileBottomAt(requestBody.Index)
	case DirectionLeft:
		err = requestBody.Board.InsertTileLeftAt(requestBody.Index)
	default:
		log.Printf("POST '/insert-tile' - Unsupported direction: %v", requestBody.Direction)
		http.Error(w, fmt.Sprintf("unsupported direction: %v", requestBody.Direction), http.StatusInternalServerError)
	}

	actions := make([]*Action, 0, 1)
	if err == nil {
		actions = append(actions,
			newPlaceTileAction(requestBody.Direction, requestBody.Index),
			newGameStateChangeAction(requestBody.Board.State))
	}

	writeJsonResponse(w, http.StatusOK, &BoardResponse{
		Board:   requestBody.Board,
		Actions: actions,
	})
}
