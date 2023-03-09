import { Board, Color, GameState } from "../BoardTypes";

import PlayerPawnView from "./PlayerPawnView";

import TileView, {
  treasures,
  type InsertTileHandler,
  type RotateRemainingTileHandler,
  type MovePlayerHandler,
} from "./TileView";

import "./BoardView.css";

const colorNames = {
  [Color.Blue]: "Blue",
  [Color.Green]: "Green",
  [Color.Red]: "Red",
  [Color.Yellow]: "Yellow",
};

interface BoardProps {
  board: Board;
  onRotateRemainingTile: RotateRemainingTileHandler;
  onInsertTile: InsertTileHandler;
  onMovePlayer: MovePlayerHandler;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    gameState,
    canPlay,
    user,
  },
  onRotateRemainingTile,
  onInsertTile,
  onMovePlayer,
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
                onMovePlayer={onMovePlayer}
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

      <div className="state">
        <div className="state__col state__tile">
          <TileView
            boardTile={remainingTile}
            canPlay={canPlay}
            gameState={gameState}
            onRotateRemainingTile={onRotateRemainingTile}
            onInsertTile={onInsertTile}
            onMovePlayer={onMovePlayer}
          />
        </div>
        {user && (
          <div className="state__col state__info">
            <div className="state__row">
              <div className="state__row__label">Your name</div>
              <div className="state__row__value">{user.name}</div>
            </div>
            <div className="state__row">
              <div className="state__row__label">Your color</div>
              <div className="state__row__value">{colorNames[user.color]}</div>
            </div>
            <div className="state__row">
              <div className="state__row__label">Your target</div>
              <div className="state__row__value">
                {treasures[user.currentTarget]}
              </div>
            </div>
          </div>
        )}
      </div>
    </>
  );
};

export default BoardView;
