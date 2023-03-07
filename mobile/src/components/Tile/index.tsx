import type { FunctionComponent, MouseEventHandler, ReactElement } from "react";
import { BoardTile } from "../../model/Board";
import type { Tile } from "../../model/Tile";

import "./index.css";

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

interface TileProps {
  /**
   *
   */
  boardTile: BoardTile;

  /**
   * Whether the user can click on the tile.
   */
  disabled?: boolean;

  /**
   *
   */
  onClick: MouseEventHandler;

  /**
   *
   */
  children?: ReactElement[];
}

/**
 *
 */
const Tile: FunctionComponent<TileProps> = ({
  boardTile: {
    tile: { treasure, shape },
    rotation,
  },
  disabled = false,
  onClick,
  children,
}: TileProps): ReactElement => {
  return (
    <button
      className={`tile tile--shape-${shape} tile--rotation-${rotation}`}
      disabled={disabled}
      onClick={onClick}
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};

export default Tile;
