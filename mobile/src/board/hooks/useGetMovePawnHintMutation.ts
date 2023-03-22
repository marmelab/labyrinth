import { AlertColor } from "@mui/material";
import { useMutation } from "react-query";

import type { BoardID, Coordinate, Error, TileInsertion } from "../BoardTypes";

type MutationError = { message: string; severity?: AlertColor };

interface MovePawnHintResponse {
  data: MutationError | Coordinate;
}

export const useGetMovePawnHintMutation = () =>
  useMutation<Coordinate | null, Error, BoardID, void>(async (id) => {
    const response = await fetch(`/api/v1/board/${id}/move-pawn-hint`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const responseContent: MovePawnHintResponse = await response.json();
    if (response.status != 200) {
      const error = responseContent.data as MutationError;
      throw {
        severity: error.severity ?? "error",
        message: error.message,
      };
    }

    return responseContent.data as Coordinate | null;
  });
