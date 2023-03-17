import type { ReactElement, ReactNode } from "react";

import { type BoardTile, Direction, GameState } from "../BoardTypes";

import "./TileView.css";

interface TreasureMap {
  [key: string]: string;
}

export const treasures: TreasureMap = {
  "": " ",
  ".": " ",
  A: "💌",
  B: "💣",
  C: "🛍",
  D: "📿",
  E: "🔭",
  F: "💎",
  G: "💰",
  H: "📜",
  I: "🗿",
  J: "🏺",
  K: "🔫",
  L: "🛡",
  M: "💈",
  N: "🛎",
  O: "⌛",
  P: "🌡",
  Q: "⛱",
  R: "🎈",
  S: "🎎",
  T: "🎁",
  U: "🔮",
  V: "📷",
  W: "🕯",
  X: "🥦",
};

type Handler = (() => Promise<void>) | undefined;

export type RotateRemainingTileHandler = () => Promise<void>;

export type InsertTileHandler = (
  direction: Direction,
  index: number
) => Promise<void>;

export type MovePlayerHandler = (line: number, row: number) => Promise<void>;

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

const Tile = ({
  remainingTile,
  boardTile: {
    tile: { treasure, shape },
    rotation,
  },
  playerTarget,
  children,
  disabled,
  onClick,
}: {
  remainingTile: boolean;
  boardTile: BoardTile;
  children: ReactNode;
  playerTarget?: string;
  disabled?: boolean;
  onClick?: Handler;
}) => {
  return (
    <button
      className={`tile tile--shape-${shape} ${
        playerTarget == treasure ? "tile--target" : ""
      } ${remainingTile ? "tile--remaining" : ""}`}
      disabled={disabled}
      onClick={onClick}
      style={{ transform: `rotate(${rotation}deg)` }}
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};

interface TileProps {
  remainingTile?: boolean;
  boardTile: BoardTile;
  canPlay: boolean;
  gameState: GameState;
  coordinates?: {
    line: number;
    row: number;
  };
  playerTarget?: string;
  onRotateRemainingTile: RotateRemainingTileHandler;
  onInsertTile: InsertTileHandler;
  onMovePlayer: MovePlayerHandler;
  children?: ReactNode;
}

const TileView = ({
  remainingTile = false,
  boardTile,
  canPlay,
  gameState,
  coordinates,
  playerTarget,
  onRotateRemainingTile,
  onInsertTile,
  onMovePlayer,
  children,
}: TileProps) => {
  if (!canPlay || gameState == GameState.End) {
    return (
      <Tile
        disabled
        boardTile={boardTile}
        playerTarget={playerTarget}
        remainingTile={remainingTile}
      >
        {children}
      </Tile>
    );
  }

  if (!coordinates) {
    return (
      <Tile
        boardTile={boardTile}
        onClick={onRotateRemainingTile}
        playerTarget={playerTarget}
        remainingTile={remainingTile}
      >
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
          remainingTile={remainingTile}
        >
          {children}
        </Tile>
      );
    }
    return (
      <Tile
        disabled
        boardTile={boardTile}
        playerTarget={playerTarget}
        remainingTile={remainingTile}
      >
        {children}
      </Tile>
    );
  }

  return (
    <Tile
      boardTile={boardTile}
      playerTarget={playerTarget}
      onClick={onMovePlayer.bind(null, coordinates.line, coordinates.row)}
      remainingTile={remainingTile}
    >
      {children}
    </Tile>
  );
};

export default TileView;
