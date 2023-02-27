package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

// newHandler is in charge of "/new" routes
func newHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Got '%v /new', expected 'POST /new'", r.Method)
		http.Error(w, fmt.Sprintf("unexpected HTTP method: %v", r.Method), http.StatusMethodNotAllowed)
		return
	}

	board, err := model.NewBoard(7, 1)
	if err != nil {
		log.Printf("Failed to initialize board: %v.", err)
		http.Error(w, "failed to initialize board", http.StatusInternalServerError)
		return
	}

	writeJsonResponse(w, http.StatusOK, board)
}

func main() {
	http.HandleFunc("/new", newHandler)

	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatalf("failed to listen: %v.", err)
	}
}
