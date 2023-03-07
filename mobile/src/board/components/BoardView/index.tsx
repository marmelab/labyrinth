import { Board, Player, GameState } from "../../entity";

import PlayerPawnView from "../PlayerPawnView";
import TileView from "../TileView";

import "./index.css";

interface BoardProps {
  board: Board;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    canPlay,
    gameState,
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
              disabled={!canPlay || gameState != GameState.MovePawn}
              onClick={() => {}}
            >
              {players
                .filter((player) => player.line == line && player.row == row)
                .map((player) => {
                  return (
                    <PlayerPawnView
                      key={`${player.color}`}
                      color={player.color}
                    />
                  );
                })}
            </TileView>
          ))
        )}
      </div>
      <TileView boardTile={remainingTile} onClick={() => {}} />
    </>
  );
};

export default BoardView;
