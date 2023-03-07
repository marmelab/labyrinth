import type { RouteObject } from "react-router-dom";

import { getByIdLoader, GetById } from "./GetById";

const routes: Array<RouteObject> = [
  {
    path: "/board/:id/view",
    loader: getByIdLoader,
    element: <GetById />,
  },
];

export default routes;
