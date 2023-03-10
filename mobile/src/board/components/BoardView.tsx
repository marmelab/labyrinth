import { Board, Color } from "../BoardTypes";

import { Box, List, Grid, Stack, Typography } from "@mui/material";

import PlayerPawnView from "./PlayerPawnView";

import TileView, {
  treasures,
  type InsertTileHandler,
  type RotateRemainingTileHandler,
  type MovePlayerHandler,
} from "./TileView";

import "./BoardView.css";

const colorNames = {
  [Color.Blue]: "Blue",
  [Color.Green]: "Green",
  [Color.Red]: "Red",
  [Color.Yellow]: "Yellow",
};

interface BoardProps {
  board: Board;
  onRotateRemainingTile: RotateRemainingTileHandler;
  onInsertTile: InsertTileHandler;
  onMovePlayer: MovePlayerHandler;
}

const BoardView = ({
  board: {
    state: { tiles, remainingTile },
    players,
    gameState,
    canPlay,
    user,
  },
  onRotateRemainingTile,
  onInsertTile,
  onMovePlayer,
}: BoardProps) => {
  return (
    <>
      <div className="board">
        {tiles.flatMap((lineTiles, line) =>
          lineTiles.map((boardTile, row) => {
            return (
              <TileView
                key={`${line * tiles.length + row}`}
                canPlay={canPlay}
                boardTile={boardTile}
                gameState={gameState}
                coordinates={{ line, row }}
                onRotateRemainingTile={onRotateRemainingTile}
                onInsertTile={onInsertTile}
                onMovePlayer={onMovePlayer}
                playerTarget={user?.currentTarget}
              >
                {players
                  .filter((player) => player.line == line && player.row == row)
                  .map((player) => {
                    return (
                      <PlayerPawnView
                        key={`${player.color}`}
                        color={player.color}
                      />
                    );
                  })}
              </TileView>
            );
          })
        )}
      </div>

      <Box width={"100%"}>
        <Grid container spacing={2}>
          <Grid
            item
            xs={4}
            display={"flex"}
            alignItems={"center"}
            justifyContent={"center"}
          >
            <TileView
              boardTile={remainingTile}
              canPlay={canPlay}
              gameState={gameState}
              onRotateRemainingTile={onRotateRemainingTile}
              onInsertTile={onInsertTile}
              onMovePlayer={onMovePlayer}
              playerTarget={user?.currentTarget}
            />
          </Grid>
          <Grid item xs={8}>
            {user && (
              <Stack>
                <Grid container spacing={2}>
                  <Grid item xs={6}>
                    <Typography fontWeight={700}>Your name</Typography>
                  </Grid>
                  <Grid item xs={4}>
                    {user.name}
                  </Grid>
                </Grid>
                <Grid container spacing={2}>
                  <Grid item xs={6}>
                    <Typography fontWeight={700}>Your color</Typography>
                  </Grid>
                  <Grid item xs={4}>
                    {colorNames[user.color]}
                  </Grid>
                </Grid>
                <Grid container spacing={2}>
                  <Grid item xs={6}>
                    <Typography fontWeight={700}>Your target</Typography>
                  </Grid>
                  <Grid item xs={4}>
                    {treasures[user.currentTarget]}
                  </Grid>
                </Grid>
              </Stack>
            )}
          </Grid>
        </Grid>
      </Box>
    </>
  );
};

export default BoardView;
