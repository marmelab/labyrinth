import { Route } from "react-router-dom";

import { GetById } from "./BoardPages";

export const BoardRoutes = [
  <Route path="board">
    <Route path=":id/view" element={<GetById />} />
  </Route>,
];
