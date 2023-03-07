import { ReactElement, useContext } from "react";
import { useLoaderData, LoaderFunctionArgs } from "react-router-dom";

import { Board } from "./entity";
import BoardView from "./components/BoardView";
import { useBoard } from "./hooks";

export function GetById(): ReactElement {
  const initialState = useLoaderData() as Board;
  const board = useBoard(initialState);

  return <BoardView board={board} />;
}
