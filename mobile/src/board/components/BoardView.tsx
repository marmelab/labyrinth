import type { ReactNode } from "react";
import { Box, Grid, Typography } from "@mui/material";

import { type Board, Color, type Player } from "../BoardTypes";

import { treasures } from "./TileView";

import "./BoardView.css";

const colorNames = {
  [Color.Blue]: "Blue",
  [Color.Green]: "Green",
  [Color.Red]: "Red",
  [Color.Yellow]: "Yellow",
};

interface BoardProps {
  remainingTile: ReactNode;
  user?: Player | null;
  children: ReactNode;
}

const BoardStateItem = ({ label, value }: { label: string; value: string }) => (
  <Grid container spacing={2}>
    <Grid item xs={6}>
      <Typography fontWeight={700}>{label}</Typography>
    </Grid>
    <Grid item xs={4}>
      {value}
    </Grid>
  </Grid>
);

const BoardView = ({ remainingTile, user, children }: BoardProps) => {
  return (
    <>
      <div className="board">{children}</div>

      <Box width={"100%"}>
        <Grid container spacing={2}>
          <Grid
            item
            xs={4}
            display={"flex"}
            alignItems={"center"}
            justifyContent={"center"}
          >
            {remainingTile}
          </Grid>
          <Grid item xs={8}>
            {user && (
              <>
                <BoardStateItem label={"You Name"} value={user.name} />
                <BoardStateItem
                  label={"You Color"}
                  value={colorNames[user.color]}
                />
                <BoardStateItem
                  label={"You Target"}
                  value={treasures[user.currentTarget]}
                />
              </>
            )}
          </Grid>
        </Grid>
      </Box>
    </>
  );
};

export default BoardView;
