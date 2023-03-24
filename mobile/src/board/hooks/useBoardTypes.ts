import { Rotation, Direction, type Coordinate } from "../BoardTypes";

export const timeout = (duration: number) =>
  new Promise<void>((resolve) => setTimeout(resolve, duration));

export const animationFrame = (callback: () => void) =>
  new Promise<void>((resolve) =>
    window.requestAnimationFrame(() => {
      callback();
      resolve();
    })
  );

const placeTileTopOffset = -45;
const placeTileRightOffset = 355;
const placeTileBottomOffset = 355;
const placeTileLeftOffset = -45;

const firstIndexOffset = 55;
const secondIndexOffset = 155;
const thirdIndexOffset = 255;

export const remainingTileViewPositions = new Map<
  Direction,
  Map<number, { top: number; left: number }>
>([
  [
    Direction.Top,
    new Map([
      [
        1,
        {
          top: placeTileTopOffset,
          left: firstIndexOffset,
        },
      ],
      [
        3,
        {
          top: placeTileTopOffset,
          left: secondIndexOffset,
        },
      ],
      [
        5,
        {
          top: placeTileTopOffset,
          left: thirdIndexOffset,
        },
      ],
    ]),
  ],
  [
    Direction.Right,
    new Map([
      [
        1,
        {
          top: firstIndexOffset,
          left: placeTileRightOffset,
        },
      ],
      [
        3,
        {
          top: secondIndexOffset,
          left: placeTileRightOffset,
        },
      ],
      [
        5,
        {
          top: thirdIndexOffset,
          left: placeTileRightOffset,
        },
      ],
    ]),
  ],
  [
    Direction.Bottom,
    new Map([
      [
        1,
        {
          top: placeTileBottomOffset,
          left: firstIndexOffset,
        },
      ],
      [
        3,
        {
          top: placeTileBottomOffset,
          left: secondIndexOffset,
        },
      ],
      [
        5,
        {
          top: placeTileBottomOffset,
          left: thirdIndexOffset,
        },
      ],
    ]),
  ],
  [
    Direction.Left,
    new Map([
      [
        1,
        {
          top: firstIndexOffset,
          left: placeTileLeftOffset,
        },
      ],
      [
        3,
        {
          top: secondIndexOffset,
          left: placeTileLeftOffset,
        },
      ],
      [
        5,
        {
          top: thirdIndexOffset,
          left: placeTileLeftOffset,
        },
      ],
    ]),
  ],
]);

const animateMoveTopOffset = 50;
const animateMoveRightOffset = -50;
const animateMoveBottomOffset = -50;
const animateMoveLeftOffset = 50;
export const moveTileIndexes = new Map<Direction, Map<number, number[]>>([
  [
    Direction.Top,
    new Map([
      [1, [1, 8, 15, 22, 29, 36, 43]],
      [3, [3, 10, 17, 24, 31, 38, 45]],
      [5, [5, 12, 19, 26, 33, 40, 47]],
    ]),
  ],
  [
    Direction.Right,
    new Map([
      [1, [13, 12, 11, 10, 9, 8, 7]],
      [3, [27, 26, 25, 24, 23, 22, 21]],
      [5, [41, 40, 39, 38, 37, 36, 35]],
    ]),
  ],
  [
    Direction.Bottom,
    new Map([
      [1, [43, 36, 29, 22, 15, 8, 1]],
      [3, [45, 38, 31, 24, 17, 10, 3]],
      [5, [47, 40, 33, 26, 19, 12, 5]],
    ]),
  ],
  [
    Direction.Left,
    new Map([
      [1, [7, 8, 9, 10, 11, 12, 13]],
      [3, [21, 22, 23, 24, 25, 26, 27]],
      [5, [35, 36, 37, 38, 39, 40, 41]],
    ]),
  ],
]);

export const movePlayerIndex = new Map<
  Direction,
  (index: number, lastTileIndex: number) => number
>([
  [
    Direction.Top,
    (index, lastTileIndex) => {
      if (index == lastTileIndex) {
        return 0;
      }
      return index + 1;
    },
  ],
  [
    Direction.Right,
    (index, lastTileIndex) => {
      if (index == 0) {
        return lastTileIndex;
      }
      return index - 1;
    },
  ],
  [
    Direction.Bottom,
    (index, lastTileIndex) => {
      if (index == 0) {
        return lastTileIndex;
      }
      return index - 1;
    },
  ],
  [
    Direction.Left,
    (index, lastTileIndex) => {
      if (index == lastTileIndex) {
        return 0;
      }
      return index + 1;
    },
  ],
]);

export const moveTileOffsets = new Map<
  Direction,
  { top?: number; left?: number }
>([
  [
    Direction.Top,
    {
      top: animateMoveTopOffset,
    },
  ],
  [
    Direction.Right,
    {
      left: animateMoveRightOffset,
    },
  ],
  [
    Direction.Bottom,
    {
      top: animateMoveBottomOffset,
    },
  ],
  [
    Direction.Left,
    {
      left: animateMoveLeftOffset,
    },
  ],
]);

export enum Action {
  RotateRemaining = "ROTATE_REMAINING",
  PlaceTile = "PLACE_TILE",
  MovePawn = "MOVE_PAWN",
  NewPlayer = "NEW_PLAYER",
}

export enum RotationDirection {
  Clockwise = "CLOCKWISE",
  AntiClockwise = "ANTICLOCKWISE",
}

export interface RotationPayload {
  direction: RotationDirection | "";
  rotation: Rotation;
}

export interface PlaceTilePayload {
  direction: Direction;
  index: number;
}

export interface MovePawnPayload {
  line: number;
  row: number;
  path: Coordinate[];
}

export interface UserAction {
  kind: Action;
  payload: RotationPayload | PlaceTilePayload | MovePawnPayload;
}
