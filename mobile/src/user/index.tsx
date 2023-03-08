import { Route } from "react-router-dom";

import { SignIn, SignOut } from "./pages";

export default [
  <Route path="auth">
    <Route path="sign-in" element={<SignIn />} />
    <Route path="sign-out" element={<SignOut />} />
  </Route>,
];
