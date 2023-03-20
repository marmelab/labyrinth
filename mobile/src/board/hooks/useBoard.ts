import { useEffect, useState, useRef } from "react";

import {
  type Board,
  GameState,
  Rotation,
  Direction,
  BoardTile,
} from "../BoardTypes";

import { boardRepository } from "../BoardRepository";

const timeout = (duration: number) =>
  new Promise<void>((resolve) => setTimeout(resolve, duration));

const animationFrame = (callback: () => void) =>
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
const moveTileIndexes = new Map<Direction, Map<number, number[]>>([
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

enum Action {
  RotateRemaining = "ROTATE_REMAINING",
  GameStateChange = "GAME_STATE_CHANGE",
  PlaceTile = "PLACE_TILE",
  PlayerTurnChange = "PLAYER_TURN_CHANGE",
}

enum RotationDirection {
  Clockwise = "CLOCKWISE",
  AntiClockwise = "ANTICLOCKWISE",
}

interface RotationPayload {
  direction: RotationDirection;
  rotation: Rotation;
}

interface PlaceTilePayload {
  direction: Direction;
  index: number;
}

interface MovePawnPayload {
  line: number;
  row: number;
}

type GameStateChangePayload = GameState;

type NewPlayerTurnPayload = null;

interface UserAction {
  kind: Action;
  payload:
    | RotationPayload
    | PlaceTilePayload
    | MovePawnPayload
    | GameStateChangePayload
    | NewPlayerTurnPayload;
}

export function useBoard(id: number | string): [Board | null, any | null] {
  const [board, setBoard] = useState<Board | null>(null);
  const [error, setError] = useState<any | null>(null);

  const setGameState = (gameState: GameState) => {
    setBoard((board) =>
      board
        ? {
            ...board,
            gameState,
          }
        : null
    );
  };

  const fetchBoard = async function () {
    try {
      const updatedBoard = await boardRepository.getById(id);
      setBoard(updatedBoard);
    } catch (e) {
      setError("Failed to load board");
      setBoard(null);
    }
  };

  const handleMercureMessage = async ({ data }: { data: string }) => {
    const actions: UserAction[] = JSON.parse(data);

    for (const { kind, payload } of actions) {
      const resolver = actionResolvers.get(kind);
      if (resolver) {
        await resolver(payload);
      }
    }
  };

  const updateTiles = async (
    direction: Direction,
    { top, left }: { top?: number; left?: number },
    index: number
  ) => {
    await animationFrame(() => {
      const tileIndexes = moveTileIndexes.get(direction)!.get(index)!;
      const lastIndex = tileIndexes[6];

      setBoard((board) => {
        if (!board) {
          return null;
        }

        const { state } = board;
        const { tiles, remainingTile } = state;

        return {
          ...board,
          state: {
            ...state,
            tiles: tiles.map((tileLine, line) => {
              return tileLine.map((boardTile, row) => {
                const index = line * 7 + row;
                if (!tileIndexes.includes(index)) {
                  return boardTile;
                }

                return {
                  ...boardTile,
                  top,
                  left,
                  opacity: index == lastIndex ? 0 : 1,
                };
              });
            }),
            remainingTile: {
              ...remainingTile,
              top: remainingTile.top
                ? remainingTile.top + (top ?? 0)
                : undefined,
              left: remainingTile.left
                ? remainingTile.left + (left ?? 0)
                : undefined,
            },
          },
        };
      });
    });

    await timeout(500);
  };

  const moveTilesHandlers = new Map<
    Direction,
    (index: number) => Promise<void>
  >([
    [
      Direction.Top,
      updateTiles.bind(null, Direction.Top, {
        top: animateMoveTopOffset,
      }),
    ],
    [
      Direction.Right,
      updateTiles.bind(null, Direction.Right, {
        left: animateMoveRightOffset,
      }),
    ],
    [
      Direction.Bottom,
      updateTiles.bind(null, Direction.Bottom, {
        top: animateMoveBottomOffset,
      }),
    ],
    [
      Direction.Left,
      updateTiles.bind(null, Direction.Left, {
        left: animateMoveLeftOffset,
      }),
    ],
  ]);

  const actionResolvers = new Map<Action, (payload: any) => Promise<void>>([
    [
      Action.RotateRemaining,
      ({ direction }: RotationPayload) =>
        new Promise((resolve) => {
          setBoard((board) => {
            if (!board) {
              return null;
            }
            const remainingTile = board.state.remainingTile;

            const newRotation =
              direction == RotationDirection.Clockwise
                ? remainingTile.rotation + 90
                : remainingTile.rotation - 90;

            return {
              ...board,
              state: {
                ...board.state,
                remainingTile: {
                  tile: remainingTile.tile,
                  rotation: newRotation,
                },
              },
            };
          });

          setTimeout(resolve, 500);
        }),
    ],
    [
      Action.GameStateChange,
      (gameState: GameStateChangePayload) =>
        new Promise((resolve) => {
          setGameState(gameState);
          setTimeout(resolve, 0);
        }),
    ],
    [
      Action.PlaceTile,
      async (payload: PlaceTilePayload) => {
        setGameState(GameState.Animating);

        // Move remaining tile
        await animationFrame(() => {
          const position = remainingTileViewPositions
            .get(payload.direction)
            ?.get(payload.index);

          setBoard((board) =>
            board
              ? {
                  ...board,
                  state: {
                    ...board.state,
                    remainingTile: {
                      ...board.state.remainingTile,
                      top: position?.top,
                      left: position?.left,
                    },
                  },
                }
              : null
          );
        });

        // Let first animation play.
        await timeout(500);

        // Move tiles
        const animateMoveTiles = moveTilesHandlers.get(payload.direction);
        animateMoveTiles && (await animateMoveTiles(payload.index));

        // Ensures board is up to date.
        await fetchBoard();
      },
    ],
    [Action.PlayerTurnChange, () => fetchBoard()],
  ]);

  useEffect(() => {
    fetchBoard();

    const mercureURL = `/.well-known/mercure?topic=${encodeURI(
      window.location.pathname
    )}`;
    const eventSource = new EventSource(mercureURL);

    eventSource.onmessage = handleMercureMessage;
    return () => {
      eventSource.close();
    };
  }, []);

  return [board, error];
}
