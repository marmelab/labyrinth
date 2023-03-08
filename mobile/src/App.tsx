import { useState } from "react";
import { Route, Routes } from "react-router-dom";

import Layout from "./shared/components/Layout";

import type {} from "./user/UserTypes";

import { BoardRoutes } from "./board";
import { UserRoutes, UserContext, NullableUser } from "./user";

import "./App.css";

export default function App() {
  const [user, setUser] = useState<NullableUser>(null);

  return (
    <UserContext.Provider value={[user, setUser]}>
      <Routes>
        <Route path="/" element={<Layout />}>
          {...BoardRoutes}
          {...UserRoutes}
        </Route>
      </Routes>
    </UserContext.Provider>
  );
}
