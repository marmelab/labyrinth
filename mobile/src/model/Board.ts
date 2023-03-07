import { Player } from "./Player";
import { Rotation, Tile } from "./Tile";

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
};

export interface BoardTile {

    tile: Tile;

    rotation: Rotation;
}