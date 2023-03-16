const GameState = {
  placeTile: 0,
  movePawn: 1,
  completed: 2,
};

export const renderGameState = ({ game_state }: { game_state: number }) => {
  switch (game_state) {
    case GameState.placeTile:
      return "Place Tile";
    case GameState.movePawn:
      return "Move Pawn";
    case GameState.completed:
      return "Completed";
  }
  return "";
};

const Colors = {
  blue: 0,
  green: 1,
  red: 2,
  yellow: 3,
};

export const renderPlayerColor = ({ color }: { color: number }) => {
  switch (color) {
    case Colors.blue:
      return "Blue";
    case Colors.green:
      return "Green";
    case Colors.red:
      return "Red";
    case Colors.yellow:
      return "Yellow";
  }
  return "";
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
