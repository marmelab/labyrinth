import type { FunctionComponent, ReactElement } from "react";
import { Outlet } from "react-router-dom";

const Layout: FunctionComponent<{}> = (): ReactElement => {
  return (
    <>
      <header></header>
      <main>
        <Outlet />
      </main>
    </>
  );
};

export default Layout;
