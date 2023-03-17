import type { ReactNode } from "react";

import { type BoardTile, Direction, GameState } from "../BoardTypes";

import "./Tile.css";

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

export type Handler = (() => Promise<void>) | undefined;

export type RotateRemainingTileHandler = () => Promise<void>;

export type InsertTileHandler = (
  direction: Direction,
  index: number
) => Promise<void>;

export type MovePlayerHandler = (line: number, row: number) => Promise<void>;

export const Tile = ({
  animate = false,
  remainingTile = false,
  boardTile: {
    tile: { treasure, shape },
    rotation,
    top = 0,
    left = 0,
    opacity = 1,
  },
  playerTarget,
  children,
  disabled,
  onClick,
}: {
  animate?: boolean;
  remainingTile?: boolean;
  boardTile: BoardTile;
  children?: ReactNode;
  playerTarget?: string;
  disabled?: boolean;
  onClick?: Handler;
}) => {
  return (
    <button
      className={`tile tile--shape-${shape} ${
        playerTarget == treasure ? "tile--target" : ""
      } ${remainingTile ? "tile--remaining" : ""} ${
        animate ? "tile--animate" : ""
      }`}
      disabled={disabled}
      onClick={onClick}
      style={{ transform: `rotate(${rotation}deg)`, top, left, opacity }}
    >
      <div className={`tile__path`}></div>
      <div className="tile__treasure">{treasures[treasure]}</div>
      {children}
    </button>
  );
};
