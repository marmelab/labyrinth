import type { FunctionComponent, MouseEventHandler, ReactElement } from "react";

import { BoardTile, Tile } from "../BoardTypes";

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

interface TileProps {
  boardTile: BoardTile;
  disabled?: boolean;
  onClick?: MouseEventHandler;
  children?: ReactElement | ReactElement[];
}

const TileView: FunctionComponent<TileProps> = ({
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

export default TileView;
