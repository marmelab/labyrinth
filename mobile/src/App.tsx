import { useState } from "react";
import { Route, Routes } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "react-query";

import Layout from "./shared/components/Layout";

import type {} from "./user/UserTypes";

import { BoardRoutes } from "./board";
import { UserRoutes, UserContext, NullableUser } from "./user";

import "./App.css";

const queryClient = new QueryClient();

export default function App() {
  const [user, setUser] = useState<NullableUser>(null);

  return (
    <QueryClientProvider client={queryClient}>
      <UserContext.Provider value={[user, setUser]}>
        <Routes>
          <Route path="/" element={<Layout />}>
            {...BoardRoutes}
            {...UserRoutes}
          </Route>
        </Routes>
      </UserContext.Provider>
    </QueryClientProvider>
  );
}
