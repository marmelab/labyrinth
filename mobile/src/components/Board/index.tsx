import { BoardState, GameState } from "../../model/Board";
import { Player } from "../../model/Player";

import PlayerComponent from "../Player";
import TileComponent from "../Tile";

import "./index.css";

interface BoardProps {
  state: BoardState;
}

/**
 *
 */
const Board = ({
  state: { tiles, remainingTile, gameState, players },
}: BoardProps) => {
  return (
    <>
      <div className="board">
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => (
            <TileComponent
              key={`${line * tiles.length + row}`}
              boardTile={boardTile}
              disabled={gameState != GameState.MovePawn}
              onClick={() => {}}
            >
              {players
                .filter(
                  ({ position }: Player) =>
                    position.line == line && position.row == row
                )
                .map(({ color }: Player) => (
                  <PlayerComponent key={`${color}`} color={color} />
                ))}
            </TileComponent>
          ))
        )}
      </div>
      <TileComponent boardTile={remainingTile} onClick={() => {}} />
    </>
  );
};

export default Board;
