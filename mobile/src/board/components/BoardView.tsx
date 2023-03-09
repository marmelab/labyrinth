import { Board, GameState } from "../BoardTypes";
import type { BoardRepository } from "../BoardRepository";

import PlayerPawnView from "./PlayerPawnView";
import TileView from "./TileView";

import "./BoardView.css";

interface BoardProps {
  board: Board;
  onRotateRemainingTile?: () => Promise<void>;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    canPlay,
    gameState,
  },
  onRotateRemainingTile,
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

      <TileView
        boardTile={remainingTile}
        disabled={!canPlay}
        onClick={onRotateRemainingTile}
      />
    </>
  );
};

export default BoardView;
