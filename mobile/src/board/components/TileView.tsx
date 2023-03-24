import {
  type BoardTile,
  Direction,
  GameState,
  TileInsertion,
} from "../BoardTypes";
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
  hint: boolean;
  lastInsertion: TileInsertion | null;
  playerTarget?: string;
  onInsertTile: InsertTileHandler;
  onMovePlayer: MovePlayerHandler;
  isAccessible: boolean;
}

const oppositeDirections = {
  [Direction.Top]: Direction.Bottom,
  [Direction.Right]: Direction.Left,
  [Direction.Bottom]: Direction.Top,
  [Direction.Left]: Direction.Right,
};

export const TileView = ({
  boardTile,
  canPlay,
  gameState,
  coordinates,
  hint,
  lastInsertion,
  playerTarget,
  onInsertTile,
  onMovePlayer,
  isAccessible,
}: TileViewProps) => {
  if (gameState == GameState.PlaceTileAnimate) {
    return (
      <Tile
        disabled={!canPlay}
        animate
        boardTile={boardTile}
        playerTarget={playerTarget}
      />
    );
  }

  if (!canPlay || gameState == GameState.End) {
    return <Tile disabled boardTile={boardTile} playerTarget={playerTarget} />;
  }

  if (gameState == GameState.PlaceTile) {
    let direction = placeTileDirection
      .get(coordinates.line)
      ?.get(coordinates.row);

    if (
      direction &&
      !(
        oppositeDirections[direction[0]] == lastInsertion?.direction &&
        direction[1] == lastInsertion?.index
      )
    ) {
      return (
        <Tile
          boardTile={boardTile}
          playerTarget={playerTarget}
          onClick={onInsertTile.bind(null, ...direction)}
          hint={hint}
        />
      );
    }
    return <Tile disabled boardTile={boardTile} playerTarget={playerTarget} />;
  }

  return (
    <Tile
      disabled={!isAccessible}
      boardTile={boardTile}
      playerTarget={playerTarget}
      onClick={onMovePlayer.bind(null, coordinates.line, coordinates.row)}
      hint={hint}
    />
  );
};

export { RemainingTileView } from "./RemainingTileView";
