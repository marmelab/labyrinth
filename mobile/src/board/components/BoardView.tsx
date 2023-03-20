import type { ReactNode } from "react";
import {
  Alert,
  AlertColor,
  Box,
  Button,
  Grid,
  Typography,
} from "@mui/material";

import {
  Color,
  type Error,
  PlaceTileHint,
  type Player,
  GameState,
} from "../BoardTypes";

import { TREASURES } from "./Tile";

import "./BoardView.css";

const colorNames = {
  [Color.Blue]: "Blue",
  [Color.Green]: "Green",
  [Color.Red]: "Red",
  [Color.Yellow]: "Yellow",
};

interface BoardProps {
  gameState: GameState;
  remainingTile: ReactNode;
  user?: Player | null;
  children: ReactNode;
  errors: Error[];
  placeTileHint?: PlaceTileHint | null;
  handleGetPlaceTileHint: () => void;
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

const BoardView = ({
  gameState,
  remainingTile,
  user,
  children,
  errors,
  handleGetPlaceTileHint,
}: BoardProps) => {
  return (
    <>
      <div className="board">
        {children}

        {remainingTile}
      </div>

      <Box width="355px">
        <Grid container spacing={2}>
          <Grid item xs={8} ml={"auto"}>
            {user && (
              <>
                <BoardStateItem label={"You Name"} value={user.name} />
                <BoardStateItem
                  label={"You Color"}
                  value={colorNames[user.color]}
                />
                <BoardStateItem
                  label={"You Target"}
                  value={TREASURES[user.currentTarget]}
                />
              </>
            )}
          </Grid>
        </Grid>
      </Box>

      {user?.isCurrentPlayer && (
        <Box width="355px">
          <Grid container>
            <Grid item xs={12} display="flex" justifyContent="flex-end">
              {gameState == GameState.PlaceTile && (
                <Button variant="outlined" onClick={handleGetPlaceTileHint}>
                  Get Hint
                </Button>
              )}
            </Grid>

            {errors.map((error, i) => (
              <Grid item xs={12} mt={2} key={i}>
                <Alert severity={error.severity}>{error.message}</Alert>
              </Grid>
            ))}
          </Grid>
        </Box>
      )}
    </>
  );
};

export default BoardView;
