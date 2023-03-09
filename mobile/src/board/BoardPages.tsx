import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";

import { BoardListItem, Direction, GameState } from "./BoardTypes";
import { boardRepository } from "./BoardRepository";

import { useBoard } from "./BoardHooks";

import BoardView from "./components/BoardView";
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
        <strong>
          You are not signed in! <br />
          You can spectate these games:
        </strong>
      )}
      <ul>
        {boards.map((board) => (
          <li>
            <Link to={`/board/${board.id}/view`}>Board #{board.id}</Link>
          </li>
        ))}
      </ul>
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

  if (board) {
    if (board.remainingSeats > 0) {
      return (
        <>
          <strong>Waiting for {board.remainingSeats} player(s)</strong>
          <ul>
            {board.players.map((player, i) => (
              <li key={i}>{player?.name ?? "?"}</li>
            ))}
          </ul>
        </>
      );
    }

    return (
      <BoardView
        board={board}
        onRotateRemainingTile={onRotateRemainingTile}
        onInsertTile={onInsertTile}
        onMovePlayer={onMovePlayer}
      />
    );
  }

  if (error) {
    throw error;
  }

  return <p>Loading</p>;
}
