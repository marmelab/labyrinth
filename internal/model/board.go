package model

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	IShapedPercentage             = 0.26
	TShapedPercentage             = 0.36
	VShapedWithTreasurePercentage = 0.11
)

var (
	randomGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

	ErrEvenRow       = errors.New("row must be odd")
	ErrInvalidAction = errors.New("this action is not allowed in this state")
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

// TileExit represents the possible exits that are available for the tile.
type TileExit int

const (
	TileExitTop TileExit = iota
	TileExitRight
	TileExitBottom
	TileExitLeft
)

func (t TileExit) Opposite() TileExit {
	switch t {
	case TileExitTop:
		return TileExitBottom
	case TileExitRight:
		return TileExitLeft
	case TileExitBottom:
		return TileExitTop
	default:
		return TileExitRight
	}
}

func (t TileExit) ExitCoordinate(line, row int) *Coordinate {
	switch t {
	case TileExitTop:
		return &Coordinate{Line: line - 1, Row: row}
	case TileExitRight:
		return &Coordinate{Line: line, Row: row + 1}
	case TileExitBottom:
		return &Coordinate{Line: line + 1, Row: row}
	default:
		return &Coordinate{Line: line, Row: row - 1}
	}
}

type TileExits []TileExit

type ShapeRotationExit map[Rotation]TileExits
type ShapeExit map[Shape]ShapeRotationExit

var (
	shapeExitMap = ShapeExit{
		ShapeI: {
			Rotation0:   {TileExitRight, TileExitLeft},
			Rotation90:  {TileExitTop, TileExitBottom},
			Rotation180: {TileExitRight, TileExitLeft},
			Rotation270: {TileExitTop, TileExitBottom},
		},
		ShapeT: {
			Rotation0:   {TileExitTop, TileExitRight, TileExitLeft},
			Rotation90:  {TileExitTop, TileExitRight, TileExitBottom},
			Rotation180: {TileExitRight, TileExitBottom, TileExitLeft},
			Rotation270: {TileExitTop, TileExitBottom, TileExitLeft},
		},
		ShapeV: {
			Rotation0:   {TileExitBottom, TileExitLeft},
			Rotation90:  {TileExitTop, TileExitLeft},
			Rotation180: {TileExitTop, TileExitRight},
			Rotation270: {TileExitRight, TileExitBottom},
		},
	}
)

func (t TileExits) Contains(target TileExit) bool {
	for _, exit := range t {
		if exit == target {
			return true
		}
	}
	return false
}

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
}

func (b Board) validatePlaceTile(index int) error {
	if b.State != GameStatePlaceTile {
		return ErrInvalidAction
	}

	if (index & 1) != 1 {
		return ErrEvenRow
	}

	return nil
}

