import { useMutation } from "react-query";

import { Board, OpponentKind } from "../BoardTypes";

type MutationError = { message: string };

type NewBoardVariables = { playerCount: number; opponentKind: OpponentKind };
type NewBoardResponse = { data: MutationError | Board };

export const useNewBoardMutation = () =>
  useMutation<Board, MutationError, NewBoardVariables, void>(
    async ({ playerCount, opponentKind }) => {
      const response = await fetch(`/api/v1/board`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ playerCount, opponentKind }),
      });

      const responseContent: NewBoardResponse = await response.json();
      if (response.status != 200) {
        throw responseContent.data;
      }

      return responseContent.data as Board;
    }
  );
