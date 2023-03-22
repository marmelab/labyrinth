package model

import (
	"errors"
	"math/rand"
	"time"

	"github.com/RyanCarrier/dijkstra"
)

const (
	IShapedPercentage             = 0.26
	TShapedPercentage             = 0.36
	VShapedWithTreasurePercentage = 0.11
)

var (
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	ErrEvenRow              = errors.New("row must be odd")
	ErrInvalidAction        = errors.New("this action is not allowed in this state")
	ErrUnsupportedDirection = errors.New("unsupported direction")
)

type GameState int

const (
	GameStatePlaceTile GameState = iota
	GameStateMovePawn
	GameStateEnd
)

var (
	players = []*Player{
		{
			Color: ColorBlue,
			Position: &Coordinate{
				Line: 0,
				Row:  0,
			},
			Score: 0,
		},
		{
			Color: ColorGreen,
			Position: &Coordinate{
				Line: -1,
				Row:  -1,
			},
			Score: 0,
		},
		{
			Color: ColorRed,
			Position: &Coordinate{
				Line: 0,
				Row:  -1,
			},
			Score: 0,
		},
		{
			Color: ColorYellow,
			Position: &Coordinate{
				Line: -1,
				Row:  0,
			},
			Score: 0,
		},
	}
	remainingPlayers = []int{0, 1, 2, 3}
)

// Coordinate is a coordinate on the board.
type Coordinate struct {
	Line int `json:"line"`
	Row  int `json:"row"`
}

// Coordinates is set of cordinates.
type Coordinates []*Coordinate

// Contans returns whether the given Coordinate is present in the coordinate
// array.
func (c Coordinates) Contains(target *Coordinate) bool {
	for _, coordinate := range c {
		if coordinate.Line == target.Line && coordinate.Row == target.Row {
			return true
		}
	}
	return false
}

// Board represents the game board.
type Board struct {

	// Tiles are the tiles that are placed on a board.
	Tiles [][]*BoardTile `json:"tiles"`

	// RemainingTile is the tile that was not placed on the board.
	RemainingTile *BoardTile `json:"remainingTile"`

	// Players holds the players that are part of the game.
	Players []*Player `json:"players"`

	// RemainingPlayers holds the players that did not got all their targets yet.
	RemainingPlayers []int `json:"remainingPlayers"`

	// RemainingPlayerIndex is the index of the current player in the remaining
	// players array.
	RemainingPlayerIndex int `json:"currentPlayerIndex"`

	// GameState is the current game state
	State GameState `json:"gameState"`

	// LastInsertion is the last tile insertion.
	LastInsertion *TileInsertion `json:"lastInsertion"`
}

func (b Board) validatePlaceTile(direction Direction, index int) error {
	if b.State != GameStatePlaceTile {
		return ErrInvalidAction
	}

	if (index & 1) != 1 {
		return ErrEvenRow
	}

	if b.LastInsertion != nil && b.LastInsertion.isOppositeTo(direction, index) {
		return ErrInvalidAction
	}

	return nil
}

