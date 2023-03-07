import { Player } from "./Player";
import { Rotation, Tile } from "./Tile";

export enum GameState {
    PlaceTile,
    MovePawn,
    End,
}

export interface BoardViewModel {
    id: number;
    remainingSeats: number;
    players: Array<Player>;
    state: BoardState;
    canPlay: boolean;
    gameState: GameState;
}

export interface BoardState {
    tiles: Array<Array<BoardTile>>;
    remainingTile: BoardTile;
};

export interface BoardTile {

    tile: Tile;

    rotation: Rotation;
}