import type { ReactElement, FormEvent } from "react";

import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import Button from "@mui/material/Button";
import FormControl from "@mui/material/FormControl";
import TextField from "@mui/material/TextField";

import { useRemoteUserRepository, useUser } from "../shared/hooks";

export function SignIn(): ReactElement {
  const navigate = useNavigate();
  const remoteUserRepository = useRemoteUserRepository();
  const [_, setUser] = useUser();

  const [name, setName] = useState<string>();

  const handleSubmit = async function (e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    if (!name) {
      return;
    }

    try {
      const signedInUser = await remoteUserRepository.signIn(name);
      setUser(signedInUser);
      navigate("/");
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <form onSubmit={handleSubmit} method="POST">
      <FormControl>
        <TextField
          id="filled-basic"
          label="Username"
          variant="filled"
          onChange={({ target: { value } }) => setName(value)}
          sx={{ mb: 2 }}
        />
        <Button type="submit" variant="contained">
          Sign In / Sign Up
        </Button>
      </FormControl>
    </form>
  );
}

export function SignOut(): ReactElement {
  const navigate = useNavigate();
  const remoteUserRepository = useRemoteUserRepository();
  const [_, setUser] = useUser();

  useEffect(() => {
    remoteUserRepository.signOut().then(() => {
      setUser(null);
      navigate("/");
    });
  });

  return <></>;
}
