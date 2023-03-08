import type { ReactNode } from "react";
import { MemoryRouter, Route, Routes } from "react-router-dom";

import Layout from "./components/Layout";

import type { UserRepository } from "../user/UserRepository";

interface RenderRouteWithOutletContextProps {
  userRepository: UserRepository;
  children: ReactNode;
}

export const RenderRouteContext = ({
  userRepository,
  children,
}: RenderRouteWithOutletContextProps) => {
  return (
    <MemoryRouter>
      <Routes>
        <Route path="/" element={<Layout userRepository={userRepository} />}>
          <Route index element={children} />
        </Route>
      </Routes>
    </MemoryRouter>
  );
};
