import { useEffect, useState } from "react";

import { Board } from "./entity";
import { BoardRepository } from "./repository";

export function useBoard(id: number | string): [Board | null, any | null] {
  const [board, setBoard] = useState<Board | null>(null);
  const [error, setError] = useState<any | null>(null);

  const fetchBoard = async function () {
    try {
      const updatedBoard = await BoardRepository.getById(id);
      setBoard(updatedBoard);
    } catch (e) {
      setError("Failed to load board");
      setBoard(null);
    }
  };

  useEffect(() => {
    fetchBoard();

    const mercureURL = `/.well-known/mercure?topic=${encodeURI(
      window.location.href
    )}`;

    const eventSource = new EventSource(mercureURL);
    eventSource.onmessage = async () => {
      await fetchBoard();
    };

    return () => {
      eventSource.close();
    };
  }, []);

  return [board, error];
}
