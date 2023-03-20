import { useEffect, useState } from "react";

import { type Error, type Board, GameState, Direction } from "../BoardTypes";

import { boardRepository } from "../BoardRepository";

import {
  Action,
  animationFrame,
  moveTileIndexes,
  moveTileOffsets,
  type PlaceTilePayload,
  remainingTileViewPositions,
  RotationDirection,
  type RotationPayload,
  timeout,
  type UserAction,
} from "./useBoardTypes";

export function useBoard(id: number | string): [Board | null, Error | null] {
  const [board, setBoard] = useState<Board | null>(null);
  const [error, setError] = useState<Error | null>(null);

  const fetchBoard = async function () {
    setError(null);
    try {
      const updatedBoard = await boardRepository.getById(id);
      setBoard(updatedBoard);
    } catch (e) {
      setError({ message: "Failed to load board", severity: "error" });
      setBoard(null);
    }
  };

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

  const handleMercureMessage = async ({ data }: { data: string }) => {
    const actions: UserAction[] = JSON.parse(data);
    if (actions.length == 0) {
      setError({
        severity: "warning",
        message: "You cannot perform this action",
      });
      return;
    }

    for (const { kind, payload } of actions) {
      const resolver = actionResolvers.get(kind);
      if (resolver) {
        await resolver(payload);
      }
    }
  };

  const handleMoveTile = async (direction: Direction, index: number) => {
    await animationFrame(() => {
      const tileIndexes = moveTileIndexes.get(direction)!.get(index)!;
      const lastIndex = tileIndexes[6];
      const { top, left } = moveTileOffsets.get(direction)!;
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
                  opacity: index === lastIndex ? 0 : 1,
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

  const actionResolvers = new Map<Action, (payload: any) => Promise<void>>([
    [
      Action.RotateRemaining,
      ({ direction, rotation }: RotationPayload) =>
        new Promise((resolve) => {
          setBoard((board) => {
            if (!board) {
              return null;
            }
            const remainingTile = board.state.remainingTile;

            if (direction == "") {
              return {
                ...board,
                state: {
                  ...board.state,
                  remainingTile: {
                    tile: remainingTile.tile,
                    rotation: rotation,
                  },
                },
              };
            }

            return {
              ...board,
              state: {
                ...board.state,
                remainingTile: {
                  tile: remainingTile.tile,
                  rotation:
                    direction == RotationDirection.Clockwise
                      ? remainingTile.rotation + 90
                      : remainingTile.rotation - 90,
                },
              },
            };
          });

          setTimeout(resolve, 500);
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
        await handleMoveTile(payload.direction, payload.index);

        // Ensures board is up to date.
        await fetchBoard();
      },
    ],
    [Action.MovePawn, () => fetchBoard()],
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
