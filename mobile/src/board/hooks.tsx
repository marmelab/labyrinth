import { useContext, useEffect, useState } from "react";

import { Board } from "./entity";
import { BoardRepository } from "./repository";

export function useBoard(initialState: Board): Board {
  const [board, setBoard] = useState<Board>(initialState);

  useEffect(() => {
    const mercureURL = `/.well-known/mercure?topic=${encodeURI(
      window.location.href
    )}`;

    const eventSource = new EventSource(mercureURL);
    eventSource.onmessage = async () => {
      try {
        const updatedBoard = await BoardRepository.getById(board.id);
        setBoard(updatedBoard);
      } catch (e) {
        console.error(e);
      }
    };

    return () => {
      eventSource.close();
    };
  }, []);

  return board;
}
