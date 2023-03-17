import { type ReactNode } from "react";

import { type BoardTile, Direction, GameState } from "../BoardTypes";
import { type InsertTileHandler, type MovePlayerHandler, Tile } from "./Tile";

const insertableIndexes = [1, 3, 5];

// First Map is (line => row)
// Second Map is (row => [direction, index])
const placeTileDirection = new Map<number, Map<number, [Direction, number]>>([
  [
    0,
    new Map(insertableIndexes.map((index) => [index, [Direction.Top, index]])),
  ],
  [
    1,
    new Map([
      [0, [Direction.Left, 1]],
      [6, [Direction.Right, 1]],
    ]),
  ],
  [
    3,
    new Map([
      [0, [Direction.Left, 3]],
      [6, [Direction.Right, 3]],
    ]),
  ],
  [
    5,
    new Map([
      [0, [Direction.Left, 5]],
      [6, [Direction.Right, 5]],
    ]),
  ],
  [
    6,
    new Map(
      insertableIndexes.map((index) => [index, [Direction.Bottom, index]])
    ),
  ],
]);

interface TileViewProps {
  boardTile: BoardTile;
  canPlay: boolean;
  gameState: GameState;
  coordinates: {
    line: number;
    row: number;
  };
  playerTarget?: string;
  onInsertTile: InsertTileHandler;
  onMovePlayer: MovePlayerHandler;
  children?: ReactNode;
}

export const TileView = ({
  boardTile,
  canPlay,
  gameState,
  coordinates,
  playerTarget,
  onInsertTile,
  onMovePlayer,
  children,
}: TileViewProps) => {
  if (!canPlay || gameState == GameState.End) {
    return (
      <Tile disabled boardTile={boardTile} playerTarget={playerTarget}>
        {children}
      </Tile>
    );
  }

  if (gameState == GameState.Animating) {
    return (
      <Tile animate boardTile={boardTile} playerTarget={playerTarget}>
        {children}
      </Tile>
    );
  }

  if (gameState == GameState.PlaceTile) {
    let direction = placeTileDirection
      .get(coordinates.line)
      ?.get(coordinates.row);
    if (direction) {
      return (
        <Tile
          boardTile={boardTile}
          playerTarget={playerTarget}
          onClick={onInsertTile.bind(null, ...direction)}
        >
          {children}
        </Tile>
      );
    }
    return (
      <Tile disabled boardTile={boardTile} playerTarget={playerTarget}>
        {children}
      </Tile>
    );
  }

  return (
    <Tile
      boardTile={boardTile}
      playerTarget={playerTarget}
      onClick={onMovePlayer.bind(null, coordinates.line, coordinates.row)}
    >
      {children}
    </Tile>
  );
};

export { RemainingTileView } from "./RemainingTileView";
