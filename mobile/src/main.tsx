import React, { ReactElement, useCallback, useState } from "react";
import ReactDOM from "react-dom/client";
import {
  createBrowserRouter,
  createRoutesFromElements,
  RouterProvider,
  Route,
} from "react-router-dom";

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

ReactDOM.createRoot(document.querySelector("#app") as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);
