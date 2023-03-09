import { Board, BoardID, BoardListItem, Direction } from "./BoardTypes";

export const boardRepository = {
  async list(page: number): Promise<BoardListItem[]> {
    const response = await fetch(`/api/v1/board?page=${page}`);
    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: BoardListItem[] } = await response.json();

    return responseContent.data;
  },
  async getById(id: BoardID): Promise<Board> {
    const response = await fetch(`/api/v1/board/${id}`);
    if (response.status != 200) {
      throw response;
    }

    const responseContent: { data: Board } = await response.json();
    return responseContent.data;
  },

  async rotateRemainingTile(id: BoardID) {
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
  async movePlayer(id: BoardID, line: number, row: number) {
    const response = await fetch(`/api/v1/board/${id}/move-player`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        line,
        row,
      }),
    });

    if (response.status != 200) {
      throw response;
    }
  },
};