func (b *Board) InsertTileTopAt(row int) error {
	if err := b.validatePlaceTile(row); err != nil {
		return err
	}

	var current = b.RemainingTile
	for line := 0; line < b.Size(); line++ {
		b.Tiles[line][row], current = current, b.Tiles[line][row]

	}

	for _, player := range b.Players {
		if player.Position.Row == row {
			player.Position.Line = (player.Position.Line + 1) % b.Size()
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	return nil
}

func (b *Board) InsertTileRightAt(line int) error {
	if err := b.validatePlaceTile(line); err != nil {
		return err
	}

	var current = b.RemainingTile
	for row := b.Size() - 1; row >= 0; row-- {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	for _, player := range b.Players {
		if player.Position.Line == line {
			player.Position.Row = player.Position.Row - 1
			if player.Position.Row < 0 {
				player.Position.Row = b.Size() - 1
			}
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	return nil
}

func (b *Board) InsertTileBottomAt(row int) error {
	if err := b.validatePlaceTile(row); err != nil {
		return err
	}

	var current = b.RemainingTile
	for line := b.Size() - 1; line >= 0; line-- {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	for _, player := range b.Players {
		if player.Position.Row == row {
			player.Position.Line = player.Position.Line - 1
			if player.Position.Line < 0 {
				player.Position.Line = b.Size() - 1
			}
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	return nil
}

func (b *Board) InsertTileLeftAt(line int) error {
	if err := b.validatePlaceTile(line); err != nil {
		return err
	}

	var current = b.RemainingTile
	for row := 0; row < b.Size(); row++ {
		b.Tiles[line][row], current = current, b.Tiles[line][row]
	}

	for _, player := range b.Players {
		if player.Position.Line == line {
			player.Position.Row = (player.Position.Row + 1) % b.Size()
		}
	}

	b.RemainingTile = current
	b.State = GameStateMovePawn
	return nil
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

	if line >= b.Size() {
		return ErrInvalidAction
	}

	if row >= b.Size() {
		return ErrInvalidAction
	}

	if !b.GetAccessibleTiles().Contains(&Coordinate{line, row}) {
		return ErrInvalidAction
	}

	currentPlayer := b.CurrentPlayer()
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

// CurrentPlayer returns the current player.
func (b Board) CurrentPlayer() *Player {
	return b.Players[b.RemainingPlayers[b.RemainingPlayerIndex]]
}

// getAccessibleNeighbors returns the tile neighbors that can be accessed.
func (b Board) getAccessibleNeighbors(line, row int) Coordinates {
	var (
		coordinates = make(Coordinates, 0, 4)
		lastIndex   = b.Size() - 1
		exits       = b.Tiles[line][row].Exits()
	)

	for _, exit := range exits {
		exitTarget := exit.ExitCoordinate(line, row)

		if exitTarget.Line < 0 || exitTarget.Line > lastIndex || exitTarget.Row < 0 || exitTarget.Row > lastIndex {
			continue
		}

		targetTile := b.Tiles[exitTarget.Line][exitTarget.Row]
		if !targetTile.Exits().Contains(exit.Opposite()) {
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
func (b Board) GetAccessibleTiles() Coordinates {
	return b.getAccessibleTilesForCoordinate(b.CurrentPlayer().Position)
}

// Size returns the board size in tiles.
func (b Board) Size() int {
	return len(b.Tiles)
}

// BoardTile represents a tile that is placed on a board with a given rotation.
type BoardTile struct {

	// Tile is the underlying tile.
	Tile *Tile `json:"tile"`

	// Rotation is the tile rotation.
	Rotation Rotation `json:"rotation"`
}

// Exists return sthe possible exits for that tile as a bitmask.
func (bt BoardTile) Exits() TileExits {
	return shapeExitMap[bt.Tile.Shape][bt.Rotation]
}

// Rotation represents a tile rotation on a board.
type Rotation int

const (
	Rotation0 Rotation = 0

	Rotation90 Rotation = 90

	Rotation180 Rotation = 180

	Rotation270 Rotation = 270
)

// generateTiles generates tile list for the given board size.
// It will only generate size*size - 3 cards, since the tiles on each corner is
// predefined (fixed V-shaped).
func generateTiles(size int) (tiles []*Tile, treasures []Treasure) {
	var (
		tileCount = size*size + 1

		// We need to generate 4 less tiles as the corners are V tiles
		generatedTiles               = tileCount - 4
		tShapedThreshold             = int(math.Round(TShapedPercentage * float64(tileCount)))
		vShapedWithTreasureThreshold = tShapedThreshold + int(math.Round(VShapedWithTreasurePercentage*float64(tileCount)))
		iShapedThreshold             = vShapedWithTreasureThreshold + int(math.Round(IShapedPercentage*float64(tileCount)))
	)

	tiles = make([]*Tile, 0, generatedTiles)
	treasures = make([]Treasure, 0, vShapedWithTreasureThreshold)

	var (
		appendTileWithTreasure = func(shape Shape, i int) {
			treasure := 'A' + Treasure(i)
			tiles = append(tiles, &Tile{
				Shape:    shape,
				Treasure: treasure,
			})
			treasures = append(treasures, treasure)
		}
		appendTileWithoutTreasure = func(shape Shape) {
			tiles = append(tiles, &Tile{
				Shape:    shape,
				Treasure: NoTreasure,
			})
		}
	)

	for i := 0; i < generatedTiles; i++ {
		if i < tShapedThreshold {
			appendTileWithTreasure(ShapeT, i)
		} else if i < vShapedWithTreasureThreshold {
			appendTileWithTreasure(ShapeV, i)
		} else if i < iShapedThreshold {
			appendTileWithoutTreasure(ShapeI)
		} else {
			appendTileWithoutTreasure(ShapeV)
		}
	}

	return tiles, treasures
}

func randomRotation() Rotation {
	switch randomGenerator.Int63n(4) {
	case 0:
		return Rotation0
	case 1:
		return Rotation90
	case 2:
		return Rotation90
	default:
		return Rotation270
	}
}

// NewBoard returns a board for the given size.
func NewBoard(size, playerCount int) (*Board, error) {
	if size != 3 && size != 7 {
		return nil, fmt.Errorf("the board size must be either 3 or 7, got: %d", size)
	}

	if playerCount < 1 || playerCount > 4 {
		return nil, fmt.Errorf("the number of players must be between 1 and 4 included, got: %d", playerCount)
	}

	var (
		tiles, treasures = generateTiles(size)
		treasureCount    = len(treasures)
		targetPerPlayer  = treasureCount / playerCount
	)

	randomGenerator.Shuffle(len(tiles), func(i, j int) {
		tiles[i], tiles[j] = tiles[j], tiles[i]
	})

	randomGenerator.Shuffle(len(treasures), func(i, j int) {
		treasures[i], treasures[j] = treasures[j], treasures[i]
	})

	board := &Board{
		Tiles:                make([][]*BoardTile, size),
		Players:              players[:playerCount],
		RemainingPlayers:     remainingPlayers[:playerCount],
		RemainingPlayerIndex: 0,
		State:                GameStatePlaceTile,
	}

	for i, player := range board.Players {
		treasureOffset := i * targetPerPlayer

		if player.Position.Line == -1 {
			player.Position.Line = size - 1
		}

		if player.Position.Row == -1 {
			player.Position.Row = size - 1
		}

		player.Targets = treasures[treasureOffset : treasureOffset+targetPerPlayer]
	}

	// The tile index is required here to track placed tiles on the board.
	// This is due to the fact that each corner has a predefined V-shaped fixed
	// tile.
	tileIndex := 0
	for line := 0; line < size; line++ {
		board.Tiles[line] = make([]*BoardTile, size)

		for row := 0; row < size; row++ {
			if line == 0 && row == 0 {
				board.Tiles[line][row] = &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation270,
				}
			} else if line == 0 && row == size-1 {
				board.Tiles[line][row] = &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation0,
				}
			} else if line == size-1 && row == 0 {
				board.Tiles[line][row] = &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation180,
				}
			} else if line == size-1 && row == size-1 {
				board.Tiles[line][row] = &BoardTile{
					Tile: &Tile{
						Shape:    ShapeV,
						Treasure: NoTreasure,
					},
					Rotation: Rotation90,
				}
			} else {
				board.Tiles[line][row] = &BoardTile{
					Tile:     tiles[tileIndex],
					Rotation: randomRotation(),
				}
				tileIndex++
			}
		}
	}

	board.RemainingTile = &BoardTile{
		Tile:     tiles[len(tiles)-1],
		Rotation: Rotation0,
	}

	return board, nil
}
