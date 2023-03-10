import { useParams } from "react-router-dom";

import { Direction, GameState } from "./BoardTypes";
import { boardRepository } from "./BoardRepository";
import { useBoard } from "./BoardHooks";
import BoardView from "./components/BoardView";

export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);

  const onRotateRemainingTile = () => boardRepository.rotateRemainingTile(id!);

  const onInsertTile = (direction: Direction, index: number) =>
    boardRepository.insertTile(id!, direction, index);

  if (board) {
    return (
      <BoardView
        board={board}
        onRotateRemainingTile={onRotateRemainingTile}
        onInsertTile={onInsertTile}
      />
    );
  }

  if (error) {
    throw error;
  }

  return <p>Loading</p>;
}
