import { Route } from "react-router-dom";

export { RemoteUserRepository } from "./UserRepository";

import { SignIn, SignOut } from "./UserPages";

export const UserRoutes = [
  <Route path="auth">
    <Route path="sign-in" element={<SignIn />} />
    <Route path="sign-out" element={<SignOut />} />
  </Route>,
];
