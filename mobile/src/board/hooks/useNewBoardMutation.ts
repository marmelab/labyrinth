import { useMutation } from "react-query";

import { Board } from "../BoardTypes";

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
