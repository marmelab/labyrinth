import { Route } from "react-router-dom";

import { SignIn } from "./pages";
import { UserRepository } from "./repository";

export default [
  <Route path="auth">
    <Route path="sign-in" element={<SignIn />} />
  </Route>,
];
