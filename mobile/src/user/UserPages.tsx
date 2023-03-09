import { type MouseEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, FormControl, TextField } from "@mui/material";

import { useUser } from "./UserHooks";
import { userRepository } from "./UserRepository";

export function SignIn() {
  const navigate = useNavigate();
  const [user, setUser] = useUser();

  const [name, setName] = useState<string>();

  const handleSubmit = async function (e: MouseEvent) {
    e.preventDefault();
    if (!name) {
      return;
    }

    try {
      const signedInUser = await userRepository.signIn(name);
      setUser(signedInUser);
      navigate("/");
    } catch (e) {
      console.error(e);
    }
  };

  return user ? (
    <p>You are signed in as {user.name}</p>
  ) : (
    <FormControl>
      <TextField
        label="Username"
        variant="filled"
        onChange={(event) => setName(event.target.value)}
        sx={{ mb: 2 }}
      />
      <Button type="submit" variant="contained" onClick={handleSubmit}>
        Sign In / Sign Up
      </Button>
    </FormControl>
  );
}
