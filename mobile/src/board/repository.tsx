import { Board } from "./entity";

export class BoardRepository {
  static async getById(id: number | string): Promise<Board> {
    const response = await fetch(`/api/v1/board/${id}`);
    if (response.status != 200) {
      throw response;
    }
    return response.json();
  }
}
