import { Rotation, Tile } from "./Tile";

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
    score: number,
};