import { BoardTile, Direction, GameState } from "../BoardTypes";

import { RotateRemainingTileHandler, Tile } from "./Tile";

interface RemainingTileViewProps {
  boardTile: BoardTile;
  canPlay: boolean;
  gameState: GameState;
  playerTarget?: string;
  onRotateRemainingTile: RotateRemainingTileHandler;
}

export const RemainingTileView = ({
  boardTile,
  canPlay,
  gameState,
  playerTarget,
  onRotateRemainingTile,
}: RemainingTileViewProps) => {
  const bt = {
    ...boardTile,
    top: boardTile.top ?? 380,
    left: boardTile.left ?? 5,
    opacity: boardTile.opacity ?? 1,
  };

  if (!canPlay || gameState == GameState.End) {
    return (
      <Tile disabled boardTile={bt} playerTarget={playerTarget} remainingTile />
    );
  }

  return (
    <Tile
      animate={gameState == GameState.Animating}
      boardTile={bt}
      playerTarget={playerTarget}
      onClick={onRotateRemainingTile}
      remainingTile
    />
  );
};
