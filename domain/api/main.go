package main

import (
	"log"
	"net/http"

	"github.com/marmelab/labyrinth/domain/internal/model"
)

func newHandler(w http.ResponseWriter, r *http.Request) {
	board, err := model.NewBoard(7, 1)
	if err != nil {
		log.Printf("Failed to initialize board: %v.", err)
		http.Error(w, "failed to initialize board", http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, board)
}

func main() {

	http.HandleFunc("/new", newHandler)

	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatalf("failed to listen: %v.", err)
	}
}
