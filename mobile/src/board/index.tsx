import { Route } from "react-router-dom";

import { List, GetById } from "./BoardPages";

export const BoardRoutes = [
  <Route path="" element={<List />} />,

  <Route path="board">
    <Route path=":id/view" element={<GetById />} />
  </Route>,
];
