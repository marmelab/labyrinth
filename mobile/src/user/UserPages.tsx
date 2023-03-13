import { type MouseEvent, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, FormControl, TextField, Typography } from "@mui/material";

import { useUserContext } from "./UserContext";
import { useSignUpMutation, useSignInMutation } from "./UserHooks";

export function SignUp() {
  const navigate = useNavigate();
  const [user, setUser] = useUserContext();
  const mutation = useSignUpMutation();

  const [email, setEmail] = useState<string>("");
  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleSubmit = async function (e: MouseEvent) {
    e.preventDefault();
    try {
      await mutation.mutateAsync({
        email,
        username,
        password,
      });
      navigate("/");
    } catch (e) {
      console.error(e);
    }
  };

  return user ? (
    <p>You are signed in as {user.username}</p>
  ) : (
    <>
      {mutation.isError && (
        <Typography color={"red"}>{mutation.error}</Typography>
      )}
      <FormControl>
        <TextField
          label="Username"
          variant="filled"
          onChange={(event) => setUsername(event.target.value)}
          sx={{ mb: 2 }}
        />
        <TextField
          label="Email"
          variant="filled"
          onChange={(event) => setEmail(event.target.value)}
          sx={{ mb: 2 }}
        />
        <TextField
          label="Password"
          variant="filled"
          type="password"
          onChange={(event) => setPassword(event.target.value)}
          sx={{ mb: 2 }}
        />
        <Button type="submit" variant="contained" onClick={handleSubmit}>
          Sign Up
        </Button>
      </FormControl>
    </>
  );
}

export function SignIn() {
  const navigate = useNavigate();
  const [user, setUser] = useUserContext();
  const mutation = useSignInMutation();

  const [username, setUsername] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleSubmit = async function (e: MouseEvent) {
    e.preventDefault();
    if (!username && !password) {
      return;
    }

    try {
      const signedInUser = await mutation.mutateAsync({ username, password });
      setUser(signedInUser);
      navigate("/");
    } catch (e) {
      console.error(e);
    }
  };

  return user ? (
    <p>You are signed in as {user.username}</p>
  ) : (
    <>
      {mutation.isError && (
        <Typography color={"red"}>{mutation.error}</Typography>
      )}
      <FormControl>
        <TextField
          label="Email"
          variant="filled"
          onChange={(event) => setUsername(event.target.value)}
          sx={{ mb: 2 }}
        />
        <TextField
          label="Password"
          variant="filled"
          type="password"
          onChange={(event) => setPassword(event.target.value)}
          sx={{ mb: 2 }}
        />
        <Button type="submit" variant="contained" onClick={handleSubmit}>
          Sign In
        </Button>
      </FormControl>
    </>
  );
}
