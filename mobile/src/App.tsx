import { useState } from "react";
import {
  createBrowserRouter,
  createRoutesFromElements,
  RouterProvider,
  Route,
  type RouteObject,
} from "react-router-dom";
import type { Router } from "@remix-run/router";

import Layout from "./shared/components/Layout";

import type { NullableUser } from "./user/UserTypes";

import BoardRoutes from "./board";
import { UserRoutes, UserContext } from "./user";

import "./App.css";

type CreateRouterFunction = (routes: RouteObject[]) => Router;

export default function App({
  createRouter = createBrowserRouter,
}: {
  createRouter?: CreateRouterFunction;
}) {
  const router = createRouter(
    createRoutesFromElements(
      <Route path="/" element={<Layout />}>
        {...BoardRoutes}
        {...UserRoutes}
      </Route>
    )
  );
  const [user, setUser] = useState<NullableUser>(null);

  return (
    <UserContext.Provider value={[user, setUser]}>
      <RouterProvider router={router} />
    </UserContext.Provider>
  );
}
