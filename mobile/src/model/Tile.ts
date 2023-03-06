export enum Shape {
    ShapeI = 0,
    ShapeT = 1,
    ShapeV = 1,
};

export enum Rotation {
    Rotation0 = 0,
    Rotation90 = 90,
    Rotation180 = 180,
    Rotation270 = 270,
};

export interface Tile {
    treasure: string,
    shape: Shape,
};