func (b *Board) InsertTileTopAt(row int) error {
	if err := b.validatePlaceTile(DirectionTop, row); err != nil {
		return err
	}

	var current = b.RemainingTile
	for line := 0; line < b.GetSize(); line++ {
		b.Tiles[line][row], current = current, b.Tiles[line][row]

	}

	for _, player := range b.Players {
		if player.Position.Row == row {
			player.Position.Line = (player.Position.Line + 1) % b.GetSize()
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	b.LastInsertion = &TileInsertion{
		Direction: DirectionTop,
		Index:     row,
	}
	return nil
}

func (b *Board) InsertTileRightAt(line int) error {
	if err := b.validatePlaceTile(DirectionRight, line); err != nil {
		return err
	}

	var current = b.RemainingTile
	for row := b.GetSize() - 1; row >= 0; row-- {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	for _, player := range b.Players {
		if player.Position.Line == line {
			player.Position.Row = player.Position.Row - 1
			if player.Position.Row < 0 {
				player.Position.Row = b.GetSize() - 1
			}
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	b.LastInsertion = &TileInsertion{
		Direction: DirectionRight,
		Index:     line,
	}
	return nil
}

func (b *Board) InsertTileBottomAt(row int) error {
	if err := b.validatePlaceTile(DirectionBottom, row); err != nil {
		return err
	}

	var current = b.RemainingTile
	for line := b.GetSize() - 1; line >= 0; line-- {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	for _, player := range b.Players {
		if player.Position.Row == row {
			player.Position.Line = player.Position.Line - 1
			if player.Position.Line < 0 {
				player.Position.Line = b.GetSize() - 1
			}
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	b.LastInsertion = &TileInsertion{
		Direction: DirectionBottom,
		Index:     row,
	}
	return nil
}

func (b *Board) InsertTileLeftAt(line int) error {
	if err := b.validatePlaceTile(DirectionLeft, line); err != nil {
		return err
	}

	var current = b.RemainingTile
	for row := 0; row < b.GetSize(); row++ {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	for _, player := range b.Players {
		if player.Position.Line == line {
			player.Position.Row = (player.Position.Row + 1) % b.GetSize()
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	b.LastInsertion = &TileInsertion{
		Direction: DirectionLeft,
		Index:     line,
	}
	return nil
}

func (b *Board) InsertTileAt(direction Direction, index int) error {
	switch direction {
	case DirectionTop:
		return b.InsertTileTopAt(index)
	case DirectionRight:
		return b.InsertTileRightAt(index)
	case DirectionBottom:
		return b.InsertTileBottomAt(index)
	case DirectionLeft:
		return b.InsertTileLeftAt(index)
	}

	return ErrUnsupportedDirection
}

func (b *Board) RotateRemainingTileClockwise() {
	switch b.RemainingTile.Rotation {
	case Rotation0:
		b.RemainingTile.Rotation = Rotation90
	case Rotation90:
		b.RemainingTile.Rotation = Rotation180
	case Rotation180:
		b.RemainingTile.Rotation = Rotation270
	case Rotation270:
		b.RemainingTile.Rotation = Rotation0
	}
}

func (b *Board) RotateRemainingTileAntiClockwise() {
	switch b.RemainingTile.Rotation {
	case Rotation0:
		b.RemainingTile.Rotation = Rotation270
	case Rotation90:
		b.RemainingTile.Rotation = Rotation0
	case Rotation180:
		b.RemainingTile.Rotation = Rotation90
	case Rotation270:
		b.RemainingTile.Rotation = Rotation180
	}
}

func (b *Board) MoveCurrentPlayerTo(line, row int) error {
	if b.State != GameStateMovePawn {
		return ErrInvalidAction
	}

	if line >= b.GetSize() {
		return ErrInvalidAction
	}

	if row >= b.GetSize() {
		return ErrInvalidAction
	}

	accessibleTiles, isShortestPath := b.GetAccessibleTiles()
	if targetTile := accessibleTiles[len(accessibleTiles)-1]; isShortestPath && (targetTile.Line != line || targetTile.Row != row) {
		return ErrInvalidAction
	} else if !accessibleTiles.Contains(&Coordinate{line, row}) {
		return ErrInvalidAction
	}

	currentPlayer := b.GetCurrentPlayer()
	currentPlayer.Position.Line = line
	currentPlayer.Position.Row = row

	currentTile := b.Tiles[line][row]
	if currentTile.Tile.Treasure == currentPlayer.Targets[0] {
		currentPlayer.Targets = currentPlayer.Targets[1:]
		currentPlayer.Score = currentPlayer.Score + 1
		currentTile.Tile.Treasure = NoTreasure
	}

	if len(currentPlayer.Targets) == 0 {

		// This removes the current player from the remaining player array
		b.RemainingPlayers = append(
			b.RemainingPlayers[:b.RemainingPlayerIndex],
			b.RemainingPlayers[b.RemainingPlayerIndex+1:]...)

		if b.RemainingPlayerIndex >= len(b.RemainingPlayers) {
			b.RemainingPlayerIndex = 0
		}
	} else {
		// We advance the remaining player index to the next player.
		b.RemainingPlayerIndex = (b.RemainingPlayerIndex + 1) % len(b.RemainingPlayers)
	}

	// TODO: Do not force last player to end the game.
	if len(b.RemainingPlayers) == 0 {
		b.State = GameStateEnd
	} else {
		b.State = GameStatePlaceTile
	}
	return nil
}

func (b Board) GetCurrentPlayer() *Player {
	if len(b.RemainingPlayers) == 0 {
		return b.Players[0]
	}
	return b.Players[b.RemainingPlayers[b.RemainingPlayerIndex]]
}

// getAccessibleNeighbors returns the tile neighbors that can be accessed.
func (b Board) getAccessibleNeighbors(line, row int) Coordinates {
	var (
		coordinates = make(Coordinates, 0, 4)
		lastIndex   = b.GetSize() - 1
		exits       = b.Tiles[line][row].GetExits()
	)

	for _, exit := range exits {
		exitTarget := exit.ExitCoordinate(line, row)

		if exitTarget.Line < 0 || exitTarget.Line > lastIndex || exitTarget.Row < 0 || exitTarget.Row > lastIndex {
			continue
		}

		targetTile := b.Tiles[exitTarget.Line][exitTarget.Row]
		if !targetTile.GetExits().Contains(exit.Opposite()) {
			continue
		}

		coordinates = append(coordinates, exitTarget)
	}

	return coordinates
}

// getAccessibleTilesForCoordinate returns the available tiles from the given coordinates.
func (b Board) getAccessibleTilesForCoordinate(coordinate *Coordinate) Coordinates {
	var (
		results = make(Coordinates, 0)
		queue   = append(make(Coordinates, 0), coordinate)
	)

	for len(queue) > 0 {
		currentTile := queue[0]
		queue = queue[1:]
		if results.Contains(currentTile) {
			continue
		}

		results = append(results, currentTile)
		queue = append(queue, b.getAccessibleNeighbors(currentTile.Line, currentTile.Row)...)
	}

	return results
}

// GetAccessibleTiles returns the tiles that are accessible by the current
// player.
func (b Board) GetAccessibleTiles() (tiles Coordinates, isShortestPath bool) {
	accessibleTiles := b.getAccessibleTilesForCoordinate(b.GetCurrentPlayer().Position)

	currentPlayer := b.GetCurrentPlayer()
	if len(currentPlayer.Targets) == 0 {
		return accessibleTiles, false
	}

	for _, coordinate := range accessibleTiles {
		currentTile := b.Tiles[coordinate.Line][coordinate.Row]
		if currentTile.Tile.Treasure == currentPlayer.Targets[0] {
			return b.getShortestPath(), true
		}
	}

	return accessibleTiles, false
}

// GetSize returns the board size in tiles.
func (b Board) GetSize() int {
	return len(b.Tiles)
}

// vertextId returns  a unique identifier for the tile based on its line and row.
func (b Board) vertextId(line, row int) int {
	return line*b.GetSize() + row
}

func (b Board) buildGraph() (graph *dijkstra.Graph, getCoordinatesByVertex func(id int) *Coordinate) {
	graph = dijkstra.NewGraph()

	var (
		size                = b.GetSize()
		coordinatesByVertex = make(map[int]*Coordinate)
	)

	graph.AddVertex(size*size + 1)
	for line := 0; line < size; line++ {
		for row := 0; row < size; row++ {

			vertexId := b.vertextId(line, row)
			graph.AddVertex(vertexId)

			coordinatesByVertex[vertexId] = &Coordinate{
				Line: line,
				Row:  row,
			}

		}
	}

	for line := 0; line < size; line++ {
		for row := 0; row < size; row++ {
			neighbors := b.getAccessibleNeighbors(line, row)
			for _, neighbor := range neighbors {
				graph.AddArc(
					b.vertextId(line, row),
					b.vertextId(neighbor.Line, neighbor.Row),
					1)
			}
		}
	}

	return graph, func(id int) *Coordinate {
		return coordinatesByVertex[id]
	}
}

func (b Board) getCurrentTargetIndex() (int, int) {
	var (
		size          = b.GetSize()
		currentPlayer = b.GetCurrentPlayer()
	)

	for line := 0; line < size; line++ {
		for row := 0; row < size; row++ {
			if b.Tiles[line][row].Tile.Treasure == currentPlayer.Targets[0] {
				return line, row
			}
		}
	}

	return size, 0
}

func (b Board) GetCurrentTargetCoordinate() *Coordinate {
	line, row := b.getCurrentTargetIndex()
	if line >= len(b.Tiles) {
		return nil
	}

	return &Coordinate{
		Line: line,
		Row:  row,
	}
}

func (b Board) getShortestPath() Coordinates {
	var (
		size = b.GetSize()

		currentPlayer = b.GetCurrentPlayer()

		graph, getCoordinatesByVertex = b.buildGraph()

		sourceVertex = b.vertextId(currentPlayer.Position.Line, currentPlayer.Position.Row)

		targetLine, targetRow = b.getCurrentTargetIndex()

		// targetVertex is the target node
		targetVertex = targetLine*size + targetRow
	)

	if sourceVertex == targetVertex {
		return Coordinates{
			currentPlayer.Position,
		}
	}

	best, err := graph.Shortest(sourceVertex, targetVertex)
	if err != nil { // No path was found.
		return nil
	}

	path := make(Coordinates, 0, len(best.Path))
	for _, vertex := range best.Path {
		path = append(path, getCoordinatesByVertex(vertex))
	}

	return path
}
