package model

type Direction string

const (
	DirectionTop    Direction = "TOP"
	DirectionRight  Direction = "RIGHT"
	DirectionBottom Direction = "BOTTOM"
	DirectionLeft   Direction = "LEFT"
)

var (
	oppositeDirection = map[Direction]Direction{
		DirectionTop:    DirectionBottom,
		DirectionRight:  DirectionLeft,
		DirectionBottom: DirectionTop,
		DirectionLeft:   DirectionRight,
	}
)

var (
	insertionDirections = []Direction{DirectionTop, DirectionRight, DirectionBottom, DirectionLeft}
	insertionIndexes    = []int{1, 3, 5}
)

type TileInsertion struct {
	Direction Direction `json:"direction"`
	Index     int       `json:"index"`
}

func (b *Board) getAvailableInsertions() []*TileInsertion {
	availableInsertions := make([]*TileInsertion, 0, 12)

	for _, direction := range insertionDirections {
		for _, index := range insertionIndexes {
			if b.LastInsertion != nil && b.LastInsertion.Direction == direction && b.LastInsertion.Index == index {
				continue
			}

			availableInsertions = append(availableInsertions, &TileInsertion{
				Direction: direction,
				Index:     index,
			})
		}
	}
	return availableInsertions
}
