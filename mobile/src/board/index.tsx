import { Route } from "react-router-dom";

import { GetById } from "./pages";
import { BoardRepository } from "./repository";

export default [
  <Route path="board">
    <Route path=":id/view" element={<GetById />} />
  </Route>,
];
