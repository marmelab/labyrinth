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

function createTileHandlerFactory(
  rotateRemainingTileHandler?: RotateRemainingTileHandler,
  onInsertTile?: InsertTileHandler
): HandlerFactory {
  if (rotateRemainingTileHandler) {
    return () => rotateRemainingTileHandler;
  }

  // First Map is (line => row)
  // Second Map is (row => listener)
  const listeners = new Map<number, Map<number, Handler>>();

  const insertableIndexes = [1, 3, 5];

  if (onInsertTile) {
    listeners.set(
      0,
      new Map(
        insertableIndexes.map((index) => [
          index,
          onInsertTile?.bind(null, Direction.Top, index),
        ])
      )
    );

    insertableIndexes.forEach((line) =>
      listeners.set(
        line,
        new Map([
          [0, onInsertTile?.bind(null, Direction.Left, line)],
          [6, onInsertTile?.bind(null, Direction.Right, line)],
        ])
      )
    );

    listeners.set(
      6,
      new Map(
        insertableIndexes.map((index) => [
          index,
          onInsertTile?.bind(null, Direction.Bottom, index),
        ])
      )
    );
  }

  return (line: number, row: number): Handler => {
    return listeners.get(line)?.get(row);
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
  const onClickFactory = createTileHandlerFactory(
    onRotateRemainingTile,
    onInsertTile
  );

  const onClick = onClickFactory(line, row);

  return (
    <button
      className={`tile tile--shape-${shape} tile--rotation-${rotation}`}
      disabled={!onClick}
      onClick={onClick}
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};

export default TileView;
