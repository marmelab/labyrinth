const gameStates = ["Place Tile", "Move Pawn", "Completed"];
export const renderGameState = ({ game_state }: { game_state: number }) => {
  return gameStates[game_state];
};

const colors = ["Blue", "Green", "Red", "Yellow"];
export const renderPlayerColor = ({ color }: { color: number }) => {
  return colors[color];
};

export const renderScore = ({
  score,
  targets,
}: {
  score: number;
  targets: string[];
}) => {
  return `${score} / ${targets.length + score}`;
};
