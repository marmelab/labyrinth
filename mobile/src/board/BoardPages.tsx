import { useParams } from "react-router-dom";

import BoardView from "./components/BoardView";
import { useBoard } from "./BoardHooks";
import { boardRepository } from "./BoardRepository";

export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);

  const onRotateRemainingTile = () => boardRepository.rotateRemainingTile(id!);

  if (board) {
    return (
      <BoardView
        board={board}
        onRotateRemainingTile={
          board.canPlay ? onRotateRemainingTile : undefined
        }
      />
    );
  }

  if (error) {
    throw error;
  }

  return <p>Loading</p>;
}
