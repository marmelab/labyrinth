import { AlertColor } from "@mui/material";
import { useMutation } from "react-query";

import type { BoardID, Error, PlaceTileHint } from "../BoardTypes";

type MutationError = { message: string; severity?: AlertColor };

interface PlaceTileHintResponse {
  data: MutationError | PlaceTileHint;
}

export const useGetPlaceTileHintMutation = () =>
  useMutation<PlaceTileHint | null, Error, BoardID, void>(async (id) => {
    const response = await fetch(`/api/v1/board/${id}/place-tile-hint`, {
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

    const data = responseContent.data as PlaceTileHint;
    if (!data) {
      return null;
    }

    return data;
  });