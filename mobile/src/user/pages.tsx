import { ReactElement } from "react";

import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";

export function SignIn(): ReactElement {
  return (
    <>
      <TextField id="filled-basic" label="Username" variant="filled" />
      <Button variant="contained">Sign In / Sign Up</Button>
    </>
  );
}
