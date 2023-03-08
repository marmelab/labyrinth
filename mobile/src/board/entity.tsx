export enum GameState {
  PlaceTile,
  MovePawn,
  End,
}

export interface Board {
  id: number;
  remainingSeats: number;
  players: Player[];
  state: BoardState;
  canPlay: boolean;
  gameState: GameState;
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
  red,
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
