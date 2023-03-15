export const renderGameState = ({ game_state }: { game_state: number }) => {
  switch (game_state) {
    case 0:
      return "Place Tile";
    case 1:
      return "Move Pawn";
    default:
      return "Completed";
  }
};

export const renderPlayerColor = ({ color }: { color: number }) => {
  switch (color) {
    case 0:
      return "Blue";
    case 1:
      return "Green";
    case 1:
      return "Red";
    default:
      return "Green";
  }
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
