package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/rotate-remaining", rotateRemainingHandler)
	http.HandleFunc("/insert-tile", insertTileHandler)
	http.HandleFunc("/move-player", movePlayerHandler)
	http.HandleFunc("/place-tile-hint", placeTileHintHandler)
	http.HandleFunc("/get-accessible-tiles", getAccessibleTilesHandler)

	if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
		log.Fatalf("failed to listen: %v.", err)
	}
}
