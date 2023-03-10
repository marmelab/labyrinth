import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";

import {
  Button,
  List as MuiList,
  ListSubheader,
  ListItem,
  ListItemText,
  Typography,
} from "@mui/material";

import { BoardListItem, Direction } from "./BoardTypes";
import { boardRepository } from "./BoardRepository";

import { useBoard } from "./BoardHooks";

import BoardView from "./components/BoardView";
import TileView from "./components/TileView";
import PlayerPawnView from "./components/PlayerPawnView";
import { useUserContext } from "../user/UserHooks";

export function List() {
  const [user, _] = useUserContext();
  const [boards, setBoards] = useState<BoardListItem[]>([]);

  useEffect(() => {
    boardRepository.list(1).then(setBoards);
  }, [user]);

  return (
    <>
      {!user && (
        <Typography fontWeight={700}>
          You are not signed in! <br />
          You can spectate these games:
        </Typography>
      )}
      <MuiList
        sx={{ width: "100%", maxWidth: 360, bgcolor: "background.paper" }}
        component="nav"
        aria-labelledby="nested-list-subheader"
        subheader={
          <ListSubheader component="div" id="nested-list-subheader">
            Boards
          </ListSubheader>
        }
      >
        {boards.map((board) => (
          <ListItem>
            <ListItemText
              primary={
                <Link to={`/board/${board.id}/view`}>Board #{board.id}</Link>
              }
            />
          </ListItem>
        ))}
      </MuiList>
    </>
  );
}

export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);

  const onRotateRemainingTile = () => boardRepository.rotateRemainingTile(id!);

  const onInsertTile = (direction: Direction, index: number) =>
    boardRepository.insertTile(id!, direction, index);

  const onMovePlayer = (line: number, row: number) => {
    return boardRepository.movePlayer(id!, line, row);
  };

  const handleJoin = () => boardRepository.joinBoard(id!);

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

    return (
      <BoardView
        remainingTile={
          <TileView
            boardTile={remainingTile}
            canPlay={canPlay}
            gameState={gameState}
            playerTarget={user?.currentTarget}
            onRotateRemainingTile={onRotateRemainingTile}
            onInsertTile={onInsertTile}
            onMovePlayer={onMovePlayer}
          />
        }
        user={user}
      >
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => {
            return (
              <TileView
                key={`${line * tiles.length + row}`}
                canPlay={canPlay}
                boardTile={boardTile}
                gameState={gameState}
                coordinates={{ line, row }}
                playerTarget={user?.currentTarget}
                onRotateRemainingTile={onRotateRemainingTile}
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
  }

  if (error) {
    throw error;
  }

  return <p>Loading</p>;
}
