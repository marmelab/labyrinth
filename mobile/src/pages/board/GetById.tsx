import React, { useEffect, useState } from "react";

import { useLoaderData, LoaderFunctionArgs } from "react-router-dom";

import Board from "../../components/Board";
import { Board as BoardModel, BoardState } from "../../model/Board";

export async function getByIdLoader({
  params: { id },
}: LoaderFunctionArgs): Promise<BoardModel> {
  const response = await fetch(`/api/v1/board/${id}`);

  if (response.status != 200) {
    throw new Error("Not Found");
  }
  return response.json();
}

export function GetById(): React.ReactElement {
  const { state: initialBoardState } = useLoaderData() as BoardModel;

  const [boardState, setBoardSate] = useState<BoardState>(initialBoardState);
  useEffect(() => {
    const mercureURL = `/.well-known/mercure?topic=${encodeURI(
      window.location.href
    )}`;

    const eventSource = new EventSource(mercureURL);
    eventSource.onmessage = ({ data }: MessageEvent) => {
      const { state } = JSON.parse(data) as BoardModel;
      setBoardSate(state);
    };

    return () => {
      eventSource.close();
    };
  }, []);

  return <Board state={boardState} />;
}
