package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

type insertTileRequestBody struct {
	Board     *model.Board `json:"board"`
	Direction string       `json:"direction"`
	Index     int          `json:"index"`
}

const (
	DirectionTop    = "TOP"
	DirectionRight  = "RIGHT"
	DirectionBottom = "BOTTOM"
	DirectionLeft   = "LEFT"
)

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

	log.Println(requestBody.Direction, requestBody.Index)

	switch requestBody.Direction {
	case DirectionTop:
		requestBody.Board.InsertTileTopAt(requestBody.Index)
	case DirectionRight:
		requestBody.Board.InsertTileRightAt(requestBody.Index)
	case DirectionBottom:
		requestBody.Board.InsertTileBottomAt(requestBody.Index)
	case DirectionLeft:
		requestBody.Board.InsertTileLeftAt(requestBody.Index)
	default:
		log.Printf("POST '/insert-tile' - Unsupported direction: %v", requestBody.Direction)
		http.Error(w, fmt.Sprintf("unsupported direction: %v", requestBody.Direction), http.StatusInternalServerError)
	}

	writeJsonResponse(w, http.StatusOK, requestBody.Board)
}
