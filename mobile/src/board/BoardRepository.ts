import { Board, BoardID, Direction } from "./BoardTypes";

export const boardRepository = {
  async getById(id: number | string): Promise<Board> {
    const response = await fetch(`/api/v1/board/${id}`);
    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: Board } = await response.json();
    return responseContent.data;
  },
  async rotateRemainingTile(id: BoardID): Promise<void> {
    const response = await fetch(`/api/v1/board/${id}/rotate-remaining`, {
      method: "POST",
    });
    if (response.status != 200) {
      throw response;
    }
  },

  async insertTile(
    id: BoardID,
    direction: Direction,
    index: number
  ): Promise<void> {
    const response = await fetch(`/api/v1/board/${id}/insert-tile`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        direction,
        index,
      }),
    });
    if (response.status != 200) {
      throw response;
    }
  },
};
