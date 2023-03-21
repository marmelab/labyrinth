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

import { type Error, Direction, PlaceTileHint } from "../BoardTypes";
import { boardRepository } from "../BoardRepository";

import { useBoard, useGetPlaceTileHintMutation } from "../BoardHooks";

import BoardView from "../components/BoardView";
import { RemainingTileView, TileView } from "../components/TileView";
import PlayerPawnView from "../components/PlayerPawnView";

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
  placeTileHint?: PlaceTileHint | null
): number | undefined => {
  if (placeTileHint) {
    const { direction, index } = placeTileHint;
    return placeTileHintIndexes.get(direction)?.get(index);
  }

  return undefined;
};

export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);
  const placeTileHint = useGetPlaceTileHintMutation();

  const onRotateRemainingTile = () => boardRepository.rotateRemainingTile(id!);

  const onInsertTile = (direction: Direction, index: number) => {
    placeTileHint.reset();
    return boardRepository.insertTile(id!, direction, index);
  };

  const onMovePlayer = (line: number, row: number) => {
    placeTileHint.reset();
    return boardRepository.movePlayer(id!, line, row);
  };

  const handleJoin = () => boardRepository.joinBoard(id!);

  const handleGetPlaceTileHint = () => {
    placeTileHint.reset();
    return placeTileHint.mutate(id!);
  };

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
              <ListItem>
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

    const hintIndex = getHintIndex(placeTileHint.data);

    return (
      <BoardView
        gameState={gameState}
        remainingTile={
          <RemainingTileView
            boardTile={remainingTile}
            canPlay={canPlay}
            gameState={gameState}
            playerTarget={user?.currentTarget}
            onRotateRemainingTile={onRotateRemainingTile}
          />
        }
        user={user}
        errors={errors}
        handleGetPlaceTileHint={handleGetPlaceTileHint}
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
