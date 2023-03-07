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
    name: string;
    color: Color;
    line: number;
    row: number;
    targets: Array<string>;
    currentTarget: string;
    score: number;
    isCurrentPlayer: boolean;
    isUser: boolean;
};