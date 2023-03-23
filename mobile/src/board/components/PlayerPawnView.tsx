import type { FunctionComponent, ReactElement } from "react";

import { Color } from "../BoardTypes";

import "./PlayerPawnView.css";

interface PlayerProps {
  color: Color;
  line: number;
  row: number;
  animate: boolean;
}

const tileOffset = 50;
const tileInnerOffset = 20;

const PlayerPawnView: FunctionComponent<PlayerProps> = ({
  color,
  line,
  row,
  animate,
}: PlayerProps): ReactElement => {
  return (
    <div
      className={`player player--color-${color} ${
        animate ? "player--animate" : ""
      }`}
      style={{
        top: tileOffset * line + tileInnerOffset,
        left: tileOffset * row + tileInnerOffset,
      }}
    ></div>
  );
};

export default PlayerPawnView;
