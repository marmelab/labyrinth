import { Rotation, Tile } from "./Tile";

export interface Board {

    tiles: Array<BoardTile>;

};

export interface BoardTile {

    tile: Tile;

    rotation: Rotation;
}