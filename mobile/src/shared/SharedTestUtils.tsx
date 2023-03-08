import type { ReactNode } from "react";
import { MemoryRouter, Route, Routes } from "react-router-dom";

import Layout from "./components/Layout";

import type { UserRepository } from "../user/UserRepository";

interface RenderRouteWithOutletContextProps {
  remoteUserRepository: UserRepository;
  children: ReactNode;
}

export const RenderRouteContext = ({
  remoteUserRepository,
  children,
}: RenderRouteWithOutletContextProps) => {
  return (
    <MemoryRouter>
      <Routes>
        <Route
          path="/"
          element={<Layout remoteUserRepository={remoteUserRepository} />}
        >
          <Route index element={children} />
        </Route>
      </Routes>
    </MemoryRouter>
  );
};
