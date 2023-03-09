import type { FunctionComponent, ReactElement } from "react";

import { type BoardTile, Direction, GameState } from "../BoardTypes";

import "./TileView.css";

interface TreasureMap {
  [key: string]: string;
}

const treasures: TreasureMap = {
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
type HandlerFactory = (line: number, row: number) => Handler;

export type RotateRemainingTileHandler = () => Promise<void>;
export type InsertTileHandler = (
  direction: Direction,
  index: number
) => Promise<void>;

/**
 * This functions creates the click handlers for the tile view
 */
function createClickHandlerFactory(
  rotateRemainingTileHandler?: RotateRemainingTileHandler,
  onInsertTile?: InsertTileHandler
): HandlerFactory {
  // If rotateRemainingTileHandler is provided, then this is the remaining tile.
  if (rotateRemainingTileHandler) {
    return () => rotateRemainingTileHandler;
  }

  // First Map is (line => row)
  // Second Map is (row => listener)
  const handlers = new Map<number, Map<number, Handler>>();

  const insertableIndexes = [1, 3, 5];

  // If onInsertTile is provided, the player can place a tile
  if (onInsertTile) {
    // Top tile listeners
    handlers.set(
      0,
      new Map(
        insertableIndexes.map((index) => [
          index,
          onInsertTile?.bind(null, Direction.Top, index),
        ])
      )
    );

    // Lefta nd right tile listeners
    insertableIndexes.forEach((line) =>
      handlers.set(
        line,
        new Map([
          [0, onInsertTile?.bind(null, Direction.Left, line)],
          [6, onInsertTile?.bind(null, Direction.Right, line)],
        ])
      )
    );

    // Bottom tile listeners
    handlers.set(
      6,
      new Map(
        insertableIndexes.map((index) => [
          index,
          onInsertTile?.bind(null, Direction.Bottom, index),
        ])
      )
    );
  }

  /**
   * Get the handler for the given line and row.
   */
  return (line: number, row: number): Handler => {
    return handlers.get(line)?.get(row);
  };
}

interface TileProps {
  boardTile: BoardTile;
  line: number;
  row: number;
  disabled?: boolean;
  onRotateRemainingTile?: RotateRemainingTileHandler;
  onInsertTile?: InsertTileHandler;
  children?: ReactElement | ReactElement[];
}

const TileView: FunctionComponent<TileProps> = ({
  boardTile: {
    tile: { treasure, shape },
    rotation,
  },
  line,
  row,
  disabled = false,
  onRotateRemainingTile,
  onInsertTile,
  children,
}: TileProps): ReactElement => {
  const handleClickFactory = createClickHandlerFactory(
    onRotateRemainingTile,
    onInsertTile
  );

  const handleClick = handleClickFactory(line, row);

  return (
    <button
      className={`tile tile--shape-${shape} tile--rotation-${rotation}`}
      disabled={!handleClick}
      onClick={handleClick}
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};

export default TileView;
