import { useEffect, useState } from "react";
import { Link as RouterLink, Outlet } from "react-router-dom";

import {
  AppBar,
  Box,
  Container,
  Link,
  MenuItem,
  Toolbar,
  Typography,
} from "@mui/material";

import { NullableUser } from "../SharedTypes";

import { useUserRepository } from "../SharedHooks";

const Layout = () => {
  const userRepository = useUserRepository();
  const [user, setUser] = useState<NullableUser>(null);

  useEffect(() => {
    userRepository.me().then((me) => {
      setUser(me);
    });
  }, []);

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
        <Outlet context={{ user, setUser }} />
      </main>
    </>
  );
};

export default Layout;
