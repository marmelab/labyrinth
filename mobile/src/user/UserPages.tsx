import { FormEvent, useEffect, useState } from "react";
import { Route, useNavigate } from "react-router-dom";
import { Button, FormControl, TextField } from "@mui/material";

import { useUserRepository, useUser } from "../shared/SharedHooks";

export function SignIn() {
  const navigate = useNavigate();
  const userRepository = useUserRepository();
  const [_, setUser] = useUser();

  const [name, setName] = useState<string>();

  const handleSubmit = async function (e: FormEvent<HTMLFormElement>) {
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

export function SignOut() {
  const navigate = useNavigate();
  const userRepository = useUserRepository();
  const [_, setUser] = useUser();

  useEffect(() => {
    userRepository.signOut().then(() => {
      setUser(null);
      navigate("/");
    });
  });

  return <></>;
}

export const UserRoutes = [
  <Route path="auth">
    <Route path="sign-in" element={<SignIn />} />
    <Route path="sign-out" element={<SignOut />} />
  </Route>,
];
