import { Route } from "react-router-dom";

import { New, List, GetById } from "./BoardPages";

export const BoardRoutes = [
  <Route path="" element={<List />} />,

  <Route path="board">
    <Route path="new" element={<New />} />
    <Route path=":id/view" element={<GetById />} />
  </Route>,
];
