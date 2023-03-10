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

const ClickableTileView = ({
  boardTile: {
    tile: { treasure, shape },
    rotation,
  },
  onClick,
  children,
}: {
  boardTile: BoardTile;
  onClick: Handler;
  children: ReactNode;
}) => {
  return (
    <button
      className={`tile tile--shape-${shape} tile--rotation-${rotation}`}
      onClick={onClick}
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};

const DisabledTileView = ({
  boardTile: {
    tile: { treasure, shape },
    rotation,
  },
  children,
}: {
  boardTile: BoardTile;
  children: ReactNode;
}) => {
  return (
    <button
      className={`tile tile--shape-${shape} tile--rotation-${rotation}`}
      disabled
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};

interface TileProps {
  boardTile: BoardTile;
  canPlay: boolean;
  gameState: GameState;
  coordinates?: {
    line: number;
    row: number;
  };
  onRotateRemainingTile: RotateRemainingTileHandler;
  onInsertTile: InsertTileHandler;
  onMovePlayer: MovePlayerHandler;
  children?: ReactNode;
}

const TileView = ({
  boardTile,
  canPlay,
  gameState,
  coordinates,
  onRotateRemainingTile,
  onInsertTile,
  onMovePlayer,
  children,
}: TileProps) => {
  if (!canPlay || gameState == GameState.End) {
    return (
      <DisabledTileView boardTile={boardTile}>{children}</DisabledTileView>
    );
  }

  if (!coordinates) {
    return (
      <ClickableTileView boardTile={boardTile} onClick={onRotateRemainingTile}>
        {children}
      </ClickableTileView>
    );
  }

  if (gameState == GameState.PlaceTile) {
    let direction = placeTileDirection
      .get(coordinates.line)
      ?.get(coordinates.row);
    if (direction) {
      return (
        <ClickableTileView
          boardTile={boardTile}
          onClick={onInsertTile.bind(null, ...direction)}
        >
          {children}
        </ClickableTileView>
      );
    }
    return (
      <DisabledTileView boardTile={boardTile}>{children}</DisabledTileView>
    );
  }

  return (
    <ClickableTileView
      boardTile={boardTile}
      onClick={onMovePlayer.bind(null, coordinates.line, coordinates.row)}
    >
      {children}
    </ClickableTileView>
  );
};

export default TileView;
