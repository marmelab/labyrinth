export enum GameState {
  PlaceTile,
  MovePawn,
  End,
}

export interface BoardListItem {
  id: number;
  remainingSeats: number;
}

export interface Board {
  id: number;
  remainingSeats: number;
  canJoin: boolean;
  players: Player[];
  state: BoardState;
  canPlay: boolean;
  gameState: GameState;
  currentPlayer?: Player | null;
  user?: Player | null;
}

export interface BoardState {
  tiles: BoardTile[][];
  remainingTile: BoardTile;
}

export interface BoardTile {
  tile: Tile;
  rotation: Rotation;
}

export enum Shape {
  ShapeI = 0,
  ShapeT = 1,
  ShapeV = 1,
}

export enum Rotation {
  Rotation0 = 0,
  Rotation90 = 90,
  Rotation180 = 180,
  Rotation270 = 270,
}

export interface Tile {
  treasure: string;
  shape: Shape;
}

export enum Color {
  Blue,
  Green,
  Red,
  Yellow,
}

export interface Player {
  name: string;
  color: Color;
  line: number;
  row: number;
  targets: string[];
  currentTarget: string;
  score: number;
  isCurrentPlayer: boolean;
  isUser: boolean;
}

export type BoardID = number | string;

export enum Direction {
  Top = "TOP",
  Right = "RIGHT",
  Bottom = "BOTTOM",
  Left = "LEFT",
}

export interface BoardRepository {
  list(page: number): Promise<BoardListItem[]>;
  getById(id: BoardID): Promise<Board>;
  rotateRemainingTile(id: BoardID): Promise<void>;
  insertTile(id: BoardID, direction: Direction, index: number): Promise<void>;
  joinBoard(id: BoardID): Promise<Board>;
}
