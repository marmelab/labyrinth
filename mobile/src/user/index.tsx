import { Route } from "react-router-dom";

import { SignIn } from "./UserPages";

export { UserContext } from "./UserContext";
export { userRepository } from "./UserRepository";

export const UserRoutes = [
  <Route path="auth">
    <Route path="sign-in" element={<SignIn />} />
  </Route>,
];
