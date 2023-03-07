export enum GameState {
  PlaceTile,
  MovePawn,
  End,
}

export interface Board {
  id: number;
  state: BoardState;
}

export interface BoardState {
  tiles: Array<Array<BoardTile>>;
  remainingTile: BoardTile;
  players: Array<Player>;
  remainingPlayers: Array<number>;
  currentPlayerIndex: number;
  gameState: GameState;
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
  red,
  Yellow,
}

export interface Coordinate {
  line: number;
  row: number;
}

export interface Player {
  color: Color;
  position: Coordinate;
  targets: Array<string>;
  score: number;
}
