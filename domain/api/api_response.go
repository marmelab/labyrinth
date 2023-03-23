package main

import "github.com/marmelab/labyrinth/domain/internal/model"

type ActionKind string

const (
	ActionKindRotateRemining ActionKind = "ROTATE_REMAINING"
	ActionKindPlaceTile      ActionKind = "PLACE_TILE"
	ActionKindMovePawn       ActionKind = "MOVE_PAWN"
)

type Action struct {
	Kind     ActionKind  `json:"kind"`
	Playload interface{} `json:"payload"`
}

type BoardResponse struct {
	Board   *model.Board `json:"board"`
	Actions []*Action    `json:"actions"`
}

func newRotateRemainingAction(direction RotationDirection, rotation model.Rotation) *Action {
	return &Action{
		Kind: ActionKindRotateRemining,
		Playload: &struct {
			Direction RotationDirection `json:"direction"`
			Rotation  model.Rotation    `json:"rotation"`
		}{
			Direction: direction,
			Rotation:  rotation,
		},
	}
}

func newPlaceTileAction(direction model.Direction, index int) *Action {
	return &Action{
		Kind: ActionKindPlaceTile,
		Playload: &struct {
			Direction model.Direction `json:"direction"`
			Index     int             `json:"index"`
		}{
			Direction: direction,
			Index:     index,
		},
	}
}

func newMovePawnAction(line, row int, path model.Coordinates) *Action {
	return &Action{
		Kind: ActionKindMovePawn,
		Playload: &struct {
			Line int               `json:"line"`
			Row  int               `json:"row"`
			Path model.Coordinates `json:"path"`
		}{
			Line: line,
			Row:  row,
			Path: path,
		},
	}
}
