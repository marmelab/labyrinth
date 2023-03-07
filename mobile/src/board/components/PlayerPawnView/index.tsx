import type { FunctionComponent, ReactElement } from "react";

import { Color } from "../../entity";

import "./index.css";

interface PlayerProps {
  color: Color;
}

const PlayerPawnView: FunctionComponent<PlayerProps> = ({
  color,
}: PlayerProps): ReactElement => {
  return <div className={`player player--color-${color}`}></div>;
};

export default PlayerPawnView;
