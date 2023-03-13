import { useContext, useEffect, type ReactNode } from "react";
import { useNavigate, Link as RouterLink, Outlet } from "react-router-dom";

import {
  AppBar,
  Box,
  Container,
  Link,
  MenuItem,
  Toolbar,
  Typography,
} from "@mui/material";

import { UserContext } from "../../user";
import { useGetIdentityQuery, useSignOutMutation } from "../../user/UserHooks";

const NavLink = function ({
  to,
  onClick,
  children,
}: {
  to?: string;
  onClick?: () => Promise<void>;
  children: ReactNode;
}) {
  const styles = { my: 2, color: "white", display: "block" };
  return (
    <>
      {to && (
        <Link component={RouterLink} to={to} sx={styles}>
          {children}
        </Link>
      )}
      {onClick && (
        <Link onClick={onClick} sx={styles}>
          {children}
        </Link>
      )}
    </>
  );
};

const Layout = () => {
  const navigate = useNavigate();
  const [user, setUser] = useContext(UserContext);
  const signOutMutation = useSignOutMutation();

  const identityQuery = useGetIdentityQuery({
    onSuccess: (user) => {
      setUser(user);
    },
  });

  const signOut = async function () {
    await signOutMutation.mutateAsync();
    setUser(null);
    navigate("/");
  };

  return (
    <>
      <AppBar>
        <Container maxWidth="xl">
          <Toolbar disableGutters>
            <Box flexGrow={1}>
              <MenuItem>
                <NavLink to="/">Home</NavLink>
              </MenuItem>
            </Box>

            <Box>
              <MenuItem
                sx={{ display: "flex", flexDirection: "row", gap: "20px" }}
              >
                {!user ? (
                  <>
                    <NavLink to="/auth/sign-in">Sign In</NavLink>
                    <NavLink to="/auth/sign-up">Sign Up</NavLink>
                  </>
                ) : (
                  <>
                    <Typography sx={{ color: "white", display: "block" }}>
                      {user.username}
                    </Typography>

                    <NavLink onClick={signOut}>Sign Out</NavLink>
                  </>
                )}
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
