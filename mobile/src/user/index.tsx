import { Route } from "react-router-dom";

import { SignUp, SignIn } from "./UserPages";

export { UserContext } from "./UserContext";

export type { User, NullableUser } from "./UserTypes";

export const UserRoutes = [
  <Route path="auth">
    <Route path="sign-up" element={<SignUp />} />
    <Route path="sign-in" element={<SignIn />} />
  </Route>,
];
