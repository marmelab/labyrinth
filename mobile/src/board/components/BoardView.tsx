import { Board, GameState } from "../BoardTypes";

import { Direction } from "../BoardTypes";

import PlayerPawnView from "./PlayerPawnView";
import TileView from "./TileView";

import "./BoardView.css";

type Listener = (() => Promise<void>) | undefined;
type ListenerFactory = (line: number, row: number) => Listener;

type RotateRemainingTypeHandler = () => Promise<void>;
type InsertTileHandler = (direction: Direction, index: number) => Promise<void>;

function createTileListenerFactory(
  gameState: GameState,
  onInsertTile?: InsertTileHandler
): ListenerFactory {
  // First Map is (line => row)
  // Second Map is (row => listener)
  const listeners = new Map<number, Map<number, Listener>>();

  const insertableIndexes = [1, 3, 5];

  if (gameState == GameState.PlaceTile) {
    listeners.set(
      0,
      new Map(
        insertableIndexes.map((index) => [
          index,
          onInsertTile?.bind(null, Direction.Top, index),
        ])
      )
    );

    insertableIndexes.forEach((line) =>
      listeners.set(
        line,
        new Map([
          [0, onInsertTile?.bind(null, Direction.Left, line)],
          [6, onInsertTile?.bind(null, Direction.Right, line)],
        ])
      )
    );

    listeners.set(
      6,
      new Map(
        insertableIndexes.map((index) => [
          index,
          onInsertTile?.bind(null, Direction.Bottom, index),
        ])
      )
    );
  }

  return (line: number, row: number): Listener => {
    return listeners.get(line)?.get(row);
  };
}

interface BoardProps {
  board: Board;
  onRotateRemainingTile?: RotateRemainingTypeHandler;
  onInsertTile?: InsertTileHandler;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    canPlay,
    gameState,
  },
  onRotateRemainingTile,
  onInsertTile,
}: BoardProps) => {
  const tileListenerFactory = createTileListenerFactory(
    gameState,
    onInsertTile
  );

  return (
    <>
      <div className="board">
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => {
            const listener = tileListenerFactory(line, row);
            return (
              <TileView
                key={`${line * tiles.length + row}`}
                boardTile={boardTile}
                disabled={!listener}
                onClick={listener}
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
        disabled={!canPlay}
        onClick={onRotateRemainingTile}
      />
    </>
  );
};

export default BoardView;
