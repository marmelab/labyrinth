export const gameState = {
  placeTile: 0,
  movePawn: 1,
  completed: 2,
};

const gameStateLabels = {
  [gameState.placeTile]: "Place Tile",
  [gameState.movePawn]: "Move Pawn",
  [gameState.completed]: "Completed",
};

export const renderGameState = ({ game_state }: { game_state: number }) => {
  return gameStateLabels[game_state];
};

export const colors = {
  blue: 0,
  green: 1,
  red: 2,
  yellow: 3,
};

export const colorLabels = {
  [colors.blue]: "Blue",
  [colors.green]: "Green",
  [colors.red]: "Red",
  [colors.yellow]: "Yellow",
};

export const renderPlayerColor = ({ color }: { color: number }) => {
  return colorLabels[color];
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
