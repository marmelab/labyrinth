import { useParams } from "react-router-dom";

import BoardView from "./components/BoardView";
import { useBoard } from "./BoardHooks";
import { boardRepository } from "./BoardRepository";

export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);

  const rotateRemainingTile = async () => {
    boardRepository.rotateRemainingTile(id!);
  };

  if (board) {
    return (
      <BoardView
        board={board}
        rotateRemainingTile={board.canPlay ? rotateRemainingTile : undefined}
      />
    );
  }

  if (error) {
    throw error;
  }

  return <p>Loading</p>;
}
