import React, { ReactElement, useCallback, useState } from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  createRoutesFromElements,
  RouterProvider,
  Route,
} from "react-router-dom";

import { User } from "./user/entity";
import { UserContext } from "./user/context";

import Layout from "./shared/components/Layout";

import BoardRoutes from "./board";
import UserRoutes from "./user";

import "./index.css";

const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<Layout />}>
      {...BoardRoutes}
      {...UserRoutes}
    </Route>
  )
);

const App = function (): ReactElement {
  const [user, setUser] = useState<User | null>(null);

  return (
    <UserContext.Provider value={[user, setUser]}>
      <RouterProvider router={router} />
    </UserContext.Provider>
  );
};

ReactDOM.createRoot(document.querySelector("#app") as HTMLElement).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
