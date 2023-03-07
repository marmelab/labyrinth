import React, { useEffect, useState } from "react";

import { useLoaderData, LoaderFunctionArgs } from "react-router-dom";

import Board from "../../components/Board";
import { BoardViewModel, BoardState } from "../../model/Board";

async function getById(id: number | string): Promise<BoardViewModel> {
  const response = await fetch(`/api/v1/board/${id}`);

  if (response.status != 200) {
    throw new Error("Not Found");
  }

  const responseBody: { data: BoardViewModel } = await response.json();
  return responseBody.data;
}

export async function getByIdLoader({
  params: { id },
}: LoaderFunctionArgs): Promise<BoardViewModel> {
  return getById(id as string);
}

export function GetById(): React.ReactElement {
  const initialState = useLoaderData() as BoardViewModel;
  const [board, setBoard] = useState<BoardViewModel>(initialState);

  useEffect(() => {
    const mercureURL = `/.well-known/mercure?topic=${encodeURI(
      window.location.href
    )}`;

    const eventSource = new EventSource(mercureURL);
    eventSource.onmessage = async () => {
      setBoard(await getById(board.id));
    };

    return () => {
      eventSource.close();
    };
  }, []);

  return <Board board={board} />;
}
