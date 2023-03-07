import type { FunctionComponent, MouseEventHandler, ReactElement } from "react";
import { Color } from "../../model/Player";

import "./index.css";

interface PlayerProps {
  color: Color;
}

const PlayerPawn: FunctionComponent<PlayerProps> = ({
  color,
}: PlayerProps): ReactElement => {
  return <div className={`player player--color-${color}`}></div>;
};

export default PlayerPawn;
