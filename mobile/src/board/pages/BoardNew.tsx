import { useState, type MouseEvent } from "react";
import { useNavigate } from "react-router-dom";

import {
  Button,
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  Typography,
} from "@mui/material";

import { useNewBoardMutation } from "../BoardHooks";

import { useUserContext } from "../../user/UserContext";

export function New() {
  const navigate = useNavigate();
  const [user, _] = useUserContext();
  const mutation = useNewBoardMutation();

  const [playerCount, setPlayerCount] = useState("1");

  const handlePlayerCountChange = (e: { target: { value: string } }) => {
    setPlayerCount(e.target.value);
  };

  const handleSubmit = async (e: MouseEvent) => {
    e.preventDefault();

    const board = await mutation.mutateAsync({ playerCount: +playerCount });
    navigate(`/board/${board.id}/view`);
  };

  if (!user) {
    return (
      <Typography fontWeight={700}>
        You must be signed in to access this page.
      </Typography>
    );
  }

  return (
    <>
      {mutation.isLoading ? (
        <Typography>Creating Board...</Typography>
      ) : (
        <FormControl>
          {mutation.isError ? (
            <Typography color={"red"}>
              An error occurred: {mutation.error.message}
            </Typography>
          ) : null}
          <FormControl fullWidth>
            <InputLabel id="player-count">Player Count</InputLabel>
            <Select
              labelId="player-count"
              value={playerCount}
              label="Player count"
              onChange={handlePlayerCountChange}
              sx={{ mb: 2 }}
            >
              {[1, 2, 3, 4].map((count) => (
                <MenuItem key={count} value={count}>
                  {count}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
          <Button variant="contained" onClick={handleSubmit}>
            Create New Game
          </Button>
        </FormControl>
      )}
    </>
  );
}
