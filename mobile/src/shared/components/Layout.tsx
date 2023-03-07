import type { FunctionComponent, ReactElement } from "react";
import { Link as RouterLink, Outlet } from "react-router-dom";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";
import MenuItem from "@mui/material/MenuItem";
import Toolbar from "@mui/material/Toolbar";

const Layout: FunctionComponent<{}> = (): ReactElement => {
  return (
    <>
      <AppBar>
        <Container maxWidth="xl">
          <Toolbar disableGutters>
            <Box sx={{ flexGrow: 1 }}>
              <MenuItem>
                <Link
                  component={RouterLink}
                  to="/"
                  sx={{ my: 2, color: "white", display: "block" }}
                >
                  Home
                </Link>
              </MenuItem>
            </Box>

            <Box>
              <MenuItem>
                <Link
                  component={RouterLink}
                  to="/auth/sign-in"
                  sx={{ my: 2, color: "white", display: "block" }}
                >
                  Sign In / Sign Up
                </Link>
              </MenuItem>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
      <main>
        <Outlet />
      </main>
    </>
  );
};

export default Layout;
