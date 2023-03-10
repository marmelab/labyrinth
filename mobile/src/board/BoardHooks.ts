import { useEffect, useState } from "react";
import { useMutation } from "react-query";

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

type MutationError = { message: string };

type NewBoardVariables = { playerCount: number };
type NewBoardResponse = { data: MutationError | Board };

export const useNewBoardMutation = () =>
  useMutation<Board, MutationError, NewBoardVariables, void>(
    async ({ playerCount }) => {
      const response = await fetch(`/api/v1/board`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ playerCount }),
      });

      const responseContent: NewBoardResponse = await response.json();
      if (response.status != 200) {
        throw responseContent.data;
      }

      return responseContent.data as Board;
    }
  );
