import { useEffect, useState } from "react";

import { Board } from "./BoardTypes";

import { boardRepository } from "./BoardRepository";

export function useBoard(id: number | string): [Board | null, any | null] {
  const [board, setBoard] = useState<Board | null>(null);
  const [error, setError] = useState<any | null>(null);

  const fetchBoard = async function () {
    try {
      const updatedBoard = await boardRepository.getById(id);
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
