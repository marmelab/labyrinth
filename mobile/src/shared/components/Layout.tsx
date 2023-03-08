import type { FunctionComponent, ReactElement } from "react";

import { useState } from "react";
import { Link as RouterLink, Outlet } from "react-router-dom";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Link from "@mui/material/Link";
import MenuItem from "@mui/material/MenuItem";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";

import type { User } from "../../user/entity";
import { RemoteUserRepository } from "../../user/repository";

const Layout: FunctionComponent<{}> = (): ReactElement => {
  const [user, setUser] = useState<User | null>(null);

  const remoteUserRepository = new RemoteUserRepository();

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
                {!user ? (
                  <Link
                    component={RouterLink}
                    to="/auth/sign-in"
                    sx={{ my: 2, color: "white", display: "block" }}
                  >
                    Sign In / Sign Up
                  </Link>
                ) : (
                  <>
                    <Typography sx={{ m: 2, color: "white", display: "block" }}>
                      {user.name}
                    </Typography>

                    <Link
                      component={RouterLink}
                      to="/auth/sign-out"
                      sx={{ my: 2, color: "white", display: "block" }}
                    >
                      Sign Out
                    </Link>
                  </>
                )}
              </MenuItem>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
      <main>
        <Outlet context={{ user, setUser, remoteUserRepository }} />
      </main>
    </>
  );
};

export default Layout;
