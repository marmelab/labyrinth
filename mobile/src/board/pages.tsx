import { useParams } from "react-router-dom";

import BoardView from "./components/BoardView";
import { useBoard } from "./hooks";

export function GetById() {
  const { id } = useParams();
  const [board, error] = useBoard(id!);

  if (board) {
    return <BoardView board={board} />;
  }

  if (error) {
    throw error;
  }

  return <p>Loading</p>;
}
