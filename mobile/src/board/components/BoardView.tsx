import type { ReactNode } from "react";
import {
  Alert,
  Box,
  Button,
  Grid,
  Table,
  TableHead,
  TableBody,
  TableRow,
  TableCell,
  Typography,
} from "@mui/material";

import { Color, type Error, type Player } from "../BoardTypes";

import { TREASURES } from "./Tile";

import "./BoardView.css";

const colorNames = {
  [Color.Blue]: "Blue",
  [Color.Green]: "Green",
  [Color.Red]: "Red",
  [Color.Yellow]: "Yellow",
};

interface BoardProps {
  canPlay: boolean;
  remainingTile: ReactNode;
  user?: Player | null;
  currentPlayer?: Player | null;
  players: Player[];
  children: ReactNode;
  errors: Error[];
  handleGetHint: () => void;
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

const BoardPlayers = ({ players }: { players: Player[] }) => (
  <Table aria-label="simple table">
    <TableHead>
      <TableRow>
        <TableCell>Name</TableCell>
        <TableCell>Score</TableCell>
        <TableCell>End</TableCell>
      </TableRow>
    </TableHead>
    <TableBody>
      {players.map((player) => (
        <TableRow key={player.name}>
          <TableCell component="th" scope="row">
            {player.name}
          </TableCell>
          <TableCell>
            {player.score} / {player.totalTargets}
          </TableCell>
          <TableCell>{player.winOrder}</TableCell>
        </TableRow>
      ))}
    </TableBody>
  </Table>
);

const BoardView = ({
  canPlay,
  remainingTile,
  user,
  currentPlayer,
  players,
  children,
  errors,
  handleGetHint,
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

      <Box width="355px">
        <Grid container>
          <Grid
            item
            xs={12}
            display="flex"
            justifyContent="flex-end"
            minHeight={50}
          >
            {canPlay && (
              <Button variant="outlined" onClick={handleGetHint}>
                Get Hint
              </Button>
            )}
          </Grid>

          {!canPlay && currentPlayer && (
            <Grid item xs={12} mb={2}>
              <Alert severity="info">
                <strong>Waiting:</strong> {currentPlayer.name}
              </Alert>
            </Grid>
          )}
          {canPlay &&
            errors.map((error, i) => (
              <Grid item xs={12} mt={2} key={i}>
                <Alert severity={error.severity}>{error.message}</Alert>
              </Grid>
            ))}
        </Grid>
      </Box>
      <Box width="355px" mt={2}>
        <BoardPlayers players={players} />
      </Box>
    </>
  );
};

export default BoardView;
