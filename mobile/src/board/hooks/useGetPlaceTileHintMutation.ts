import { AlertColor } from "@mui/material";
import { useMutation } from "react-query";

import { type Board, type Error, type PlaceTileHint } from "../BoardTypes";

type MutationError = { message: string; severity?: AlertColor };

interface PlaceTileHintResponseData {
  placeTileHint: PlaceTileHint | null;
}

interface PlaceTileHintResponse {
  data: MutationError | PlaceTileHintResponseData;
}

export const useGetPlaceTileHintMutation = () =>
  useMutation<PlaceTileHint | null, Error, Board, void>(async (board) => {
    const response = await fetch(`/api/v1/board/${board.id}/place-tile-hint`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const responseContent: PlaceTileHintResponse = await response.json();
    if (response.status != 200) {
      const error = responseContent.data as MutationError;
      throw {
        severity: error.severity ?? "error",
        message: error.message,
      };
    }

    const data = responseContent.data as PlaceTileHintResponseData;
    if (!data) {
      return null;
    }
    return data.placeTileHint;
  });
