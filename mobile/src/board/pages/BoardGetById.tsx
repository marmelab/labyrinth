import { useParams } from "react-router-dom";

import {
  Button,
  List as MuiList,
  ListSubheader,
  ListItem,
  ListItemText,
  Typography,
  Alert,
} from "@mui/material";

import {
  type Error,
  Direction,
  PlaceTileHint,
  GameState,
  Coordinate,
} from "../BoardTypes";
import { boardRepository } from "../BoardRepository";

import { useBoard, useGetPlaceTileHintMutation } from "../BoardHooks";

import BoardView from "../components/BoardView";
import { RemainingTileView, TileView } from "../components/TileView";
import PlayerPawnView from "../components/PlayerPawnView";
import { useGetMovePawnHintMutation } from "../hooks/useGetMovePawnHintMutation";

const placeTileHintIndexes = new Map<Direction, Map<number, number>>([
  [
    Direction.Top,
    new Map([
      [1, 1],
      [3, 3],
      [5, 5],
    ]),
  ],
  [
    Direction.Right,
    new Map([
      [1, 13],
      [3, 27],
      [5, 41],
    ]),
  ],
  [
    Direction.Bottom,
    new Map([
      [1, 43],
      [3, 45],
      [5, 47],
    ]),
  ],
  [
    Direction.Left,
    new Map([
      [1, 7],
      [3, 21],
      [5, 35],
    ]),
  ],
]);

const getHintIndex = (
  placeTileHint?: PlaceTileHint | null,
  movePawnHint?: (Coordinate & { size: number }) | null
): number | undefined => {
  if (placeTileHint) {
    const { direction, index } = placeTileHint;
    return placeTileHintIndexes.get(direction)?.get(index);
  }

  if (movePawnHint) {
    return movePawnHint.line * 7 + movePawnHint.row;
  }

  return undefined;
};

type Handler = (...args: any[]) => Promise<any>;
export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);

  const placeTileHint = useGetPlaceTileHintMutation();
  const movePawnHint = useGetMovePawnHintMutation();

  const handleUserAction = <F extends Handler>(handler: F): F => {
    return ((...args: any[]) => {
      placeTileHint.reset();
      movePawnHint.reset();
      return handler(...args);
    }) as Handler as F;
  };

  const onRotateRemainingTile = handleUserAction(() =>
    boardRepository.rotateRemainingTile(id!)
  );

  const onInsertTile = handleUserAction(
    (direction: Direction, index: number) => {
      return boardRepository.insertTile(id!, direction, index);
    }
  );

  const onMovePlayer = handleUserAction((line: number, row: number) => {
    return boardRepository.movePlayer(id!, line, row);
  });

  const handleJoin = () => boardRepository.joinBoard(id!);

  const handleGetHint = handleUserAction(async () => {
    if (board?.gameState == GameState.PlaceTile) {
      await placeTileHint.mutateAsync(id!);
    } else if (board?.gameState == GameState.MovePawn) {
      await movePawnHint.mutateAsync(id!);
    }
  });

  if (board) {
    if (board.remainingSeats > 0) {
      return (
        <>
          <Typography fontWeight={700}>
            Waiting for {board.remainingSeats} player(s)
          </Typography>
          <MuiList
            sx={{ width: "100%", maxWidth: 360, bgcolor: "background.paper" }}
            component="nav"
            aria-labelledby="nested-list-subheader"
            subheader={
              <ListSubheader component="div" id="nested-list-subheader">
                Players
              </ListSubheader>
            }
          >
            {board.players.map((player, i) => (
              <ListItem key={player?.color ?? i}>
                <ListItemText primary={player?.name ?? "?"} />
              </ListItem>
            ))}
          </MuiList>
          {board.canJoin && (
            <Button variant="contained" onClick={handleJoin}>
              Join Game
            </Button>
          )}
        </>
      );
    }

    const {
      gameState,
      canPlay,
      state: { tiles, remainingTile },
      players,
      user,
    } = board;

    const errors: Error[] = [];
    if (error) {
      errors.push(error);
    }

    if (placeTileHint.isError) {
      errors.push(placeTileHint.error);
    }

    if (movePawnHint.isError) {
      errors.push(movePawnHint.error);
    }

    const hintIndex = getHintIndex(
      placeTileHint.data,
      movePawnHint.data
        ? { ...movePawnHint.data, size: board.state.tiles.length }
        : undefined
    );

    return (
      <BoardView
        canPlay={board.canPlay}
        remainingTile={
          <RemainingTileView
            boardTile={remainingTile}
            canPlay={canPlay}
            gameState={gameState}
            playerTarget={user?.currentTarget}
            onRotateRemainingTile={onRotateRemainingTile}
          />
        }
        currentPlayer={board.currentPlayer}
        user={user}
        errors={errors}
        handleGetHint={handleGetHint}
      >
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => {
            const index = line * tiles.length + row;
            return (
              <TileView
                key={`${line * tiles.length + row}`}
                canPlay={canPlay}
                boardTile={boardTile}
                gameState={gameState}
                coordinates={{ line, row }}
                hint={hintIndex == index}
                playerTarget={user?.currentTarget}
                onInsertTile={onInsertTile}
                onMovePlayer={onMovePlayer}
                isAccessible={
                  !!board.accessibleTiles?.coordinates?.find(
                    (coordinate) =>
                      coordinate.line === line && coordinate.row === row
                  )
                }
              >
                {players
                  .filter((player) => player.line == line && player.row == row)
                  .map((player) => {
                    return (
                      <PlayerPawnView
                        key={`${player.color}`}
                        color={player.color}
                      />
                    );
                  })}
              </TileView>
            );
          })
        )}
      </BoardView>
    );
  } else if (error) {
    return <Alert severity={error.severity}>{error.message}</Alert>;
  }

  return <Alert severity="info">Loading</Alert>;
}
