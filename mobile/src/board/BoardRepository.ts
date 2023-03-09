import { Board } from "./BoardTypes";

export type BoardID = number | string;

export interface BoardRepository {
  getById(id: BoardID): Promise<Board>;
  rotateRemainingTile(id: BoardID): Promise<void>;
}
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
};
