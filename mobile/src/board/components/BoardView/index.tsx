import { Board, Player, GameState } from "../../entity";

import PlayerPawnView from "../PlayerPawnView";
import TileView from "../TileView";

import "./index.css";

interface BoardProps {
  board: Board;
}

/**
 *
 */
const BoardView = ({
  board: {
    state: { tiles, remainingTile, gameState, players },
  },
}: BoardProps) => {
  return (
    <>
      <div className="board">
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => (
            <TileView
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
                  <PlayerPawnView key={`${color}`} color={color} />
                ))}
            </TileView>
          ))
        )}
      </div>
      <TileView boardTile={remainingTile} onClick={() => {}} />
    </>
  );
};

export default BoardView;
