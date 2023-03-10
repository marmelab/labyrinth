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
  onRotateRemainingTile: RotateRemainingTileHandler;
  onInsertTile: InsertTileHandler;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    gameState,
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
                canPlay={canPlay}
                boardTile={boardTile}
                gameState={gameState}
                coordinates={{ line, row }}
                onRotateRemainingTile={onRotateRemainingTile}
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
        boardTile={remainingTile}
        canPlay={canPlay}
        gameState={gameState}
        onRotateRemainingTile={onRotateRemainingTile}
        onInsertTile={onInsertTile}
      />
    </>
  );
};

export default BoardView;
