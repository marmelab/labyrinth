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

import type { User } from "../../user/UserTypes";
import type { UserRepository } from "../../user/UserRepository";
import { NullableUser } from "../SharedTypes";

interface LayoutProps {
  userRepository: UserRepository;
}

const Layout = ({ userRepository }: LayoutProps) => {
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
        <Outlet context={{ user, setUser, userRepository }} />
      </main>
    </>
  );
};

export default Layout;
