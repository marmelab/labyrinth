import { BoardState, GameState } from "../../model/Board";
import Tile from "../Tile";

import "./index.css";

interface BoardProps {
  state: BoardState;
}

/**
 *
 */
const Board = ({ state: { tiles, remainingTile, gameState } }: BoardProps) => {
  return (
    <>
      {" "}
      <div className="board">
        {tiles.flatMap((line) =>
          line.map((boardTile) => (
            <Tile
              boardTile={boardTile}
              disabled={gameState != GameState.MovePawn}
              onClick={() => {}}
            />
          ))
        )}
      </div>
      <Tile boardTile={remainingTile} onClick={() => {}} />
    </>
  );
};

export default Board;
