import { Route } from "react-router-dom";

import { GetById } from "./BoardPages";

export default [
  <Route path="board">
    <Route path=":id/view" element={<GetById />} />
  </Route>,
];
