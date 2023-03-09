import { Board, GameState } from "../BoardTypes";

import { Direction } from "../BoardTypes";

import PlayerPawnView from "./PlayerPawnView";
import TileView, {
  type InsertTileHandler,
  type RotateRemainingTileHandler,
} from "./TileView";

import "./BoardView.css";

interface BoardProps {
  board: Board;
  onRotateRemainingTile?: RotateRemainingTileHandler;
  onInsertTile?: InsertTileHandler;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    canPlay,
  },
  onRotateRemainingTile,
  onInsertTile,
}: BoardProps) => {
  return (
    <>
      <div className="board">
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => {
            return (
              <TileView
                key={`${line * tiles.length + row}`}
                line={line}
                row={row}
                boardTile={boardTile}
                onInsertTile={onInsertTile}
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
            );
          })
        )}
      </div>

      <TileView
        line={-1}
        row={-1}
        boardTile={remainingTile}
        disabled={!canPlay}
        onRotateRemainingTile={canPlay ? onRotateRemainingTile : undefined}
      />
    </>
  );
};

export default BoardView;
